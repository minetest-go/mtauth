package db

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
)

type setting_key string

const (
	SETTING_SHARED_SECRET setting_key = "shared_secret"
)

type SettingsRepository struct {
	db *sql.DB
}

type Setting struct {
	Key   setting_key
	Value string
}

func NewSettingsRepository(db *sql.DB) *SettingsRepository {
	return &SettingsRepository{db: db}
}

func (repo *SettingsRepository) GetByKey(key setting_key) (*Setting, error) {
	rows, err := repo.db.Query("select key,value from settings where key = $1", key)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}
	entry := &Setting{}
	err = rows.Scan(&entry.Key, &entry.Value)
	return entry, err
}

func (repo *SettingsRepository) Create(entry *Setting) error {
	_, err := repo.db.Exec("insert into settings(key,value) values($1,$2)", entry.Key, entry.Value)
	return err
}

func (repo *SettingsRepository) SetupDefaults() error {
	shared_secret_entry, err := repo.GetByKey(SETTING_SHARED_SECRET)
	if err != nil {
		return err
	}
	if shared_secret_entry == nil {
		key := make([]byte, 32)
		_, err = rand.Read(key)
		if err != nil {
			return err
		}
		shared_secret_entry = &Setting{
			Key:   SETTING_SHARED_SECRET,
			Value: base64.StdEncoding.EncodeToString(key),
		}
		err = repo.Create(shared_secret_entry)
		if err != nil {
			return err
		}
	}
	return nil
}
