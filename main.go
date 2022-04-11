package main

import (
	"database/sql"
	"fmt"
	"mtauth/db"
	"mtauth/web"
	"mtauth/worldconfig"
	"net/http"

	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

func main() {
	cfg, err := worldconfig.Parse("world.mt")
	if err != nil {
		panic(err)
	}

	var db_ *sql.DB
	auth_backend := cfg[worldconfig.CONFIG_AUTH_BACKEND]
	switch auth_backend {
	case worldconfig.BACKEND_SQLITE3:
		db_, err = sql.Open("sqlite", "auth.sqlite")
		if err != nil {
			panic(err)
		}
		err = db.Migrate(db_)
		if err != nil {
			panic(err)
		}
		break
	case worldconfig.BACKEND_POSTGRES:
		connStr := cfg[worldconfig.CONFIG_PSQL_AUTH_CONNECTION]
		db_, err = sql.Open("postgres", connStr)
		if err != nil {
			panic(err)
		}
		err = db_.Ping()
		if err != nil {
			panic(err)
		}
	default:
		panic("unsupported backend: " + auth_backend)
	}

	authrepo := db.NewAuthRepository(db_)
	privrepo := db.NewPrivilegeRepository(db_)
	web.Setup(authrepo, privrepo)
	fmt.Printf("Listening on port %d\n", 8080)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
