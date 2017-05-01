package dtemplate

// import (
// 	"sort"
// 	"bytes"
// 	"fmt"
// 	"regexp"
// 	"strings"
// )

// var _spaceOrComma = regexp.MustCompile(`\s+|,`)

// type eventNameFunctionMap struct {
// 	EventName   string
// 	FunctionKey string
// }

// type eventifyMap struct {
// 	Events   []*eventNameFunctionMap
// 	Elements []*MapNode
// }

// func (em *eventifyMap) Keys() []string {
// 	keyMap := map[string]bool{}
// 	for _, e := range em.Events {
// 		keyMap[e.EventName] = true
// 	}
// 	keys := make([]string, 0, len(keyMap))
// 	for k, _ := range keyMap{
// 		keys = append(keys, k)
// 	}
// 	sort.Sort(sort.StringSlice(keys))
// 	return keys
// }

// func Eventify(lang string, am AttributeMap) (*eventifyMap, error) {
// 	var out bytes.Buffer
// 	eventMap := []*eventifyMap{}

// 	for val, mapNode := range am {
// 		evtPairs := strings.Split(val, ` `)
// 		for _, evtPair := range evtPairs {
// 			evtArray := string.Split(evtPair, `:`)
// 			var event, key string
// 			event = evtArray[0]
// 			if 1 == len(evtArray) {
// 				key = event
// 			} else {
// 				key = evtArray[1]
// 			}
// 			if 1 < len(evtArray) {
// 				return ``, fmt.Errorf(
// 					`Event description appears to have too many :'s in %s - (%s)`,
// 					evtPair, val)
// 			}

// 		}
// 	}
// }

// var _eventify_typescript_script = `
// let eventify = function(events:Map<string,(e:Event)=>void>):void {
// 	{{$parentNode = .Parent}};
// 	let el: HTMLElement;
// 	{{range .Eventify}}{{$events = .Events}}{{range .Elements}}
// 	el = {{$parentNode}}{{.}};
// 	{{range $events}}
// 	el.addEventListener('{{.EventName}}', events.get('{{.FunctionKey}}'));
// 	{{end}}
// 	{{end}}{{end}}
// };
// `
// var _eventify_typescript_interface = `
// interface Events{{.Name}} {
// 	{{range .Eventify}}

// }
// `
