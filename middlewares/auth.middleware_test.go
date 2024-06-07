package middlewares_test

import (
	"ngobar/berkicau/middlewares"
	"testing"
)

func TestEncryptSuccess(t *testing.T) {
	uid := uint(54321)
	secretKey := "supersecretkey"
	token, err := middlewares.Authenticate(uid, secretKey)
	if err != nil {
		t.Fatalf("error occurred while encrypting: %v", err.Error())
	}

	claims, err := middlewares.DecryptToken(token, secretKey)
	if err != nil {
		t.Fatalf("error occurred while decrypting: %v", err.Error())
	}

	if claims.UID != uid {
		t.Fatalf("unexpected payload, expect: %d, got: %d", uid, claims.UID)
	}
}