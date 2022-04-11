package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgresAuthRepository(connStr string) (*PostgresAuthRepository, error) {
	db, err := sql.Open("postgres", connStr)
	return &PostgresAuthRepository{db: db}, err
}

type PostgresAuthRepository struct {
	db *sql.DB
}

func (repo *PostgresAuthRepository) GetByUsername(username string) (*AuthEntry, error) {
	rows, err := repo.db.Query("select id,name,password,last_login from auth where name = $1", username)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}
	entry := &AuthEntry{}
	err = rows.Scan(&entry.ID, &entry.Name, &entry.Password, &entry.LastLogin)
	return entry, err
}
