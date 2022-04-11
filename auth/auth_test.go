package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test:enter from the database column
const CREDENTIALS = "#1#TxqLUa/uEJvZzPc3A0xwpA#oalXnktlS0bskc7bccsoVTeGwgAwUOyYhhceBu7wAyITkYjCtrzcDg6W5Co5V+oWUSG13y7TIoEfIg6rafaKzAbwRUC9RVGCeYRIUaa0hgEkIe9VkDmpeQ/kfF8zT8p7prOcpyrjWIJR+gmlD8Bf1mrxoPoBLDbvmxkcet327kQ9H4EMlIlv+w3XCufoPGFQ1UrfWiVqqK8dEmt/ldLPfxiK1Rg8MkwswEekymP1jyN9Cpq3w8spVVcjsxsAzI5M7QhSyqMMrIThdgBsUqMBOCULdV+jbRBBiA/ClywtZ8vvBpN9VGqsQuhmQG0h5x3fqPyR2XNdp9Ocm3zHBoJy/w"

func TestAuth(t *testing.T) {
	salt, verifier, err := ParseDBPassword(CREDENTIALS)
	assert.NoError(t, err)

	// valid credentials
	success, err := VerifyAuth("test", "enter", salt, verifier)
	assert.NoError(t, err)
	assert.True(t, success)

	// invalid password
	success, err = VerifyAuth("test", "bogus", salt, verifier)
	assert.NoError(t, err)
	assert.False(t, success)

	// invalid user
	success, err = VerifyAuth("testx", "enter", salt, verifier)
	assert.NoError(t, err)
	assert.False(t, success)

	// ad-hoc creation
	salt, verifier, err = CreateAuth("a", "b")

	// valid credentials
	success, err = VerifyAuth("a", "b", salt, verifier)
	assert.NoError(t, err)
	assert.True(t, success)
}
