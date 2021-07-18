package xmlparse

import (
	"fmt"
)

%%{
	machine xml;

	action mark { _mark = p; }
	action push { stack.Push(fromMark(0)); }

	action endEntity { 
		// Handle not receiving the closing ';'
		ent := fromMark(0)
		l := len(ent)
		if ';'==ent[l-1] {
			doc.AddNode(ENTITY, fromMark(-1))
		} else {
			doc.AddNode(ENTITY, ent)
		}
	}
	action endComment { 
		doc.AddNode(COMMENT, fromMark(-3))
	}
	action endCDATA { doc.AddNode(CDATA, fromMark(-3)) }
	action endExclam { doc.AddNode(EXCLAM, fromMark(-1)) }
	action endPI { doc.AddNode(PI, fromMark(-2)) }
	action endAmpersand { 
		doc.AddNode(ENTITY, "amp")
		p--	// Step back so that we re-process the char that followed the &
	}

	action endElemName { 
		el, _ := stack.Pop();
		doc.OpenElement(el);
	}

	action autoCloseTag { 
		doc.CloseElement(``); 
	}

	action openTag {
		// Allows the document to auto-close void elements
		doc.EndOpenTag();
	}

	action setAttribute {
		k, v := stack.KVPair()
		doc.SetAttribute(k,v)
	}
	action endTag {
		el, _ := stack.Pop();
		doc.CloseElement(el);
	}
	action pushString {
		stack.Push(fromMark(0))
	}

	ELEM = ((any - "\t" - "\n" - "\f" - " " - "/" - ">" - "\"" - "'" - "=" - "!" - "?")+) >mark %pushString;

	STRINGDELIM = "\"" (any - "\"")* >mark %pushString "\"";
	STRINGAPO = "'" (any - "'")* >mark %pushString "'";
	STRING = STRINGDELIM | STRINGAPO | ELEM;

	RAW = (any - "&" - "<")+ >mark %{ doc.AddNode(RAW, fromMark(0)) };

	ENTITYCHAR = [a-zA-Z0-9_] | "-";

	ENTITY = "&" (ENTITYCHAR+) >mark ";"? %endEntity ;
	AMPERSAND = "&" (any - ENTITYCHAR) %endAmpersand;

	ENDCOMMENT = "-->";

	COMMENT = "<!--" (any* -- ENDCOMMENT) >mark ENDCOMMENT %endComment;
	CDATA = "<![CDATA[" (any* -- "]]>") >mark "]]>" %endCDATA ;
	PI = "<?" (any* -- "?>") >mark "?>" %endPI ;
	EXCLAM = "<!" (any-"[" - "-") >mark (any - ">")* ">" %endExclam;

	ATTRIBVALUE = space* "=" space* STRING;
	ATTRIB = space+ ELEM ATTRIBVALUE? %setAttribute;

	CLOSETAG = "/>" %autoCloseTag;
	OPENTAG = ">" %openTag;

	TAG = "<" ELEM >mark %endElemName ATTRIB* space* (CLOSETAG | OPENTAG);

	ETAG = "</" space* ELEM space* ">" %endTag;

	main := |*
		AMPERSAND => {};
		ENTITY => {};
		RAW => {};
		COMMENT => {};
		CDATA => {};
		EXCLAM => {};
		PI => {};
		TAG => {};
		ETAG => {};
		*|;

}%%

%% write data;

func Parse(data []byte) (*Document, error) {
	doc := NewDocument()
	// p is our current position
	p,cs := 0,0
	pe := len(data)
	eof := pe
	ts, te, act := 0,0,0
	if 0<ts || 0<te || 0<act {}
	_mark := 0


	// Stack is used for storing strings
	stack := NewStack()

	fromMark := func(delta int) string {
		return string(data[_mark:p+delta])
	}

	loop := 0

	%% write init;
	for p<eof {
		//fmt.Println("In p<eof loop")
		%% write exec;

		loop++
		if loop>100 {
			f := p+100
			if f>len(data) {
				f = len(data)
			}
			return nil, fmt.Errorf("We seem to have got stuck around character position %d: %s", p, string(data[p:f]))
		}
	}

	return doc, nil
}
