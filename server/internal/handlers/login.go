package handlers

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

var (
	fakehash = func() []byte {
		hash, _ := bcrypt.GenerateFromPassword([]byte("TlI4jGmI9tCr9c6VzaCtfyow3")) // randomly generated
		return hash
	}
)

func loginUser(ctx context.Context)
