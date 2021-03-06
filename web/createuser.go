package web

import (
	"encoding/json"
	"mtauth/auth"
	"mtauth/db"
	"net/http"
	"time"
)

type CreateUser struct {
	Name     string   `json:"name"`
	Password string   `json:"password"`
	Privs    []string `json:"privs"`
}

func NewCreateUserController(authrepo db.AuthRepository, privrepo db.PrivilegeRepository) *CreateUserController {
	return &CreateUserController{
		authrepo: authrepo,
		privrepo: privrepo,
	}
}

type CreateUserController struct {
	authrepo db.AuthRepository
	privrepo db.PrivilegeRepository
}

func (ac *CreateUserController) CreateUser(resp http.ResponseWriter, req *http.Request) {
	create_user := &CreateUser{}
	err := json.NewDecoder(req.Body).Decode(create_user)
	if err != nil {
		SendError(resp, 500, err.Error())
		return
	}

	existing_auth, err := ac.authrepo.GetByUsername(create_user.Name)
	if err != nil {
		SendError(resp, 500, err.Error())
		return
	}
	if existing_auth != nil {
		SendError(resp, 409, "auth already exists")
		return
	}

	salt, verifier, err := auth.CreateAuth(create_user.Name, create_user.Password)
	if err != nil {
		SendError(resp, 500, err.Error())
		return
	}
	dbpw := auth.CreateDBPassword(salt, verifier)

	auth_entry := db.AuthEntry{
		Name:      create_user.Name,
		Password:  dbpw,
		LastLogin: int(time.Now().UnixMilli() / 1000),
	}

	err = ac.authrepo.Create(&auth_entry)
	if err != nil {
		SendError(resp, 500, err.Error())
		return
	}

	if create_user.Privs != nil {
		for _, priv := range create_user.Privs {
			priv_entry := db.PrivilegeEntry{
				ID:        *auth_entry.ID,
				Privilege: priv,
			}
			err = ac.privrepo.Create(&priv_entry)
			if err != nil {
				SendError(resp, 500, err.Error())
				return
			}
		}
	}

	SendJson(resp, auth_entry)
}
