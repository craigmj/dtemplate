package dtemplate

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func Watch(dir string, C chan bool, filter func(n string) bool) {
	w, err := fsnotify.NewWatcher()
	if nil != err {
		log.Fatal(err)
	}
	if err := w.Add(dir); nil != err {
		log.Fatal(err)
	}
	for {
		select {
		case e := <-w.Events:
			if filter(e.Name) {
				fmt.Printf("EVT %s: %s\n", e.Op.String(), e.Name)
				C <- true
			}
		case err := <-w.Errors:
			log.Fatal(err)
		}
	}
}
