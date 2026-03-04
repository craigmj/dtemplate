package dtemplate

import (
	"flag"
	"os"
	"runtime"

	"github.com/golang/glog"
)

func Main() {
	args := flag.NewFlagSet("dtemplate", flag.ExitOnError)
	dir := args.String("dir", ".", "Source directory of template files")
	out := args.String("out", "", "Output file")
	name := args.String(`name`, `DTemplate`, `Name for the generate class/function`)
	watch := args.Bool(`watch`, false, `Watch and recompile on changes`)
	includeQuerySelect := args.Bool("include-query-select", false, "Include the query selection functions in generated js")
	lang := args.String(`lang`, `js`, `Language: js|ts|cljs`)
	config := args.String(`config`, `dtemplate.yml`, `Configuration file for dtemplate`)
	export := args.Bool(`export`, true, `Include JS modules export on main DTemplate function`)
	pathSeparator := args.String(`separator`, `_`, `Separator for templates in subdirectories`)
	if `android` == runtime.GOOS {
		args.Parse(os.Args[2:])
	} else {
		args.Parse(os.Args)
	}

	/*
		ilog.Printf("os.Args[%d] = %v", len(os.Args), os.Args)
		ilog.Printf("os.Args[2:]= %v", os.Args[2:])
		ilog.Printf("runtime.GOOS = %s", runtime.GOOS)
		ilog.Fatalf("name='%s', pathSeparator = %s, args=%v\n", *name, *pathSeparator, args.Args())
	*/

	if err := generateTemplates(*config, *dir, *out, *lang, *name, *includeQuerySelect, *watch, *export, *pathSeparator); nil != err {
		glog.Fatal(err)
	}

}
