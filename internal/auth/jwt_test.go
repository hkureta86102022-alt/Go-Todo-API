package auth

import (
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	t.Setenv("JWT_SECRET", "test")

	token, err := GenerateJWT(-1)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if token == "" {
		t.Fatal("expected token, got empty string")
	}
}
