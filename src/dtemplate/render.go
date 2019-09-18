package dtemplate

import (
	"encoding/json"
	"io"
	"os"
	"text/template"
	`strings`
	// "github.com/golang/glog"
)

func render(out, lang, name string, templates []*Template, includeQuerySelect bool) error {
	var err error
	t := template.Must(template.New("").Funcs(map[string]interface{}{
		"JS": func(in interface{}) string {
			raw, err := json.Marshal(in)
			if nil != err {
				return err.Error()
			}
			return string(raw)
		},
		"clj_string": func(in string) string {
			return strings.ReplaceAll(in, `"`, `\"`)
		},
		"clj_name": func(in string) string {
			return strings.ReplaceAll(in, `_`, "-")
		},
	}).Parse(_template[lang]))
	var outf io.Writer
	if "" == out {
		outf = os.Stdout
	} else {
		outfile, err := os.Create(out)
		if nil != err {
			return err
		}
		defer outfile.Close()
		outf = outfile
	}
	if err = t.Execute(outf, map[string]interface{}{
		"Class":              name,
		"T":                  mapTemplates(templates),
		"Templates":          templates,
		"includeQuerySelect": includeQuerySelect,
	}); nil != err {
		return err
	}
	return nil
}

var _template = map[string]string{
	`ts`: `// dtemplate generated - do not edit
{{$class := .Class -}}
export namespace EL {
	{{range .Templates -}}
	export type {{.Name}} =	{{.Node.TypescriptType}};
	{{end}}
}
export namespace R {
	{{range .Templates -}}		
	export interface {{.Name}} {
		{{range .Indices}}{{if not .IsThis -}}
		{{.Name}}: {{.Node.TypescriptType}},
		{{end}}{{end -}}
	};
	{{end}}
}	// end namespace R
{{range .Templates -}}
export class {{.Name}} {
	protected static _template : {{.Node.TypescriptType}};
	public el : {{.Node.TypescriptType}};
	public $ : R.{{.Name}};
	constructor() {
		let t = {{.Name}}._template;
		if (! t ) {
			let d = document.createElement('div');
			d.innerHTML = ` + "`{{.Html}}`" + `;
			t = d.firstElementChild{{.This.Path}} as {{.Node.TypescriptType}};
			{{.Name}}._template = t;
		}
		let n = t.cloneNode(true) as {{.Node.TypescriptType}};
		{{range .Indices}}{{if .IsThis}}
		n = n{{.Path}} as {{.Node.TypescriptType}};
		{{- end}}{{end}}
		this.$ = {
		{{- range .Indices}}{{if not .IsThis}}
			{{.Name}}: n{{.Path}} as {{.Node.TypescriptType}},
		{{- end}}{{end}}
		};
		/*
		{{range .Indices}}
		{{if eq "this" .Name}}{{else}}
		if (!this.$.{{.Name}}) {
			console.error("Failed to resolve item {{.Name}} on path {{.Path}} of ", n);
			debugger;
		} else {
			console.log("{{.Name}} resolved to ", this.$.{{.Name}});
		}
		{{end}}{{end}}
		*/
		this.el = n;
	}
}
{{end}}
`,
	`js`: `// dtemplate generated - do not edit
let {{.Class}} = (function() {
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
`,
`cljs`:cljsTemplate,
}
