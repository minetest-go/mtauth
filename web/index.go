package web

import (
	"mtauth/web/templates"
	"net/http"
)

var indexTmpl = templates.Parse("index.html")

func Index(resp http.ResponseWriter, req *http.Request) {
	indexTmpl.Execute(resp, nil)
}
