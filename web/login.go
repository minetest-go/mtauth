package web

import (
	"fmt"
	"html/template"
	"mtauth/db"
	"mtauth/web/templates"
	"net/http"
)

func NewLoginController(repos *db.Repositories) *LoginController {
	return &LoginController{
		repos: repos,
		tmpl:  templates.Parse("login.html"),
	}
}

type LoginController struct {
	repos *db.Repositories
	tmpl  *template.Template
}

func (ac *LoginController) Login(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		err := req.ParseForm()
		if err != nil {
			SendError(resp, 500, err.Error())
			return
		}

		username := req.Form.Get("username")
		password := req.Form.Get("password")
		fmt.Printf("username: %s, password: %s\n", username, password)

	}

	ac.tmpl.Execute(resp, nil)
}
