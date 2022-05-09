package db

import "database/sql"

type SettingsRepository struct {
	db *sql.DB
}

type Setting struct {
	Key   string
	Value string
}

func NewSettingsRepository(db *sql.DB) *SettingsRepository {
	return &SettingsRepository{db: db}
}

func (repo *SettingsRepository) GetByKey(key string) (*Setting, error) {
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
