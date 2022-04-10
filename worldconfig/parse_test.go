package worldconfig

import (
	"fmt"
	"testing"
)

func TestParseSqlite(t *testing.T) {
	cfg := Parse("./testdata/world.mt.sqlite")
	if cfg[CONFIG_AUTH_BACKEND] != BACKEND_SQLITE3 {
		t.Fatal("not sqlite3")
	}
}

func TestParsePostgres(t *testing.T) {
	cfg := Parse("./testdata/world.mt.postgres")
	fmt.Println(cfg)
	if cfg[CONFIG_AUTH_BACKEND] != BACKEND_POSTGRES {
		t.Fatal("not postgres")
	}

	if cfg[CONFIG_PSQL_AUTH_CONNECTION] != "host=/var/run/postgresql user=postgres password=enter dbname=postgres" {
		t.Fatal("param err")
	}
}
