package dtemplate

import (
	"log"
	"os"
	"path/filepath"

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
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if nil != err {
			return err
		}
		if info.IsDir() {
			if err := w.Add(path); nil != err {
				log.Fatalf(`Failed to add dir %s`, path)
			}
		}
		return nil
	}); nil != err {
		log.Fatal(err)
	}
	for {
		select {
		case e := <-w.Events:
			// fmt.Printf("EVT %s: %s\n", e.Op.String(), e.Name)
			if filter(e.Name) {
				C <- true
			}
			switch e.Op {
			case fsnotify.Create:
				info, err := os.Lstat(e.Name)
				if nil != err {
					log.Fatal(err)
				}
				if info.IsDir() {
					w.Add(e.Name)
				}
			}
		case err := <-w.Errors:
			log.Fatal(err)
		}
	}
}
