package web

import (
	"fmt"
	"html/template"
	"mtauth/db"
	"mtauth/web/templates"
	"net/http"
)

func NewLoginController(authrepo db.AuthRepository, privrepo db.PrivilegeRepository) *LoginController {
	return &LoginController{
		authrepo: authrepo,
		privrepo: privrepo,
		tmpl:     templates.Parse("login.html"),
	}
}

type LoginController struct {
	authrepo db.AuthRepository
	privrepo db.PrivilegeRepository
	tmpl     *template.Template
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
