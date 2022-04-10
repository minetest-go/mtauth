package main

import (
	"crypto/subtle"
	"mtauth/srp"
)

func VerifyAuth(username, password string, salt, verifier []byte) (bool, error) {
	// client
	pubA, privA, err := srp.InitiateHandshake()
	if err != nil {
		return false, err
	}

	// server
	B, _, K, err := srp.Handshake(pubA, verifier)

	// client
	clientK, err := srp.CompleteHandshake(pubA, privA, []byte(username), []byte(password), salt, B)
	if err != nil {
		return false, err
	}

	// server
	if subtle.ConstantTimeCompare(clientK, K) != 1 {
		return false, nil
	}

	return true, nil
}
