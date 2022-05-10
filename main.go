package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"mtauth/db"
	"mtauth/web"
	"mtauth/worldconfig"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

func main() {
	cfg, err := worldconfig.Parse("world.mt")
	if err != nil {
		panic(err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	repos, err := db.Setup(cwd, cfg)
	if err != nil {
		panic(err)
	}

	shared_secret_entry, err := repos.Settings.GetByKey(db.SETTING_SHARED_SECRET)
	if err != nil {
		panic(err)
	}
	if shared_secret_entry == nil {
		key := make([]byte, 32)
		_, err = rand.Read(key)
		if err != nil {
			panic(err)
		}
		shared_secret_entry = &db.Setting{
			Key:   db.SETTING_SHARED_SECRET,
			Value: base64.StdEncoding.EncodeToString(key),
		}

	}

	web.Setup(repos)
	fmt.Printf("Listening on port %d\n", 8080)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
