package services

import (
	"github.com/vertinofff/blog-api/config"
	"testing"
)

func TestAccessTokenValidation(t *testing.T) {
	s := NewTokenService(&config.Config{JWT: config.JWTConfig{Secret: "access-secret-which-is-long-enough", RefreshSecret: "refresh-secret-which-is-long-enough", AccessTokenExpireDuration: 1, RefreshTokenExpireDuration: 2}})
	if _, err := s.GetAccessClaims("not-a-token"); err == nil {
		t.Fatal("malformed token accepted")
	}
	td, err := s.GenerateToken(7, "alice", "a@example.test")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = s.GetAccessClaims(td.AccessToken); err != nil {
		t.Fatalf("valid access token rejected: %v", err)
	}
}
