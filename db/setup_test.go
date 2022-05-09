package db

import (
	"mtauth/worldconfig"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	cfg := make(map[string]string)
	cfg[worldconfig.CONFIG_AUTH_BACKEND] = worldconfig.BACKEND_SQLITE3

	repos, err := Setup(cfg)
	assert.NoError(t, err)
	assert.NotNil(t, repos)
}
