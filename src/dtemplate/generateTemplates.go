package dtemplate

import (
	"fmt"
	"os"
	"path/filepath"
	// "github.com/golang/glog"
)

func generateTemplates(sourceDir, destDir string, lang, name string, includeQuerySelect, watch bool) error {
	C := make(chan bool)
	go func() {
		defer close(C)
		C <- true // First generation
		if !watch {
			return
		}
		Watch(sourceDir, C, func(n string) bool {
			return ".html" == filepath.Ext(n)
		})
	}()

	nameSeparator := "_"
	switch lang {
	case "cljs":
		nameSeparator="-"
	}
	for range C {
		f := func() {
			// If we're watching, we catch all errors and log but ignore them.
			if watch {
				defer func() {
					if err := recover(); nil != err {
						fmt.Fprintf(os.Stderr, "ERR: %v\n", err)
					}
				}()
			}
			templates, err := loadTemplates(sourceDir, nameSeparator)
			if nil != err {
				panic(err)
			}
			if err := render(destDir, lang, name, templates, includeQuerySelect); nil != err {
				panic(err)
			}
			fmt.Println("Generated templates in ", sourceDir)
		}
		f()
	}

	return nil
}
