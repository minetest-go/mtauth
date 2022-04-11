package db

import (
	"database/sql"
	"errors"

	_ "modernc.org/sqlite"
)

type SQliteAuthRepository struct {
	db *sql.DB
}

func NewSQliteAuthRepository(filename string) (*SQliteAuthRepository, error) {
	db, err := sql.Open("sqlite", "file:"+filename+"?mode=ro")
	if err != nil {
		return nil, err
	}

	return &SQliteAuthRepository{db: db}, nil
}

func (repo *SQliteAuthRepository) Migrate() error {
	result, err := repo.db.Query("pragma journal_mode;")
	if err != nil {
		return err
	}

	if !result.Next() {
		return errors.New("no results returned")
	}

	var mode string
	err = result.Scan(&mode)
	if err != nil {
		return err
	}

	if mode != "wal" {
		_, err = repo.db.Exec("pragma journal_mode = wal;")
		if err != nil {
			return errors.New("couldn't switch the db-journal to wal-mode, please stop the minetest-engine to allow doing this or do it manually: " + err.Error())
		}
	}

	return nil
}

func (repo *SQliteAuthRepository) Close() error {
	return repo.db.Close()
}
