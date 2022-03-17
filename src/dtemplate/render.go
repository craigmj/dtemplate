package dtemplate

import (
	"encoding/json"
	"io"
	"os"
	"text/template"
	`strings`
	// "github.com/golang/glog"
)

func render(out, lang, name string, templates []*Template, includeQuerySelect, export bool) error {
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
		"Export":			  export,
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
	`js`: jsTemplate,
`cljs`:cljsTemplate,
}
