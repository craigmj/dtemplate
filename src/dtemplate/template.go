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

	"gopkg.in/yaml.v2"

	`xmlparse`
	`config`
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
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if nil != err {
			return err
		}
		ext := filepath.Ext(path)
		if ".html" != ext {
			// Only consider .html files
			return nil
		}
		relPath := path[len(dir)+1 : len(path)-len(ext)]
		raw, err := ioutil.ReadFile(path)
		if nil != err {
			return err
		}

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

		if err := processNodes(node.Node, settings, cfg); nil!=err {
			return fmt.Errorf(`Failed processing nodes in %s: %s`, path, err.Error())
		}

		// With libxml2, our node is already the first-child
		// element
		t := &Template{
			Name:    strings.Replace(relPath, "/", nameSeparator, -1),
			Node:    node,
			Raw:     raw,
			Indices: findIndices(`data-set`, node),
		}
		templates = append(templates, t)

		return nil
	})
	if nil != err {
		return nil, err
	}
	return templates, nil
}

func splitMetadata(in io.Reader) ([]byte, []byte, error) {
	scan := bufio.NewScanner(in)
	var yml *bytes.Buffer
	var data *bytes.Buffer
	// Possible states are
	// 0 => looking for START OF XML, START OF YML, or YML header
	// 1 => scanning YML
	// 2 => scanning XML
	state := 0 //
	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		if "" == line {
			continue
		}
		switch state {
		case 0:
			if '<' == line[0] {
				data = bytes.NewBuffer(scan.Bytes())
				data.WriteString("\n")
				state = 2
				continue
			}
			if "---" == scan.Text() {
				// This has to be a start of yml
				yml = bytes.NewBuffer([]byte{})
				state = 1
				continue
			}
			// At this point we must have YML starting, without
			// the --- prefix, which is fine- we can manage that
			yml = bytes.NewBuffer(scan.Bytes())
			yml.WriteString("\n")
			state = 1
			continue
		case 1:
			if "---" == scan.Text() {
				data = bytes.NewBuffer([]byte{})
				state = 2
				continue
			}
			yml.Write(scan.Bytes())
			yml.WriteString("\n")
			continue
		case 2:
			data.Write(scan.Bytes())
			data.WriteString("\n")
		}
	}
	if nil == data {
		return nil, nil, fmt.Errorf(`No XML data found in file: did you finish the metadata with '---' on a separate line ?`)
	}
	if nil != yml {
		return yml.Bytes(), data.Bytes(), nil
	}
	return []byte{}, data.Bytes(), nil
}

func processNodes(node xmlparse.Node, settings map[string]interface{}, cfg *config.Config) error {
	return xmlparse.Walk(node, func(n xmlparse.Node, depth int) error {
		el, ok := n.(*xmlparse.Element)
		if !ok {
			return nil
		}
		if ``!=n.GetAttribute(`dtemplate-process`) {
			proc := n.GetAttribute(`dtemplate-process`)

			c, ok := settings[proc]
			if !ok {
				c, ok = cfg.Process[proc]
				if !ok {
					c = proc
				}
			}
			cstring, ok := c.(string)
			if !ok {
				return fmt.Errorf(`Unable to convert setting %s to a string`, proc)
			}

			args := strings.Split(strings.TrimSpace(cstring), ` `)
			cmd := exec.Command(args[0], args[1:]...)

			var out bytes.Buffer
			cmd.Stdout = &out
			rin, win, err := os.Pipe()
			if nil!=err {
				return err
			}
			cmd.Stdin = rin
			raw := el.InnerRawText()
			go func() {
				io.Copy(win, strings.NewReader(raw))
				win.Close()
			}()
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); nil!=err {
				return fmt.Errorf(`Failed running %s : %w`, strings.Join(args, ` `), err)
			}
			el.SetInnerText(out.String())
			// fmt.Println("--- Converted")
			// fmt.Println(raw)
			// fmt.Println("--- to")
			// fmt.Println(out.String())
			// fmt.Println("---------------------")
		}
		return nil
	})
}