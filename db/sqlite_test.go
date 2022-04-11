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

	db, err := NewSQliteAuthRepository(dbfile.Name() + "?mode=ro")
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

func TestSQliteRepo(t *testing.T) {
	// init stuff
	dbfile, err := os.CreateTemp(os.TempDir(), "auth.sqlite")
	assert.NoError(t, err)
	assert.NotNil(t, dbfile)
	copyFileContents("testdata/auth.wal.sqlite", dbfile.Name())

	// open db
	db, err := NewSQliteAuthRepository(dbfile.Name())
	assert.NoError(t, err)
	assert.NoError(t, db.Migrate())

	// existing entry
	entry, err := db.GetByUsername("test")
	assert.NoError(t, err)
	assert.NotNil(t, entry)
	assert.Equal(t, "test", entry.Name)
	assert.Equal(t, "#1#TxqLUa/uEJvZzPc3A0xwpA#oalXnktlS0bskc7bccsoVTeGwgAwUOyYhhceBu7wAyITkYjCtrzcDg6W5Co5V+oWUSG13y7TIoEfIg6rafaKzAbwRUC9RVGCeYRIUaa0hgEkIe9VkDmpeQ/kfF8zT8p7prOcpyrjWIJR+gmlD8Bf1mrxoPoBLDbvmxkcet327kQ9H4EMlIlv+w3XCufoPGFQ1UrfWiVqqK8dEmt/ldLPfxiK1Rg8MkwswEekymP1jyN9Cpq3w8spVVcjsxsAzI5M7QhSyqMMrIThdgBsUqMBOCULdV+jbRBBiA/ClywtZ8vvBpN9VGqsQuhmQG0h5x3fqPyR2XNdp9Ocm3zHBoJy/w", entry.Password)
	assert.Equal(t, int64(2), *entry.ID)
	assert.Equal(t, 1649603232, entry.LastLogin)

	// non-existing entry
	entry, err = db.GetByUsername("bogus")
	assert.NoError(t, err)
	assert.Nil(t, entry)

	// create entry
	new_entry := &AuthEntry{
		Name:      "createduser",
		Password:  "blah",
		LastLogin: 456,
	}
	assert.NoError(t, db.Create(new_entry))
	assert.NotNil(t, new_entry.ID)

	// check newly created entry
	entry, err = db.GetByUsername("createduser")
	assert.NoError(t, err)
	assert.NotNil(t, entry)
	assert.Equal(t, new_entry.Name, entry.Name)
	assert.Equal(t, new_entry.Password, entry.Password)
	assert.Equal(t, *new_entry.ID, *entry.ID)
	assert.Equal(t, new_entry.LastLogin, entry.LastLogin)

	// change things
	new_entry.Name = "x"
	new_entry.Password = "y"
	new_entry.LastLogin = 123
	assert.NoError(t, db.Update(new_entry))
	entry, err = db.GetByUsername("x")
	assert.NoError(t, err)
	assert.NotNil(t, entry)
	assert.Equal(t, new_entry.Name, entry.Name)
	assert.Equal(t, new_entry.Password, entry.Password)
	assert.Equal(t, *new_entry.ID, *entry.ID)
	assert.Equal(t, new_entry.LastLogin, entry.LastLogin)

	// remove new user
	assert.NoError(t, db.Delete(*new_entry.ID))
	entry, err = db.GetByUsername("x")
	assert.NoError(t, err)
	assert.Nil(t, entry)

	// cleanup
	assert.NoError(t, db.Close())
}
