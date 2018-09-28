package xmlparse

import (
	"bytes"
	"fmt"
)

type NodeType uint32

const (
	RAW NodeType = 1 << iota
	ENTITY
	EXCLAM
	CDATA
	COMMENT
	PI
	ELEM
	DOCUMENT
)

type Node interface {
	NextSibling() Node
	Parent() *Element
	String() string
	SetNextSibling(nextsibling Node)
	Type() NodeType
	GetAttribute(k string) string
}

type base struct {
	nodeType NodeType
	nextsibling Node
	parent *Element
}

func (b base) Type() NodeType {
	return b.nodeType
}
func (b base) NextSibling() Node {
	return b.nextsibling
}
func (b base) Parent() *Element {
	return b.parent
}
func (b *base) SetNextSibling(n Node) {
	if nil==b {
		return
	}
	b.nextsibling = n
}

type RawNode struct {
	base
	content  string
}

type Element struct {
	base
	tag        string
	attributes map[string]string
	children   []Node
}

func (r *RawNode) GetAttribute(k string) string {
	panic("Attempt to call GetAttribute on a non-element node")
}
func (e *Element) GetAttribute(k string) string {
	return e.attributes[k]
}
func (e *Element) RemoveAttribute(k string) {
	delete(e.attributes, k)
}
func (e *Element) LocalName() string {
	return e.tag
}

func (r *RawNode) String() string {
	switch r.nodeType {
	case RAW:
		return r.content
	case ENTITY:
		return fmt.Sprintf("&%s;", r.content)
	case CDATA:
		return fmt.Sprintf("<![CDATA[%s]]>", r.content)
	case COMMENT:
		return fmt.Sprintf("<!--%s-->", r.content)
	case PI:
		return fmt.Sprintf("<?%s?>", r.content)
	case EXCLAM:
		return fmt.Sprintf("<!%s>", r.content)
	}
	return fmt.Sprintf("UNRECOGNIZED TYPE %d", r.nodeType)
}

func (e *Element) FirstChild() Node {
	if 0==len(e.children) {
		return nil
	}

	return e.children[0]
}

func (e *Element) LastChild() Node {
	if 0==len(e.children) {
		return nil
	}
	return e.children[len(e.children)-1]
}

func (e *Element) String() string {
	var out bytes.Buffer
	if `` == e.tag { // THIS IS THE DOCUMENT
		for _, c := range e.children {
			fmt.Fprintf(&out, c.String())
		}
		return out.String()
	}
	fmt.Fprintf(&out, "<%s", e.tag)
	for k, v := range e.attributes {
		fmt.Fprintf(&out, " %s=\"%s\"", k, v)
	}
	// You cannot have an empty div - it causes all sorts of grief...
	if 0 == len(e.children) && `div` != e.tag {
		fmt.Fprintf(&out, "/>")
	} else {
		fmt.Fprintf(&out, ">")
		if 0 == len(e.children) {
			// if an element must be non-empty, we have to
			// add some 'random' content to it here.
			fmt.Fprintf(&out, ` `)
		} else {
			for _, c := range e.children {
				fmt.Fprintf(&out, c.String())
			}
		}
		fmt.Fprintf(&out, "</%s>", e.tag)
	}
	return out.String()
}

type Document struct {
	Root    *Element
	Current *Element
}

func NewDocument() *Document {
	d := &Document{
		Root: &Element{
			base: base{ 
				nodeType: DOCUMENT,
				nextsibling: nil,
				parent: nil,
			},
			tag:      ``,
			children: []Node{},
		},
	}
	d.Current = d.Root
	return d
}

func (d *Document) AddNode(nt NodeType, content string) {
	b := base{
		nodeType: nt,
		nextsibling: nil,
		parent: d.Current,
	}

	node := &RawNode{base: b, content: content}
	sib := d.Current.LastChild()
	if nil!=sib {
		d.Current.LastChild().SetNextSibling(node)
	}
	d.Current.children = append(d.Current.children, node)
}

func (d *Document) OpenElement(tag string) {
	b := base{
		nodeType: ELEM,
		nextsibling: nil,
		parent: d.Current,
	}
	el := &Element{
		base: b,
		tag:        tag,
		attributes: map[string]string{},
		children:   []Node{},
	}
	sib := d.Current.LastChild()
	if nil!=sib {
		d.Current.LastChild().SetNextSibling(el)
	}
	d.Current.children = append(d.Current.children, el)
	d.Current = el
}

func (d *Document) SetAttribute(k, v string) {
	// Could check whether we have duplicated an attribute here...
	d.Current.attributes[k] = v
}

func (d *Document) CloseElement(tag string) {
	// TODO: Need to check that our current element is indeed the tagged element we are closing.
	// Then move up the tree until we find the proper tag to close
	if `` != tag && d.Current.tag != tag {
		find := d.Current
		for ; find != nil && find.tag != tag; find = find.parent {
		}
		if nil == find {
			// Closing a tag we haven't found in the tree - just ignore
		} else {
			d.Current = find.parent // Close the closest parent-matching node we have
		}
	} else {
		d.Current = d.Current.parent
	}
}

// EndOpenTag is triggered when a tag open reaches it's end. We use this as an opportunity
// to auto-close void tags.
func (d *Document) EndOpenTag() {
	tag := d.Current.tag
	for _, v := range []string{
		"area", "base",
		"br", "col", "command", "embed", "hr", "img", "input", "keygen",
		"link", "meta", "param", "source", "track", "wbr"} {
		if tag == v {
			d.Current = d.Current.parent
			return
		}
	}
}

func (d *Document) String() string {
	return d.Root.String()
}