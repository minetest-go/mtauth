package db

import (
	"database/sql"
	"embed"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

//go:embed migrations/*.sql
var migrations embed.FS

func init() {
	source.Register("embed", &driver{})
}

type driver struct {
	httpfs.PartialDriver
}

func (d *driver) Open(rawURL string) (source.Driver, error) {
	err := d.PartialDriver.Init(http.FS(migrations), "migrations")
	if err != nil {
		return nil, err
	}

	return d, nil
}

func MigrateMTAuth(db *sql.DB) (uint, error) {
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		return 0, err
	}

	m, err := migrate.NewWithDatabaseInstance("embed://", "sqlite", driver)
	if err != nil {
		return 0, err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return 0, err
	}

	v, _, _ := m.Version()
	return v, nil
}
