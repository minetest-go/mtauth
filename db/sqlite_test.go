package db

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func TestCheckJournalModeDelete(t *testing.T) {
	dbfile, err := os.CreateTemp(os.TempDir(), "auth.sqlite")
	assert.NoError(t, err)
	assert.NotNil(t, dbfile)
	copyFileContents("testdata/auth.sqlite", dbfile.Name())

	db, err := NewSQliteAuthRepository(dbfile.Name())
	assert.NoError(t, err)
	assert.Error(t, db.Migrate())
	assert.NoError(t, db.Close())
}

func TestCheckJournalModeWal(t *testing.T) {
	dbfile, err := os.CreateTemp(os.TempDir(), "auth.wal.sqlite")
	assert.NoError(t, err)
	assert.NotNil(t, dbfile)
	copyFileContents("testdata/auth.wal.sqlite", dbfile.Name())

	db, err := NewSQliteAuthRepository(dbfile.Name())
	assert.NoError(t, err)
	assert.NoError(t, db.Migrate())
	assert.NoError(t, db.Close())
}
