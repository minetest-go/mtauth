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

	var authrepo db.AuthRepository
	var privrepo db.PrivilegeRepository

	auth_backend := cfg[worldconfig.CONFIG_AUTH_BACKEND]
	switch auth_backend {
	case worldconfig.BACKEND_SQLITE3:
		authrepo, err = db.NewSQliteAuthRepository("auth.sqlite")
		if err != nil {
			panic(err)
		}
		// TODO: common db connection, standalone migration function
		privrepo, err = db.NewSQlitePrivilegeRepository("auth.sqlite")
		if err != nil {
			panic(err)
		}
		break
	default:
		panic("unsupported backend: " + auth_backend)
	}

	web.Setup(authrepo, privrepo)
	fmt.Printf("Listening on port %d\n", 8080)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
