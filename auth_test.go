package main

import (
	"encoding/base64"
	"fmt"
	"strings"
	"testing"
)

// test:enter from the database column
const CREDENTIALS = "#1#TxqLUa/uEJvZzPc3A0xwpA#oalXnktlS0bskc7bccsoVTeGwgAwUOyYhhceBu7wAyITkYjCtrzcDg6W5Co5V+oWUSG13y7TIoEfIg6rafaKzAbwRUC9RVGCeYRIUaa0hgEkIe9VkDmpeQ/kfF8zT8p7prOcpyrjWIJR+gmlD8Bf1mrxoPoBLDbvmxkcet327kQ9H4EMlIlv+w3XCufoPGFQ1UrfWiVqqK8dEmt/ldLPfxiK1Rg8MkwswEekymP1jyN9Cpq3w8spVVcjsxsAzI5M7QhSyqMMrIThdgBsUqMBOCULdV+jbRBBiA/ClywtZ8vvBpN9VGqsQuhmQG0h5x3fqPyR2XNdp9Ocm3zHBoJy/w"

func TestAuth(t *testing.T) {
	parts := strings.Split(CREDENTIALS, "#")
	if len(parts) != 4 {
		t.FailNow()
	}
	if parts[1] != "1" {
		t.FailNow()
	}
	fmt.Printf("%s\n", parts[2])
	//version := parts[1]
	salt, err := base64.RawStdEncoding.DecodeString(parts[2])
	if err != nil {
		t.Fatal(err)
	}
	verifier, err := base64.RawStdEncoding.DecodeString(parts[3])
	if err != nil {
		t.Fatal(err)
	}

	// valid credentials
	success, err := VerifyAuth("test", "enter", salt, verifier)
	if err != nil {
		t.Fatal(err)
	}
	if !success {
		t.FailNow()
	}

	// invalid password
	success, err = VerifyAuth("test", "bogus", salt, verifier)
	if err != nil {
		t.Fatal(err)
	}
	if success {
		t.FailNow()
	}

	// invalid user
	success, err = VerifyAuth("testx", "enter", salt, verifier)
	if err != nil {
		t.Fatal(err)
	}
	if success {
		t.FailNow()
	}
}
