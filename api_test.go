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

func GetSkalinContactCustomAttributeForTest() string {
	return os.Getenv("TEST_SKALIN_CONTACT_CUSTOM_ATTRIBUTE_ID")
}
func GetSkalinExistingCustomerRefIdForTest() string {
	return os.Getenv("TEST_SKALIN_EXISTING_CUSTOMER_REF_ID")
}

func GetSkalinExistingCustomerIdForTest() string {
	return os.Getenv("TEST_SKALIN_EXISTING_CUSTOMER_ID")
}

func GetSkalinExistingContactIdForTest() string {
	return os.Getenv("TEST_SKALIN_EXISTING_CONTACT_ID")
}
