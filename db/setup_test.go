package db

import (
	"mtauth/worldconfig"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	cfg := make(map[string]string)
	cfg[worldconfig.CONFIG_AUTH_BACKEND] = worldconfig.BACKEND_SQLITE3

	repos, err := Setup(os.TempDir(), cfg)
	assert.NoError(t, err)
	assert.NotNil(t, repos)
}

func TestSetupUnsupported(t *testing.T) {
	cfg := make(map[string]string)
	cfg[worldconfig.CONFIG_AUTH_BACKEND] = worldconfig.BACKEND_FILES

	repos, err := Setup(os.TempDir(), cfg)
	assert.Error(t, err)
	assert.Nil(t, repos)
}
