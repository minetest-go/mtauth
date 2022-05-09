package db

import (
	"database/sql"
	"errors"
	"mtauth/worldconfig"
)

type Repositories struct {
	Auth AuthRepository
	Priv PrivilegeRepository
}

func Setup(cfg map[string]string) (*Repositories, error) {
	repos := &Repositories{}

	var err error
	var auth_db *sql.DB
	auth_backend := cfg[worldconfig.CONFIG_AUTH_BACKEND]
	switch auth_backend {
	case worldconfig.BACKEND_SQLITE3:
		auth_db, err = sql.Open("sqlite", "auth.sqlite")
		if err != nil {
			return nil, err
		}
		err = MigrateAuth(auth_db)
		if err != nil {
			return nil, err
		}
		break
	case worldconfig.BACKEND_POSTGRES:
		connStr := cfg[worldconfig.CONFIG_PSQL_AUTH_CONNECTION]
		auth_db, err = sql.Open("postgres", connStr)
		if err != nil {
			return nil, err
		}
		err = auth_db.Ping()
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unsupported backend: " + auth_backend)
	}

	repos.Auth = NewAuthRepository(auth_db)
	repos.Priv = NewPrivilegeRepository(auth_db)

	return repos, nil
}
