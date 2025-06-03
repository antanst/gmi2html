package gmi2html

import "html/template"

// Templates for different line types

var (
	textLineTmpl          = template.Must(template.New("textLine").Parse(`<p class="gemini-textline">{{.}}</p>`))
	h1Tmpl                = template.Must(template.New("h1").Parse(`<h1 class="gemini-heading-1">{{.}}</h1>`))
	h2Tmpl                = template.Must(template.New("h2").Parse(`<h2 class="gemini-heading-2">{{.}}</h2>`))
	h3Tmpl                = template.Must(template.New("h3").Parse(`<h3 class="gemini-heading-3">{{.}}</h3>`))
	listItemTmpl          = template.Must(template.New("listItem").Parse(`<p class="gemini-list-item">• {{.}}</p>`))
	blockquoteTmpl        = template.Must(template.New("blockquote").Parse(`<blockquote class="gemini-blockquote">{{.}}</blockquote>`))
	preformattedTmplStart = template.Must(template.New("preformatted").Parse(`<pre class="gemini-preformatted">`))
	preformattedTmplEnd   = template.Must(template.New("preformatted").Parse(`</pre>`))
	linkTmpl              = template.Must(template.New("link").Parse(`<div class="gemini-link-container"><a href="{{.URL}}">{{.Description}}</a></div>`))
)
