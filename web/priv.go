package web

import (
	"mtauth/db"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func NewPrivController(repos *db.Repositories) *PrivController {
	return &PrivController{repos: repos}
}

type PrivController struct {
	repos *db.Repositories
}

func (ac *PrivController) GetPrivs(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		SendError(resp, 500, err.Error())
		return
	}

	list, err := ac.repos.Priv.GetByID(id)
	if err != nil {
		SendError(resp, 500, err.Error())
	} else {
		SendJson(resp, list)
	}
}
