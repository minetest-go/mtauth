package web

import (
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
	templates.Get("login.html").Execute(resp, true)
}
