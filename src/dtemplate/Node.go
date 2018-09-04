package dtemplate

import (
	"io"
	"io/ioutil"
	"strings"

	"xmlparse"
)

type Node struct {
	Node xmlparse.Node
}

func ParseNode(in io.Reader) (*Node, error) {
	raw, err := ioutil.ReadAll(in)
	if nil!=err {
		return nil, err
	}
	d, err := xmlparse.Parse(raw)
	if nil!=err {
		return nil, err
	}
	return &Node{d.Root}, nil
}

func (n *Node) IsElement() bool {
	if nil == n {
		panic("IsElement called on nil *Node")
	}
	if nil == n.Node {
		panic("IsElement called on node.NODE NIL")
	}
	return n.Node.Type() == xmlparse.ELEM
}

func (n *Node) FirstChild() *Node {
	c := n.Node.(*xmlparse.Element).FirstChild()

	if nil == c {
		return nil
	}
	return &Node{c}
}

func (n *Node) NextSibling() *Node {
	s := n.Node.NextSibling()
	if nil == s {
		return nil
	}
	return &Node{s}
}

func (n *Node) GetAttribute(attr string) string {
	return n.Node.(*xmlparse.Element).GetAttribute(attr)
}

func (n *Node) RemoveAttribute(attr string) {
	n.Node.(*xmlparse.Element).RemoveAttribute(attr)
}

func (n *Node) Render() string {
	if nil == n.Node {
		return "n.Node IS NIL"
	}
	return n.Node.String()
}

func (n *Node) TypescriptType() string {
	e := n.Node.(*xmlparse.Element).LocalName()
	e = strings.ToLower(e)
	t, ok := _typescript_element_map[e]
	// if e == `input` {
	// 	inputType := n.GetAttribute(`type`)
	// 	switch strings.ToLower(inputType) {
	// 	case `radio`:
	// 		return `HTMLRadioElement`
	// 	case `checkbox`:
	// 		return `HTMLCheckboxElement`
	// 	}
	// }
	if ok {
		return t
	}
	return `HTMLUnknownElement`
}

// This maps html elements to TypeScript types
var _typescript_element_map = map[string]string{
	`table`:      `HTMLTableElement`,
	`thead`:      `HTMLTableSectionElement`,
	`tbody`:      `HTMLTableSectionElement`,
	`caption`:    `HTMLTableCaptionElement`,
	`td`:         `HTMLTableCellElement`,
	`base`:       `HTMLBaseElement`,
	`a`:          `HTMLAnchorElement`,
	`area`:       `HTMLAreaElement`,
	`media`:      `HTMLMediaElement`,
	`quote`:      `HTMLQuoteElement`,
	`body`:       `HTMLBodyElement`,
	`br`:         `HTMLBRElement`,
	`button`:     `HTMLButtonElement`,
	`canvas`:     `HTMLCanvasElement`,
	`col`:        `HTMLTableColElement`,
	`datalist`:   `HTMLDataListElement`,
	`del`:        `HTMLModElement`,
	`ins`:        `HTMLModElement`,
	`div`:        `HTMLDivElement`,
	`dl`:         `HTMLDListElement`,
	`embed`:      `HTMLEmbedElement`,
	`fieldset`:   `HTMLFieldSetElement`,
	`form`:       `HTMLFormElement`,
	`h1`:         `HTMLHeadingElement`,
	`h2`:         `HTMLHeadingElement`,
	`h3`:         `HTMLHeadingElement`,
	`h4`:         `HTMLHeadingElement`,
	`h5`:         `HTMLHeadingElement`,
	`h6`:         `HTMLHeadingElement`,
	`head`:       `HTMLHeadElement`,
	`hr`:         `HTMLHRElement`,
	`html`:       `HTMLHtmlElement`,
	`iframe`:     `HTMLIFrameElement`,
	`img`:        `HTMLImageElement`,
	`input`:      `HTMLInputElement`,
	`label`:      `HTMLLabelElement`,
	`legend`:     `HTMLLegendElement`,
	`li`:         `HTMLLIElement`,
	`link`:       `HTMLLinkElement`,
	`map`:        `HTMLMapElement`,
	`menu`:       `HTMLMenuElement`,
	`meta`:       `HTMLMetaElement`,
	`object`:     `HTMLObjectElement`,
	`ol`:         `HTMLOListElement`,
	`opgroup`:    `HTMLOptGroupElement`,
	`option`:     `HTMLOptionElement`,
	`p`:          `HTMLParagraphElement`,
	`param`:      `HTMLParamElement`,
	`pre`:        `HTMLPreElement`,
	`progress`:   `HTMLProgressElement`,
	`script`:     `HTMLScriptElement`,
	`select`:     `HTMLSelectElement`,
	`source`:     `HTMLSourceElement`,
	`span`:       `HTMLSpanElement`,
	`style`:      `HTMLStyleElement`,
	`textarea`:   `HTMLTextAreaElement`,
	`title`:      `HTMLTitleElement`,
	`tr`:         `HTMLTableRowElement`,
	`track`:      `HTMLTrackElement`,
	`ul`:         `HTMLUListElement`,
	`nextid`:     `HTMLNextIdElement`,
	`applet`:     `HTMLAppletElement`,
	`address`:    `HTMLBlockElement`,
	`blockQuote`: `HTMLBlockElement`,
	`center`:     `HTMLBlockElement`,
	`listing`:    `HTMLBlockElement`,
	`plainText`:  `HTMLBlockElement`,
	`xmp`:        `HTMLBlockElement`,
	`dd`:         `HTMLDDElement`,
	`font`:       `HTMLFontElement`,
	`frame`:      `HTMLFrameElement`,
	`frameset`:   `HTMLFrameSetElement`,
	`directory`:  `HTMLDirectoryElement`,
	`phrase`:     `HTMLPhraseElement`,
	`basefont`:   `HTMLBaseFontElement`,
	`marquee`:    `HTMLMarqueeElement`,
	`dt`:         `HTMLDTElement`,
	`bgsound`:    `HTMLBGSoundElement`,
	`isindex`:    `HTMLIsIndexElement`,
	`webview`:    `MSHTMLWebViewElement`,
}
