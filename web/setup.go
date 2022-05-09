package web

import (
	"mtauth/db"
	"mtauth/web/static"
	"net/http"

	"github.com/gorilla/mux"
)

func Setup(repos *db.Repositories) {
	auth_controller := NewAuthController(repos)
	priv_controller := NewPrivController(repos)
	createuser_controller := NewCreateUserController(repos)
	login_controller := NewLoginController(repos)
	oauth_controller := NewOAuthController()

	r := mux.NewRouter()
	// api surface
	r.HandleFunc("/api/auth/{username}", auth_controller.GetAuth)
	r.HandleFunc("/api/auth/{username}/verify", auth_controller.Verify).Methods(http.MethodPost)
	r.HandleFunc("/api/createuser", createuser_controller.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/api/user_privileges/{id}", priv_controller.GetPrivs)
	r.HandleFunc("/oauth/authorize", oauth_controller.Authorize)
	// forms
	r.HandleFunc("/", Index)
	r.HandleFunc("/login", login_controller.Login)
	// static files
	r.PathPrefix("/").HandlerFunc(http.FileServer(http.FS(static.Files)).ServeHTTP)

	http.Handle("/", r)
}
