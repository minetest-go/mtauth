package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckJournalModeDelete(t *testing.T) {
	ok, err := CheckJournalMode("testdata/auth.sqlite")
	assert.NoError(t, err)
	assert.False(t, ok)
}

func TestCheckJournalModeWal(t *testing.T) {
	ok, err := CheckJournalMode("testdata/auth.wal.sqlite")
	assert.NoError(t, err)
	assert.True(t, ok)
}
