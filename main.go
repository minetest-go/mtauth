package main

import (
	"fmt"
	"mtauth/db"
	"mtauth/web"
	"mtauth/worldconfig"
	"net/http"
)

func main() {
	cfg, err := worldconfig.Parse("world.mt")
	if err != nil {
		panic(err)
	}

	var repo db.AuthRepository
	auth_backend := cfg[worldconfig.CONFIG_AUTH_BACKEND]
	switch auth_backend {
	case worldconfig.BACKEND_SQLITE3:
		repo, err = db.NewSQliteAuthRepository("auth.sqlite")
		if err != nil {
			panic(err)
		}
		break
	default:
		panic("unsupported backend: " + auth_backend)
	}

	web.Setup(repo)
	fmt.Printf("Listening on port %d\n", 8080)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
