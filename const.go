package skalinsdk

import (
	"errors"
	"os"
)

const (
	SKALIN_API_URL  = "https://api.skalin.io/v1"
	SKALIN_AUTH_URL = "https://auth.skalin.io/oauth/token"
	jsonContentType = "application/json"
)

var (
	ErrUndefined     = errors.New("undefined error")
	ErrAuthorization = errors.New("No authorization token was found")
)

func GetAPIUrl() string {
	if v := os.Getenv("SKALIN_API_URL"); v != "" {
		return v
	}
	return SKALIN_API_URL
}

func BuildUrl(path string) string {
	return GetAPIUrl() + path
}
