package templates

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplates(t *testing.T) {
	err := Init()
	assert.NoError(t, err)
	tmpl := Get("login.html")
	assert.NotNil(t, tmpl)

	buf := bytes.Buffer{}
	w := bufio.NewWriter(&buf)
	err = tmpl.Execute(w, true)
	assert.NoError(t, err)
	assert.NoError(t, w.Flush())
	assert.True(t, len(buf.Bytes()) > 0)
}
