package templates

import (
	"embed"
	"html/template"
)

//go:embed *.html
var files embed.FS

func Parse(name string) *template.Template {
	return template.Must(template.ParseFS(files, "layout.html", name))
}
