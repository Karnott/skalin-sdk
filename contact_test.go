package skalinsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetContacts(t *testing.T) {
	var fakeContactsData = `[
  {
    "id": "1",
		"refId": "2",
		"firsName": "Mon super prenom", 
		"lastName": "Mon super nom de famille"
  }
  ]`
	var fakeContactsResponse = `
	{
		"status":"success",
		"data": ` + fakeContactsData + `
	}`

	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		expectedBody := []byte(fakeContactsResponse)
		_expectedContacts := []byte(fakeContactsData)
		var expectedContacts []Contact
		err := json.Unmarshal(_expectedContacts, &expectedContacts)
		if !assert.NoError(t, err) {
			return
		}

		mockApi.On(
			"send",
			http.MethodGet,
			BuildUrl(SAVE_CONTACT_PATH),
			jsonContentType,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			http.StatusOK,
		).Return(nil, expectedBody, nil)

		skalinAPI := &skalin{api: mockApi}
		contacts, err := skalinAPI.GetContacts(nil)
		mockApi.AssertExpectations(t)
		if !assert.NoError(t, err) || !assert.Equal(t, len(expectedContacts), len(contacts)) {
			return
		}
		expectedContact := expectedContacts[0]
		contact := contacts[0]
		assert.Equal(t, expectedContact.Id, contact.Id)
	})

	t.Run("With error", func(t *testing.T) {
		mockApi := new(MockAPI)
		mockApi.On(
			"send",
			http.MethodGet,
			BuildUrl(SAVE_CONTACT_PATH),
			jsonContentType,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, fmt.Errorf("Status code != %v: %v", http.StatusInternalServerError, http.StatusOK))

		skalinAPI := &skalin{api: mockApi}
		contacts, err := skalinAPI.GetContacts(nil)
		mockApi.AssertExpectations(t)
		if !assert.Error(t, err) {
			return
		}
		assert.Nil(t, contacts)
	})

	t.Run("Call API", func(t *testing.T) {
		if GetSkalinAppClientID() == "" || GetSkalinClientApiID() == "" || GetSkalinClientApiSecret() == "" {
			return
		}
		skalinApi, err := New(GetSkalinAppClientID(), GetSkalinClientApiID(), GetSkalinClientApiSecret())
		if !assert.NoError(t, err) {
			return
		}
		contacts, err := skalinApi.GetContacts(nil)
		if !assert.NoError(t, err) {
			return
		}
		if len(contacts) == 0 {
			return
		}
		for _, contact := range contacts {
			assert.NotEqual(t, "", contact.Id)
		}
	})
}
