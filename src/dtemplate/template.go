package dtemplate

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	`os/exec`
	`sync`

	"gopkg.in/yaml.v2"
	// "github.com/juju/errors"

	`dtemplate/xmlparse`
	`dtemplate/config`
)

type Index struct {
	Name string
	Node *Node
	Pos  []int
}

func newIndex(name string, node *Node, pos []int) *Index {
	i := &Index{
		Name: name,
		Node: node,
		Pos:  make([]int, len(pos)),
	}
	copy(i.Pos, pos)
	return i
}

func (idx *Index) IsThis() bool {
	return `this` == idx.Name
}

func (idx *Index) Path() string {
	if nil == idx {
		return ``
	}
	path := make([]string, len(idx.Pos))
	for i, p := range idx.Pos {
		// path[i] = fmt.Sprintf(`.children[%d]`, p)
		path[i] = fmt.Sprintf(`.childNodes[%d]`, p)
	}
	if 0<len(path) {
		path[0]=``
	}
	return strings.Join(path, ``)
}

func (idx *Index) CljPath() string {
	if nil==idx {
		return ``
	}
	path := make([]string, len(idx.Pos))
	for i, p := range idx.Pos {
		path[i] = fmt.Sprintf("%d", p)
	}
	// if idx.IsThis() && 0<len(path) {
	if 0 < len(path) {
		path[0] = ``
	}
	return strings.Join(path, ` `)
}

type Template struct {
	Name    string
	Node    *Node
	Raw     []byte
	Indices []*Index
}

// This returns the `this` index, or the
// root Node if there is no `this` Node
func (t *Template) This() *Index {
	if 0 == len(t.Indices) {
		return nil
	}
	i := t.Indices[0]
	if i.IsThis() {
		return i
	}
	return &Index{
		Name: `this`,
		Node: t.Node,
		Pos:  []int{},
	}
}

func (t *Template) Html() string {
	return t.Node.Render()
}

func mapTemplates(templates []*Template) map[string]string {
	m := map[string]string{}
	for _, t := range templates {
		m[t.Name] = string(t.Raw)
	}
	return m
}

func findFirstElement(n *Node) *Node {
	c := n.FirstChild()
	for nil != c {
		if c.IsElement() {
			return c
		}
		c = c.NextSibling()
	}
	return nil
}

func findIndices(attr string, node *Node) []*Index {
	indices := []*Index{}
	// Check whether the parent node is a attributed node
	for a,v := range node.Attributes() {
		name := a
		index := false
		if attr == a {
			index = true
			name = v
			node.RemoveAttribute(attr)
		} else if ``!=a && '$'==a[1] {
			node.RemoveAttribute(a)
			name = a[1:]
			index = true
		}
		if index {
			indices = append(indices, newIndex(name, node, []int{}))
		}
	}
	// name := node.GetAttribute(attr)
	// if `` != name {
	// 	indices = append(indices, newIndex(name, node, []int{}))
	// 	// Remove the attribute after capturing the index
	// 	node.RemoveAttribute(attr)
	// }
	findIndices_recurse(attr, node, []int{}, &indices)
	return indices
}

func findIndices_recurse(attr string, parent *Node, path []int, indices *[]*Index) {
	i := 0
	n := parent.FirstChild()
	lenp := len(path)
	cpath := make([]int, lenp+1)
	copy(cpath, path)
	for nil != n {
		cpath[lenp] = i
		if n.IsElement() {
			name := n.GetAttribute(attr)
			if `` != name {
				if `this` == name {
					*indices = []*Index{newIndex(name, n, cpath)}
					// If we've found our 'this' node, we will
					// reset all our indices so that the first one points
					// to 'this', then we'll only recurse over our
					// children of 'this'.
					findIndices_recurse(attr, n, []int{}, indices)
					return
				}
				*indices = append(*indices, newIndex(name, n, cpath))
				// Remove the attribute after capturing the index
				n.RemoveAttribute(attr)
			}
			// Only if this is an element, we recurse
			findIndices_recurse(attr, n, cpath, indices)
		}
		n = n.NextSibling()
		i++
	}
}

func loadTemplates(dir, nameSeparator string, cfg *config.Config) ([]*Template, error) {

	if strings.HasSuffix(dir, `/`) {
		dir = dir[0: len(dir)-1]
	}
	templates := []*Template{}

	var wait sync.WaitGroup
	ERR := make(chan error)
	var templatesLock sync.Mutex

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if nil != err {
			return err
		}
		ext := filepath.Ext(path)
		if ".html" != ext {
			// Only consider .html files
			return nil
		}
		wait.Add(1)
		go func(path, ext string) {
			if err := func(path, ext string) error {
				defer wait.Done()
				relPath := path[len(dir)+1 : len(path)-len(ext)]
				raw, err := ioutil.ReadFile(path)
				if nil != err {
					return err
				}

				// fmt.Println(`ABSOLUTE RAW = ----`,string(raw),`------`)
				metaRaw, xmlRaw, err := splitMetadata(bytes.NewReader(raw))
				if nil != err {
					return fmt.Errorf(`Failed to parse %s: %s`, path, err.Error())
				}

				settings := map[string]interface{}{}
				if err := yaml.Unmarshal(metaRaw, &settings); nil != err {
					return fmt.Errorf(`Failed parsing yaml metadata in %s: %s`, path, err.Error())
				}

				node, err := ParseNode(bytes.NewReader(xmlRaw))
				if nil != err {
					return fmt.Errorf(`Failed to parse HTML in %s`, path)
				}
				// fmt.Println(`RAW bytes = `, string(xmlRaw))
				// fmt.Println(`-----`)
				// fmt.Println(`unprocessed node = `)
				// fmt.Println(node.Node.RawString())

				if err := processNodes(path, &node.Node, settings, cfg); nil!=err {
					return fmt.Errorf(`Failed processing nodes in %s: %s`, path, err.Error())
				}

				// childTemplates are constructed by walking the node tree looking for 
				// [dtemplate-child] elements. We don't remove them until all have been found,
				// and only after they are all removed, do we set indices and the actual content,
				// so that we don't get child templates appearing inside parent templates
				childTemplates := []*Template{}
				if err := xmlparse.Walk(&node.Node, func(n *xmlparse.Node, depth int) error {
					el, ok := (*n).(*xmlparse.Element)
					if !ok {
						return nil
					}
					if ``!=el.GetAttribute(`dtemplate-child`) {
						childTemplateName := strings.Replace(
								relPath + "." +
								strings.Join(getAncestorAttributes(el, `dtemplate-child`), `.`), "/", nameSeparator, -1)
						fmt.Println(`Creating childTemplate with name `, childTemplateName)
						childTemplate := &Template{
							Name: childTemplateName,
							Node: &Node{*n},
						}
						childTemplates = append(childTemplates, childTemplate)
						//log.Infof(`Found child template %s : raw = %s`, childTemplate.Name, childTemplate.Node.
					}
					return nil
				}); nil!=err {
					return fmt.Errorf(`Failed processing child template in %s: %s`, path, err.Error())
				}
				// Next, we remove all the child template nodes
				for _, c := range childTemplates {
					c.Node.Node.Remove()
				}
				// Only now can we set Raw content and indices
				for _, c := range childTemplates {
					c.Raw = []byte(c.Node.Node.RawString())
					c.Indices = findIndices(`data-set`, c.Node)
					fmt.Println(`---`, c.Name)
					fmt.Println(string(c.Raw))
				}



				// With libxml2, our node is already the first-child
				// element
				t := &Template{
					Name:    strings.Replace(relPath, "/", nameSeparator, -1),
					Node:    node,
					Raw:     []byte((*node).Node.RawString()),
					Indices: findIndices(`data-set`, node),
				}
				if 0<len(childTemplates) {
					fmt.Println(`---`, t.Name)
					fmt.Println(string(t.Raw))					
				}
				// fmt.Println(`-----`)
				// fmt.Println(`raw node=`)
				// fmt.Println((t.Node.Node).RawString())
				// fmt.Println(`---`)
				templatesLock.Lock()
				defer templatesLock.Unlock()
				templates = append(templates, t)
				// add all child templates
				templates = append(templates, childTemplates...)
				return nil
			}(path, ext); nil!=err {
				ERR <- err
			}
		}(path, ext)

		return nil
	})
	if nil != err {
		return nil, err
	}
	go func() {
		wait.Wait()
		close(ERR)
	}()

	for err = range ERR {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	if nil!=err {
		return nil, err
	}
	return templates, nil
}

func getAncestorAttributes(e *xmlparse.Element, attr string) []string {
	if nil==e.Parent() {
		if e.HasAttribute(attr) {
			return []string{e.GetAttribute(attr)}
		} else {
			return []string{}
		}
	}
	if e.HasAttribute(attr) {
		return append(getAncestorAttributes(e.Parent(), attr), e.GetAttribute(attr))
	} else {
		return getAncestorAttributes(e.Parent(), attr)
	}
}

func LineReader(in io.Reader) (LINE chan string, ERR chan error) {
	LINE=make(chan string)
	ERR = make(chan error,1)
	go func() {
		defer func() {
			close(LINE)
			close(ERR)
		}()
		read := bufio.NewReader(in)
		for {
			line, err := read.ReadString('\n')
			if nil!=err && io.EOF != err {
				ERR <- err
				return
			}
			// trim trailing \n
			if (0<len(line) && line[len(line)-1]=='\n') {
				line = line[0:len(line)-1]
			}
			LINE <- line
			if io.EOF == err {
				return
			}
		}
	}()
	return LINE, ERR
}

func splitMetadata(in io.Reader) ([]byte, []byte, error) {
	var yml *bytes.Buffer
	var data *bytes.Buffer
	// Possible states are
	// 0 => looking for START OF XML, START OF YML, or YML header
	// 1 => scanning YML
	// 2 => scanning XML
	state := 0 //
	LINES, ERR := LineReader(in)
	lineCount := 0
	for rawline := range LINES {
		lineCount++
		// trim the newline from the read rawline
		line := strings.TrimSpace(rawline)
		if "" == line {
			continue
		}
		switch state {
		case 0:
			if '<' == line[0] {
				data = bytes.NewBuffer([]byte(rawline))
				data.WriteString("\n")
				state = 2
				continue
			}
			if "---" == line {
				// This has to be a start of yml
				yml = bytes.NewBuffer([]byte{})
				state = 1
				continue
			}
			// At this point we must have YML starting, without
			// the --- prefix, which is fine- we can manage that
			yml = bytes.NewBuffer([]byte(rawline))
			yml.WriteString("\n")
			state = 1
			continue
		case 1:
			if "---" == line {
				data = bytes.NewBuffer([]byte{})
				state = 2
				continue
			}
			yml.WriteString(rawline)
			yml.WriteString("\n")
			continue
		case 2:
			data.WriteString(rawline)
			data.WriteString("\n")
		}
	}
	if err := <- ERR; nil!=err {
		return []byte{}, nil, err
	}
	if nil == data {
		return nil, nil, fmt.Errorf(`No XML data found in file after %d lines: did you finish the metadata with '---' on a separate line ?`, lineCount)
	}
	if nil != yml {
		return yml.Bytes(), data.Bytes(), nil
	}
	return []byte{}, data.Bytes(), nil
}

func processNodes(srcFile string, node *xmlparse.Node, settings map[string]interface{}, cfg *config.Config) error {
	return xmlparse.Walk(node, func(n *xmlparse.Node, depth int) error {
		el, ok := (*n).(*xmlparse.Element)
		if !ok {
			return nil
		}
		absFile, err := filepath.Abs(srcFile)
		if nil!=err {
			return fmt.Errorf(`Failed to read source file '%s': %w`, srcFile, err)
		}
		// We process -include before -process, which means that we can run
		// -process on a -include'd file (eg. for scss)
		if ``!=el.GetAttribute(`dtemplate-include`) {
			file := el.GetAttribute(`dtemplate-include`)
			absFile, err = filepath.Abs(filepath.Join(filepath.Dir(srcFile), file))
			if nil!=err {
				return fmt.Errorf(`Failed to resolve include file '%s' in '%s': %w`, file, srcFile, err)
			}
			in, err := os.ReadFile(absFile)
			if nil!=err {
				return fmt.Errorf(`Failed to open include file '%s' in '%s': %w`,
					file, srcFile, err)
			}
			el.SetInnerText(string(in))
			el.RemoveAttribute(`dtemplate-include`);
		}
		if ``!=el.GetAttribute(`dtemplate-process`) {
			proc := el.GetAttribute(`dtemplate-process`)

			var process *config.Process
			ms, ok := settings[proc]
			if ok {
				process = config.NewProcessFromMap(ms.(map[string]interface{}))
			} else {
				process, ok = cfg.Process[proc]
				if !ok {
					process = &config.Process{Exec: proc}
				}
			}

			// @TODO Should properly parse the Exec to enable delimited arguments
			args := strings.Split(strings.TrimSpace(process.Exec), ` `)
			cwd, _ := os.Getwd()
			// replace special templated strings in command arguments
			replacements := map[string]string {
				`%TEMPLATE_DIR%`: filepath.Dir(absFile),
				`%TEMPLATE_FILE%`: absFile,
				`%.%`: cwd,
			}
			for i, a := range args {
				for k, v := range replacements {
					a = strings.ReplaceAll(a,k,v)
				}
				args[i] = a
			}
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Dir = filepath.Dir(absFile)

			var out bytes.Buffer
			cmd.Stdout = &out
			rin, win, err := os.Pipe()
			if nil!=err {
				return err
			}
			cmd.Stdin = rin
			raw := el.InnerRawText()
			// fmt.Println("---")
			// fmt.Println("DIR = ", cmd.Dir)
			// fmt.Println(strings.Join(args, " "))
			// fmt.Println("--- Converting")
			// fmt.Println(raw)
			go func() {
				fmt.Fprintln(win, process.Prefix)
				io.Copy(win, strings.NewReader(raw))
				fmt.Fprintln(win, process.Suffix)
				win.Close()
			}()
			cmd.Stderr = os.Stderr


			if err := cmd.Run(); nil!=err {
				return fmt.Errorf(`Failed running [ %s ] : %w`, strings.Join(args, ` `), err)
			}
			el.SetInnerText(out.String())
			// fmt.Println("--- to")
			// fmt.Println((*n).RawString())
			// fmt.Println("---------------------")
			// fmt.Println("--- parent is")
			// fmt.Println((*n).Parent().RawString())
			el.RemoveAttribute(`dtemplate-process`)
		}
		return nil
	})
}