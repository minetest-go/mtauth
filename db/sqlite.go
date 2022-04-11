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
	db, err := sql.Open("sqlite", "file:"+filename)
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

func (repo *SQliteAuthRepository) GetByUsername(username string) (*AuthEntry, error) {
	rows, err := repo.db.Query("select id,name,password,last_login from auth where name = ?", username)
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

func (repo *SQliteAuthRepository) Create(entry *AuthEntry) error {
	result, err := repo.db.Exec("insert into auth(id,name,password,last_login) values(?,?,?,?)", entry.ID, entry.Name, entry.Password, entry.LastLogin)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	// assign returned id
	entry.ID = &id
	return nil
}

func (repo *SQliteAuthRepository) Update(entry *AuthEntry) error {
	_, err := repo.db.Exec("update auth set name = ?, password = ?, last_login = ? where id = ?", entry.Name, entry.Password, entry.LastLogin, entry.ID)
	return err
}

func (repo *SQliteAuthRepository) Delete(id int64) error {
	_, err := repo.db.Exec("delete from auth where id = ?", id)
	return err
}
