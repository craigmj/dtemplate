
//line src/xmlparse/xmlparser.rl:1
package xmlparse

import (
	"fmt"
)


//line src/xmlparse/xmlparser.rl:102



//line src/xmlparse/xmlparser.go:15
var _xml_actions []byte = []byte{
	0, 1, 0, 1, 10, 1, 12, 1, 14, 
	1, 15, 2, 0, 12, 2, 1, 17, 
	2, 2, 19, 2, 3, 20, 2, 4, 
	21, 2, 5, 22, 2, 6, 16, 2, 
	8, 23, 2, 9, 23, 2, 11, 24, 
	2, 12, 7, 2, 12, 10, 2, 13, 
	18, 3, 0, 12, 10, 3, 12, 7, 
	0, 3, 12, 7, 10, 4, 12, 7, 
	0, 10, 
}

var _xml_key_offsets []int16 = []int16{
	0, 0, 8, 19, 31, 43, 55, 67, 
	79, 80, 91, 103, 115, 127, 139, 151, 
	163, 164, 165, 170, 171, 172, 184, 196, 
	208, 220, 222, 223, 224, 225, 226, 227, 
	229, 230, 231, 232, 233, 234, 235, 236, 
	237, 238, 240, 251, 263, 267, 279, 291, 
	292, 293, 295, 297, 299, 299, 308, 308, 
	308, 308, 308, 308, 308, 308, 
}

var _xml_trans_keys []byte = []byte{
	45, 95, 48, 57, 65, 90, 97, 122, 
	12, 33, 39, 47, 63, 9, 10, 32, 
	34, 61, 62, 11, 13, 32, 39, 47, 
	62, 9, 12, 33, 34, 61, 63, 11, 
	13, 32, 39, 47, 62, 9, 12, 33, 
	34, 61, 63, 11, 13, 32, 39, 47, 
	61, 62, 63, 9, 12, 33, 34, 11, 
	13, 32, 39, 47, 61, 62, 63, 9, 
	12, 33, 34, 11, 13, 32, 39, 47, 
	61, 62, 63, 9, 12, 33, 34, 62, 
	11, 13, 32, 33, 34, 39, 47, 9, 
	12, 61, 63, 11, 13, 32, 39, 47, 
	62, 9, 12, 33, 34, 61, 63, 11, 
	13, 32, 39, 47, 62, 9, 12, 33, 
	34, 61, 63, 11, 13, 32, 33, 34, 
	39, 47, 62, 9, 12, 61, 63, 11, 
	13, 32, 33, 34, 39, 47, 62, 9, 
	12, 61, 63, 11, 13, 32, 33, 34, 
	39, 47, 61, 62, 63, 9, 12, 11, 
	13, 32, 33, 34, 39, 47, 61, 62, 
	63, 9, 12, 34, 34, 32, 47, 62, 
	9, 13, 39, 39, 11, 13, 32, 33, 
	34, 39, 47, 62, 9, 12, 61, 63, 
	11, 13, 32, 39, 47, 62, 9, 12, 
	33, 34, 61, 63, 11, 13, 32, 39, 
	47, 61, 62, 63, 9, 12, 33, 34, 
	11, 13, 32, 39, 47, 61, 62, 63, 
	9, 12, 33, 34, 45, 91, 62, 45, 
	45, 45, 45, 45, 62, 67, 68, 65, 
	84, 65, 91, 93, 93, 93, 62, 93, 
	11, 13, 32, 39, 47, 9, 12, 33, 
	34, 61, 63, 11, 13, 32, 39, 47, 
	62, 9, 12, 33, 34, 61, 63, 32, 
	62, 9, 13, 11, 13, 32, 39, 47, 
	62, 9, 12, 33, 34, 61, 63, 11, 
	13, 32, 39, 47, 62, 9, 12, 33, 
	34, 61, 63, 63, 63, 62, 63, 38, 
	60, 38, 60, 45, 59, 95, 48, 57, 
	65, 90, 97, 122, 
}

var _xml_single_lengths []byte = []byte{
	0, 2, 5, 6, 6, 8, 8, 8, 
	1, 7, 6, 6, 8, 8, 10, 10, 
	1, 1, 3, 1, 1, 8, 6, 8, 
	8, 2, 1, 1, 1, 1, 1, 2, 
	1, 1, 1, 1, 1, 1, 1, 1, 
	1, 2, 5, 6, 2, 6, 6, 1, 
	1, 2, 2, 2, 0, 3, 0, 0, 
	0, 0, 0, 0, 0, 0, 
}

var _xml_range_lengths []byte = []byte{
	0, 3, 3, 3, 3, 2, 2, 2, 
	0, 2, 3, 3, 2, 2, 1, 1, 
	0, 0, 1, 0, 0, 2, 3, 2, 
	2, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 3, 3, 1, 3, 3, 0, 
	0, 0, 0, 0, 0, 3, 0, 0, 
	0, 0, 0, 0, 0, 0, 
}

var _xml_index_offsets []int16 = []int16{
	0, 0, 6, 15, 25, 35, 46, 57, 
	68, 70, 80, 90, 100, 111, 122, 134, 
	146, 148, 150, 155, 157, 159, 170, 180, 
	191, 202, 205, 207, 209, 211, 213, 215, 
	218, 220, 222, 224, 226, 228, 230, 232, 
	234, 236, 239, 248, 258, 262, 272, 282, 
	284, 286, 289, 292, 295, 296, 303, 304, 
	305, 306, 307, 308, 309, 310, 
}

var _xml_indicies []byte = []byte{
	1, 1, 1, 1, 1, 0, 3, 4, 
	3, 5, 6, 3, 3, 3, 2, 9, 
	9, 8, 3, 10, 11, 8, 3, 3, 
	7, 14, 14, 13, 3, 15, 16, 13, 
	3, 3, 12, 19, 19, 18, 3, 20, 
	21, 22, 3, 18, 3, 17, 14, 14, 
	23, 3, 15, 24, 16, 3, 23, 3, 
	12, 25, 25, 18, 3, 20, 21, 22, 
	3, 18, 3, 12, 26, 3, 28, 28, 
	24, 3, 29, 30, 3, 24, 3, 27, 
	33, 33, 32, 3, 20, 22, 32, 3, 
	3, 31, 25, 25, 32, 3, 20, 22, 
	32, 3, 3, 12, 35, 35, 34, 3, 
	29, 30, 20, 22, 34, 3, 27, 37, 
	37, 36, 3, 29, 30, 15, 16, 36, 
	3, 12, 39, 39, 38, 3, 29, 30, 
	20, 21, 22, 3, 38, 12, 37, 37, 
	40, 3, 29, 30, 15, 24, 16, 3, 
	40, 12, 42, 41, 44, 43, 45, 46, 
	47, 45, 3, 42, 48, 44, 49, 39, 
	39, 34, 3, 29, 30, 20, 22, 34, 
	3, 12, 51, 51, 8, 3, 10, 11, 
	8, 3, 3, 50, 54, 54, 53, 3, 
	55, 21, 56, 3, 53, 3, 52, 57, 
	57, 53, 3, 55, 21, 56, 3, 53, 
	3, 50, 59, 60, 58, 62, 61, 63, 
	3, 65, 64, 67, 66, 68, 66, 68, 
	69, 66, 70, 3, 71, 3, 72, 3, 
	73, 3, 74, 3, 75, 3, 77, 76, 
	79, 78, 80, 78, 81, 80, 78, 83, 
	83, 5, 3, 3, 5, 3, 3, 82, 
	86, 86, 85, 3, 3, 87, 85, 3, 
	3, 84, 88, 89, 88, 3, 91, 91, 
	90, 3, 3, 87, 90, 3, 3, 82, 
	83, 83, 92, 3, 3, 89, 92, 3, 
	3, 82, 94, 93, 96, 95, 97, 96, 
	95, 99, 100, 98, 101, 101, 102, 103, 
	105, 106, 105, 105, 105, 105, 104, 104, 
	107, 108, 109, 110, 111, 112, 113, 
}

var _xml_trans_targs []byte = []byte{
	52, 53, 3, 0, 25, 42, 47, 3, 
	4, 22, 8, 56, 5, 4, 7, 8, 
	56, 5, 6, 7, 8, 9, 56, 6, 
	9, 7, 55, 10, 12, 16, 19, 10, 
	4, 11, 13, 21, 13, 14, 15, 14, 
	15, 17, 18, 17, 18, 4, 8, 56, 
	20, 20, 23, 24, 23, 6, 24, 8, 
	56, 24, 26, 27, 32, 26, 57, 28, 
	29, 30, 29, 30, 31, 58, 33, 34, 
	35, 36, 37, 38, 39, 40, 39, 40, 
	41, 59, 43, 45, 43, 44, 43, 60, 
	44, 60, 46, 45, 46, 48, 49, 48, 
	49, 61, 51, 1, 2, 50, 51, 50, 
	50, 53, 54, 50, 50, 50, 50, 50, 
	50, 50, 
}

var _xml_trans_actions []byte = []byte{
	0, 1, 1, 0, 0, 0, 0, 0, 
	41, 41, 41, 41, 1, 0, 1, 0, 
	0, 0, 44, 44, 44, 5, 44, 0, 
	0, 50, 0, 1, 1, 0, 0, 0, 
	44, 44, 44, 50, 0, 1, 44, 50, 
	0, 1, 11, 0, 5, 3, 3, 3, 
	1, 0, 1, 54, 0, 58, 58, 58, 
	58, 62, 1, 0, 0, 0, 0, 0, 
	1, 1, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 1, 1, 0, 0, 
	0, 0, 1, 1, 0, 5, 5, 5, 
	0, 0, 5, 11, 0, 1, 1, 0, 
	0, 0, 1, 0, 0, 47, 0, 29, 
	14, 0, 0, 32, 35, 23, 17, 20, 
	38, 26, 
}

var _xml_to_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 7, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 
}

var _xml_from_state_actions []byte = []byte{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 9, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 
}

var _xml_eof_trans []int16 = []int16{
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 0, 0, 0, 0, 0, 
	0, 0, 0, 102, 104, 105, 105, 108, 
	109, 110, 111, 112, 113, 114, 
}

const xml_start int = 50
const xml_first_final int = 50
const xml_error int = 0

const xml_en_main int = 50


//line src/xmlparse/xmlparser.rl:105

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

	
//line src/xmlparse/xmlparser.go:255
	{
	cs = xml_start
	ts = 0
	te = 0
	act = 0
	}

//line src/xmlparse/xmlparser.rl:127
	for p<eof {
		//fmt.Println("In p<eof loop")
		
//line src/xmlparse/xmlparser.go:267
	{
	var _klen int
	var _trans int
	var _acts int
	var _nacts uint
	var _keys int
	if p == pe {
		goto _test_eof
	}
	if cs == 0 {
		goto _out
	}
_resume:
	_acts = int(_xml_from_state_actions[cs])
	_nacts = uint(_xml_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		 _acts++
		switch _xml_actions[_acts - 1] {
		case 15:
//line NONE:1
ts = p

//line src/xmlparse/xmlparser.go:290
		}
	}

	_keys = int(_xml_key_offsets[cs])
	_trans = int(_xml_index_offsets[cs])

	_klen = int(_xml_single_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + _klen - 1)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + ((_upper - _lower) >> 1)
			switch {
			case data[p] < _xml_trans_keys[_mid]:
				_upper = _mid - 1
			case data[p] > _xml_trans_keys[_mid]:
				_lower = _mid + 1
			default:
				_trans += int(_mid - int(_keys))
				goto _match
			}
		}
		_keys += _klen
		_trans += _klen
	}

	_klen = int(_xml_range_lengths[cs])
	if _klen > 0 {
		_lower := int(_keys)
		var _mid int
		_upper := int(_keys + (_klen << 1) - 2)
		for {
			if _upper < _lower {
				break
			}

			_mid = _lower + (((_upper - _lower) >> 1) & ^1)
			switch {
			case data[p] < _xml_trans_keys[_mid]:
				_upper = _mid - 2
			case data[p] > _xml_trans_keys[_mid + 1]:
				_lower = _mid + 2
			default:
				_trans += int((_mid - int(_keys)) >> 1)
				goto _match
			}
		}
		_trans += _klen
	}

_match:
	_trans = int(_xml_indicies[_trans])
_eof_trans:
	cs = int(_xml_trans_targs[_trans])

	if _xml_trans_actions[_trans] == 0 {
		goto _again
	}

	_acts = int(_xml_trans_actions[_trans])
	_nacts = uint(_xml_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _xml_actions[_acts-1] {
		case 0:
//line src/xmlparse/xmlparser.rl:10
 _mark = p; 
		case 1:
//line src/xmlparse/xmlparser.rl:13
 
		// Handle not receiving the closing ';'
		ent := fromMark(0)
		l := len(ent)
		if ';'==ent[l-1] {
			doc.AddNode(ENTITY, fromMark(-1))
		} else {
			doc.AddNode(ENTITY, ent)
		}
	
		case 2:
//line src/xmlparse/xmlparser.rl:23
 
		doc.AddNode(COMMENT, fromMark(-3))
	
		case 3:
//line src/xmlparse/xmlparser.rl:26
 doc.AddNode(CDATA, fromMark(-3)) 
		case 4:
//line src/xmlparse/xmlparser.rl:27
 doc.AddNode(EXCLAM, fromMark(-1)) 
		case 5:
//line src/xmlparse/xmlparser.rl:28
 doc.AddNode(PI, fromMark(-2)) 
		case 6:
//line src/xmlparse/xmlparser.rl:29
 
		doc.AddNode(ENTITY, "amp")
		p--	// Step back so that we re-process the char that followed the &
	
		case 7:
//line src/xmlparse/xmlparser.rl:34
 
		el, _ := stack.Pop();
		doc.OpenElement(el);
	
		case 8:
//line src/xmlparse/xmlparser.rl:39
 
		doc.CloseElement(``); 
	
		case 9:
//line src/xmlparse/xmlparser.rl:43

		// Allows the document to auto-close void elements
		doc.EndOpenTag();
	
		case 10:
//line src/xmlparse/xmlparser.rl:48

		k, v := stack.KVPair()
		doc.SetAttribute(k,v)
	
		case 11:
//line src/xmlparse/xmlparser.rl:52

		el, _ := stack.Pop();
		doc.CloseElement(el);
	
		case 12:
//line src/xmlparse/xmlparser.rl:56

		stack.Push(fromMark(0))
	
		case 13:
//line src/xmlparse/xmlparser.rl:66
 doc.AddNode(RAW, fromMark(0)) 
		case 16:
//line src/xmlparse/xmlparser.rl:91
te = p
p--

		case 17:
//line src/xmlparse/xmlparser.rl:92
te = p
p--

		case 18:
//line src/xmlparse/xmlparser.rl:93
te = p
p--

		case 19:
//line src/xmlparse/xmlparser.rl:94
te = p
p--

		case 20:
//line src/xmlparse/xmlparser.rl:95
te = p
p--

		case 21:
//line src/xmlparse/xmlparser.rl:96
te = p
p--

		case 22:
//line src/xmlparse/xmlparser.rl:97
te = p
p--

		case 23:
//line src/xmlparse/xmlparser.rl:98
te = p
p--

		case 24:
//line src/xmlparse/xmlparser.rl:99
te = p
p--

//line src/xmlparse/xmlparser.go:477
		}
	}

_again:
	_acts = int(_xml_to_state_actions[cs])
	_nacts = uint(_xml_actions[_acts]); _acts++
	for ; _nacts > 0; _nacts-- {
		_acts++
		switch _xml_actions[_acts-1] {
		case 14:
//line NONE:1
ts = 0

//line src/xmlparse/xmlparser.go:491
		}
	}

	if cs == 0 {
		goto _out
	}
	p++
	if p != pe {
		goto _resume
	}
	_test_eof: {}
	if p == eof {
		if _xml_eof_trans[cs] > 0 {
			_trans = int(_xml_eof_trans[cs] - 1)
			goto _eof_trans
		}
	}

	_out: {}
	}

//line src/xmlparse/xmlparser.rl:130

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
