package templates

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplates(t *testing.T) {
	tmpl := Parse("index.html")
	assert.NotNil(t, tmpl)

	buf := bytes.Buffer{}
	w := bufio.NewWriter(&buf)
	err := tmpl.Execute(w, nil)
	assert.NoError(t, err)
	assert.NoError(t, w.Flush())
	assert.True(t, len(buf.Bytes()) > 0)
}
