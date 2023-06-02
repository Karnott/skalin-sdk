package skalinsdk

import (
	"os"
)

func GetSkalinClientApiID() string {
	return os.Getenv("TEST_SKALIN_CLIENT_API_ID")
}
func GetSkalinClientApiSecret() string {
	return os.Getenv("TEST_SKALIN_CLIENT_API_SECRET")
}
func GetSkalinAppClientID() string {
	return os.Getenv("TEST_SKALIN_APP_CLIENT_ID")
}
