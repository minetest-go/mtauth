package templates

import (
	"embed"
	"html/template"
)

//go:embed *.html
var files embed.FS
var t *template.Template

func Init() error {
	var err error
	t, err = template.ParseFS(files, "*.html")
	return err
}

func Get(name string) *template.Template {
	return t.Lookup(name)
}
