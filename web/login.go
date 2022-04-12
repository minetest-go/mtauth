package web

import (
	"fmt"
	"mtauth/db"
	"mtauth/web/templates"
	"net/http"
)

func NewLoginController(authrepo db.AuthRepository, privrepo db.PrivilegeRepository) *LoginController {
	return &LoginController{
		authrepo: authrepo,
		privrepo: privrepo,
	}
}

type LoginController struct {
	authrepo db.AuthRepository
	privrepo db.PrivilegeRepository
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

	templates.Get("login.html").Execute(resp, nil)
}
