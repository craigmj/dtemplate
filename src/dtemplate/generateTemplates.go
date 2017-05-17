package dtemplate

import (
	"fmt"
	"path/filepath"

	"github.com/golang/glog"
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

	for range C {
		templates, err := loadTemplates(sourceDir)
		if nil != err {
			glog.Fatal(err)
			return err
		}
		if err := render(destDir, lang, name, templates, includeQuerySelect); nil != err {
			glog.Fatal(err)
			return err
		}
		fmt.Println("Generated templates in ", sourceDir)
	}

	return nil
}
