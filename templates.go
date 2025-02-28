package gmi2html

import "html/template"

// Templates for different line types

var (
	textLineTmpl     = template.Must(template.New("textLine").Parse(`<p class="gemini-textline">{{.}}</p>`))
	h1Tmpl           = template.Must(template.New("h1").Parse(`<h1 class="gemini-heading-1">{{.}}</h1>`))
	h2Tmpl           = template.Must(template.New("h2").Parse(`<h2 class="gemini-heading-2">{{.}}</h2>`))
	h3Tmpl           = template.Must(template.New("h3").Parse(`<h3 class="gemini-heading-3">{{.}}</h3>`))
	listItemTmpl     = template.Must(template.New("listItem").Parse(`<p class="gemini-list-item">• {{.}}</p>`))
	blockquoteTmpl   = template.Must(template.New("blockquote").Parse(`<blockquote class="gemini-blockquote">{{.}}</blockquote>`))
	preformattedTmpl = template.Must(template.New("preformatted").Parse(`<pre class="gemini-preformatted">{{.}}</pre>`))
	linkTmpl         = template.Must(template.New("link").Parse(`<div class="gemini-link-container"><a href="{{.URL}}">{{.Description}}</a></div>`))

	// Container template with typography-focused CSS
	containerTmpl = template.Must(template.New("container").Parse(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{.Title}}</title>
    <style>
        :root {
            --text-color: #333;
            --bg-color: #fff;
            --link-color: #0066cc;
            --link-hover: #004080;
            --quote-bg: #f5f5f5;
            --quote-border: #ddd;
            --pre-bg: #f8f8f8;
            --pre-border: #eaeaea;
        }
        
        @media (prefers-color-scheme: dark) {
            :root {
                --text-color: #eee;
                --bg-color: #292929;
                --link-color: #4a9eff;
                --link-hover: #77b6ff;
                --quote-bg: #333;
                --quote-border: #444;
                --pre-bg: #2a2a2a;
                --pre-border: #3a3a3a;
            }
        }
        
body {
    font-family: "Source Serif Pro", "Georgia", "Cambria", serif;
    color: var(--text-color);
    /* background-color: var(--bg-color); */
    max-width: 34rem;
    margin: 0 auto;
    padding: 1rem 1rem;
    font-size: 16px;
    text-align: justify;
    hyphens: auto;
    -webkit-hyphens: auto;
    -ms-hyphens: auto;
}
        
        .gemini-container {
            width: 100%;
        }
        
        .gemini-heading-1 {
            font-size: 2rem;
            margin-top: 1rem;
            margin-bottom: 1rem;
            font-weight: 700;
            line-height: 1.2;
        }
        
        .gemini-heading-2 {
            font-size: 1.6rem;
            margin-top: 0.8rem;
            margin-bottom: 0.8rem;
            font-weight: 600;
            line-height: 1.3;
        }
        
        .gemini-heading-3 {
            font-size: 1.3rem;
            margin-top: 0.7rem;
            margin-bottom: 0.7rem;
            font-weight: 600;
            line-height: 1.4;
        }
        
        .gemini-textline {
            margin-bottom: 1rem;
        }
        
        .gemini-list-item {
            margin: 0.5rem 0;
            padding-left: 0.5rem;
        }
        
        .gemini-link-container {
            margin: 1rem 0;
        }
        
        .gemini-link-container a {
            color: var(--link-color);
            text-decoration: none;
            border-bottom: 1px solid transparent;
            transition: border-color 0.2s ease;
        }
        
        .gemini-link-container a:hover {
            color: var(--link-hover);
            border-bottom-color: var(--link-hover);
        }
        
        .gemini-blockquote {
            background-color: var(--quote-bg);
            border-left: 3px solid var(--quote-border);
            margin: 1.5rem 0;
            padding: 0.8rem 1rem;
            font-style: italic;
        }
        
        .gemini-preformatted {
            background-color: var(--pre-bg);
            border: 1px solid var(--pre-border);
            border-radius: 3px;
            padding: 1rem;
            overflow-x: auto;
            font-family: monospace;
            font-size: 0.9rem;
            margin: 1rem 0;
            white-space: pre-wrap;
        }
    </style>
</head>
<body>
    <div class="gemini-container">
        {{.Content}}
    </div>
</body>
</html>`))
)
