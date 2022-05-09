package main

import (
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

	repos, err := db.Setup(cfg)
	if err != nil {
		panic(err)
	}

	web.Setup(repos)
	fmt.Printf("Listening on port %d\n", 8080)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
