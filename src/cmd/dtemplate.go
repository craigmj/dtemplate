package main

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"github.com/golang/glog"
)

func main() {
	dir := flag.String("dir", ".", "Source directory of template files")
	out := flag.String("out", "", "Output file")
	includeQuerySelect := flag.Bool("include-query-select", false, "Do not include the query selection functions in generated js")
	flag.Parse()

	templates := make(map[string]string)

	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if nil != err {
			return err
		}
		ext := filepath.Ext(path)
		if ".html" != ext {
			// Only consider .html files
			return nil
		}
		relPath := path[len(*dir)+1 : len(path)-len(ext)]
		raw, err := ioutil.ReadFile(path)
		if nil != err {
			return err
		}
		templates[relPath] = string(raw)
		return nil
	})
	if nil != err {
		glog.Fatal(err)
	}

	t := template.Must(template.New("").Funcs(map[string]interface{}{
		"JS": func(in interface{}) string {
			raw, err := json.Marshal(in)
			if nil != err {
				return err.Error()
			}
			return string(raw)
		},
	}).Parse(_template))
	var outf io.Writer
	if "" == *out {
		outf = os.Stdout
	} else {
		outfile, err := os.Create(*out)
		if nil != err {
			glog.Fatal(err)
		}
		defer outfile.Close()
		outf = outfile
	}
	if err = t.Execute(outf, map[string]interface{}{
		"T":                  templates,
		"includeQuerySelect": *includeQuerySelect,
	}); nil != err {
		glog.Fatal(err)
	}
}

var _template = `
let DTemplate = (function() {
{{if not .includeQuerySelect}}
	class joinIterators {
		constructor (iters) {
			this.iters = iters;
			this.i = 0;
		}
		next() {
			if (this.i == this.iters.length) {
				return { 'value':undefined, 'done':true }
			}
			let r = this.iters[this.i].next();
			if (!r.done) {
				return r;
			}
			this.i++;
			return this.next();
		}
		[Symbol.iterator]() {
			return this;
		}
	}

	class querySelectorAllIterator {
		constructor(qs) {
			this.qs = qs;
			this.i = 0;
		}
		next() {
			if (this.i == this.qs.length) {
				return { 'value':undefined, 'done':true }
			}
			return { 'value': this.qs.item(this.i++), 'done':false };
		}
		[Symbol.iterator]() {
			return this;
		}
	}

	let QuerySelectorAllIterate = function(el, query) {
		let els = [];
		if ('function'==typeof el.matches) {
			if (el.matches(query)) {
				els.push(el);
			}
		} else if ('function'==typeof el.matchesSelector) {
			if (el.matchesSelector(query)) {
				els.push(el);
			}
		}
		let qs = el.querySelectorAll(query);
		let i = qs[Symbol.iterator];
		if ('function'==typeof i) {
			return new joinIterators([els[Symbol.iterator](), qs[Symbol.iterator]()])
		}
		return new joinIterators([els[Symbol.iterator](), new querySelectorAllIterator(qs)]);
	}
{{- end}}	

	let templates =
		{{.T | JS}};

	let mk = function(k, html) {
		let el = document.createElement('div');
		el.innerHTML = html;

		let c = el.firstElementChild;
		while ((null!=c) && (Node.ELEMENT_NODE!=c.nodeType)) {
			c = c.nextSibling;
		}
		if (null==c) {
			console.error("FAILED TO FIND ANY ELEMENT CHILD OF ", k, ":", el)
			return mk('error', '<em>No child elements in template ' + k + '</em>');
		}
		el = c;
		let et = el.querySelector('[data-set="this"]');
		if (null!=et) {
			el = et;
			el.removeAttribute('data-set');
		}
		return el;
	}

	for (let i in templates) {
		templates[i] = mk(i, templates[i]);
	}

	return function(t, dest={}) {
		// Return a deep copy of the node
		let n = templates[t].cloneNode(true);
		try {
			for (let el of QuerySelectorAllIterate(n, '[data-set]')) {
				let a = el.getAttribute('data-set');
				if (a.substr(0,1)=='$') {
					a = a.substr(1);
					el = jQuery(el);
				}
				dest[a] = el;
			}
		} catch (err) {
			console.error("ERROR in DTemplate(" + t + "): ", err);
			debugger;
		}
		return [n,dest];
	}
})();
`
