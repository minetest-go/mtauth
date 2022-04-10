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
	journal_mode, err := CheckJournalMode(filename)
	if err != nil {
		return nil, err
	}

	if !journal_mode {
		return nil, errors.New("journal-mode is not wal")
	}

	db, err := sql.Open("sqlite", "file:"+filename+"?mode=ro")
	if err != nil {
		return nil, err
	}

	sq := &SQliteAuthRepository{db: db}
	return sq, nil
}

func CheckJournalMode(filename string) (bool, error) {
	db, err := sql.Open("sqlite", "file:"+filename+"?mode=ro")
	if err != nil {
		return false, err
	}

	result, err := db.Query("pragma journal_mode;")
	if err != nil {
		return false, err
	}

	if !result.Next() {
		return false, errors.New("no results returned")
	}

	var mode string
	err = result.Scan(&mode)
	return mode == "wal", err
}

func ChangeJournalMode(filename string) error {
	return nil
}
