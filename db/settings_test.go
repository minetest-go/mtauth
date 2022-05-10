package db

import (
	"database/sql"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSettings(t *testing.T) {
	// init stuff
	dbfile, err := os.CreateTemp(os.TempDir(), "settings.sqlite")
	assert.NoError(t, err)

	// open db
	db, err := sql.Open("sqlite", "file:"+dbfile.Name())
	assert.NoError(t, err)

	_, err = MigrateMTAuth(db)
	assert.NoError(t, err)

	repo := NewSettingsRepository(db)
	assert.NotNil(t, repo)

	s, err := repo.GetByKey("notfound")
	assert.NoError(t, err)
	assert.Nil(t, s)

	// create
	s = &Setting{
		Key:   SETTING_SHARED_SECRET,
		Value: "123",
	}
	assert.NoError(t, repo.Create(s))

	// fetch newly creates
	s, err = repo.GetByKey(SETTING_SHARED_SECRET)
	assert.NoError(t, err)
	assert.NotNil(t, s)
	assert.Equal(t, SETTING_SHARED_SECRET, s.Key)
	assert.Equal(t, "123", s.Value)
}
