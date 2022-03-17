package dtemplate

import (
	"flag"

	"github.com/golang/glog"
)

func Main() {
	dir := flag.String("dir", ".", "Source directory of template files")
	out := flag.String("out", "", "Output file")
	name := flag.String(`name`, `DTemplate`, `Name for the generate class/function`)
	watch := flag.Bool(`watch`, false, `Watch and recompile on changes`)
	includeQuerySelect := flag.Bool("include-query-select", false, "Do not include the query selection functions in generated js")
	lang := flag.String(`lang`, `js`, `Language: js|ts|cljs`)
	config := flag.String(`config`, `dtemplate.yml`, `Configuration file for dtemplate`)
	export := flag.Bool(`export`, true, `Include JS modules export on main DTemplate function`)
	pathSeparator := flag.String(`separator`, `_`, `Separator for templates in subdirectories`)
	flag.Parse()

	if err := generateTemplates(*config, *dir, *out, *lang, *name, *includeQuerySelect, *watch, *export, *pathSeparator); nil != err {
		glog.Fatal(err)
	}

}
