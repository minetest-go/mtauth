package web

import (
	"io/ioutil"
	"mtauth/auth"
	"mtauth/db"
	"net/http"

	"github.com/gorilla/mux"
)

func NewAuthController(repo db.AuthRepository) *AuthController {
	return &AuthController{repo: repo}
}

type AuthController struct {
	repo db.AuthRepository
}

func (ac *AuthController) GetAuth(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	username := vars["username"]
	entry, err := ac.repo.GetByUsername(username)
	if err != nil {
		SendError(resp, 500, err.Error())
	}
	if entry != nil {
		SendJson(resp, entry)
	} else {
		SendError(resp, 404, "not found")
	}
}

func (ac *AuthController) Verify(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	username := vars["username"]
	entry, err := ac.repo.GetByUsername(username)
	if err != nil {
		SendError(resp, 500, err.Error())
		return
	}
	if entry == nil {
		SendError(resp, 404, "not found")
		return
	}

	salt, verifier, err := auth.ParseDBPassword(entry.Password)
	if err != nil {
		SendError(resp, 500, err.Error())
		return
	}

	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		SendError(resp, 500, err.Error())
		return
	}

	ok, err := auth.VerifyAuth(username, string(data), salt, verifier)
	if err != nil {
		SendError(resp, 500, err.Error())
		return
	}
	if ok {
		resp.WriteHeader(http.StatusOK)
	} else {
		resp.WriteHeader(http.StatusUnauthorized)
	}
}
