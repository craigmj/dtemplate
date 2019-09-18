package dtemplate

var cljsTemplate = `
{{- $class := .Class -}}
(ns {{.Class}})

(defn- find-child [el path]
	(if (empty? path) el
	  (recur (.item (.-childNodes el) (first path)) (rest path))))

;; generated dtemplate file - do not edit
{{range .Templates -}}
{{$t := .}}
(defn {{.Name}} []
	(let [div (.createElement js/document "div")]
		(set! (.-innerHTML div) "{{.Html | clj_string}}")
		(let [el (find-child (.-firstElementChild div) (list {{.This.CljPath}}))]
			{
				:el el
				{{range .Indices}}{{if not .IsThis -}}
				:{{.Name}} (find-child el (list {{.CljPath}}))
				{{end}}{{end -}}
			})))
{{range .Indices}}{{if not .IsThis -}}
(defn {{$t.Name}}-{{.Name}}! [el txt]
	(set! (.-textContent (:{{.Name}} el)) txt) el)
(defn {{$t.Name}}-conj-{{.Name}}! [el child]
	(.appendChild (:{{.Name}} el) child) el)
{{end}}{{end}}{{end}}
`