package web

import (
	"mtauth/db"
	"net/http"

	"github.com/gorilla/mux"
)

func Setup(repo db.AuthRepository) {
	controller := NewAuthController(repo)

	r := mux.NewRouter()
	r.HandleFunc("/api/auth/{username}", controller.GetAuth)
	r.HandleFunc("/api/auth/{username}/verify", controller.Verify).Methods(http.MethodPost)

	http.Handle("/", r)
}
