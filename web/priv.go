package web

import (
	"mtauth/db"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func NewPrivController(repo db.PrivilegeRepository) *PrivController {
	return &PrivController{repo: repo}
}

type PrivController struct {
	repo db.PrivilegeRepository
}

func (ac *PrivController) GetPrivs(resp http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		SendError(resp, 500, err.Error())
	}

	list, err := ac.repo.GetByID(id)
	if err != nil {
		SendError(resp, 500, err.Error())
	}
	SendJson(resp, list)
}
