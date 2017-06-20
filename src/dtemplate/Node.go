package dtemplate

import (
	"io"
	"strings"

	// "github.com/lestrrat/go-libxml2"
	"github.com/lestrrat/go-libxml2/clib"
	"github.com/lestrrat/go-libxml2/parser"
	// "github.com/lestrrat/go-libxml2/dom"
	"github.com/lestrrat/go-libxml2/types"
)

type Node struct {
	Node types.Node
}

func ParseNode(in io.Reader) (*Node, error) {
	parser := parser.New(
		// parser.XMLParseRecover,
		//parser.XMLParseNoEnt,
		// parser.XMLParseCompact,  // default
		parser.XMLParseNoBlanks, // default
		// parser.XMLParseNoError,  // default
		// parser.XMLParseNoWarning, // default
	)
	d, err := parser.ParseReader(in)
	if err != nil {
		return nil, err
	}
	n, err := d.DocumentElement()
	return &Node{n}, err
}

func (n *Node) IsElement() bool {
	if nil == n {
		panic("IsElement called on nil *Node")
	}
	if nil == n.Node {
		panic("IsElement called on node.NODE NIL")
	}
	return n.Node.NodeType() == clib.ElementNode
}

func (n *Node) FirstChild() *Node {
	c, err := n.Node.FirstChild()
	if nil != err {
		return nil
	}
	if nil == c {
		return nil
	}
	return &Node{c}
}

func (n *Node) NextSibling() *Node {
	s, err := n.Node.NextSibling()
	if nil != err {
		return nil
	}
	if nil == s {
		return nil
	}
	return &Node{s}
}

func (n *Node) GetAttribute(attr string) string {
	a, err := n.Node.(types.Element).GetAttribute(attr)
	if nil != err {
		return ``
	}
	return a.Value()
}

func (n *Node) RemoveAttribute(attr string) {
	n.Node.(types.Element).RemoveAttribute(attr)
}

func (n *Node) Render() string {
	if nil == n.Node {
		return "n.Node IS NIL"
	}
	return n.Node.String()
}

func (n *Node) TypescriptType() string {
	e := n.Node.(types.Element).LocalName()
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
