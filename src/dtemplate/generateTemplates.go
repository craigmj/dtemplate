package dtemplate

import (
	"fmt"
	"os"
	"path/filepath"
	// "github.com/golang/glog"

	`dtemplate/config`
)

func generateTemplates(configFilename, sourceDir, destDir string, lang, name string, includeQuerySelect, watch, export bool, pathSeparator string) error {
	cfg, err := config.ReadConfig(configFilename)
	if nil!=err {
		return err
	}
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

	nameSeparator := pathSeparator
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
			templates, err := loadTemplates(sourceDir, nameSeparator, cfg)
			if nil != err {
				panic(err)
			}
			if err := render(destDir, lang, name, templates, includeQuerySelect, export); nil != err {
				panic(err)
			}
			fmt.Println("Generated templates in ", sourceDir)
		}
		f()
	}

	return nil
}
