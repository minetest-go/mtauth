package web

import (
	"mtauth/db"
	"net/http"

	"github.com/gorilla/mux"
)

func Setup(authrepo db.AuthRepository, privrepo db.PrivilegeRepository) {
	auth_controller := NewAuthController(authrepo)
	priv_controller := NewPrivController(privrepo)
	createuser_controller := NewCreateUserController(authrepo, privrepo)

	r := mux.NewRouter()
	r.HandleFunc("/api/auth/{username}", auth_controller.GetAuth)
	r.HandleFunc("/api/auth/{username}/verify", auth_controller.Verify).Methods(http.MethodPost)
	r.HandleFunc("/api/createuser", createuser_controller.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/api/user_privileges/{id}", priv_controller.GetPrivs)

	http.Handle("/", r)
}
