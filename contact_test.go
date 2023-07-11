package skalinsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCustomMarshaller(t *testing.T) {
	customAttribute := CustomAttributes{
		"customAttribute1": "customValue1",
		"customAttribute2": "customValue2",
	}
	contact := &Contact{
		CustomAttributes: customAttribute,
		FirstName:        "Mon super prenom",
	}
	contactBytes, err := json.Marshal(contact)
	if !assert.NoError(t, err) {
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal(contactBytes, &result)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, contact.FirstName, result["firstName"])
	assert.Equal(t, customAttribute["customAttribute1"], result["customAttribute1"])
	assert.Equal(t, customAttribute["customAttribute2"], result["customAttribute2"])
}

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

		skalinAPI := &skalinAPI{api: mockApi}
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

		skalinAPI := &skalinAPI{api: mockApi}
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

func TestSaveContact(t *testing.T) {
	var fakeContactData = `{
		"id": "",
		"refId": "2",
		"firsName": "Mon super prenom", 
		"lastName": "Mon super nom de famille",
		"email": "Mon super nom de famille",
		"phone": "Mon super nom de famille",
		"tags": ["tag1"]
  }`
	var fakeContactResponse = `
	{
		"status":"success",
		"data": ` + fakeContactData + `
	}`

	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		expectedBody := []byte(fakeContactResponse)
		var expectedContact Contact
		err := json.Unmarshal([]byte(fakeContactData), &expectedContact)
		if !assert.NoError(t, err) {
			return
		}

		_expectedContact, err := json.Marshal(expectedContact)
		if !assert.NoError(t, err) {
			return
		}
		mockApi.On(
			"send",
			http.MethodPost,
			BuildUrl(SAVE_CONTACT_PATH),
			jsonContentType,
			mock.Anything,
			_expectedContact,
			mock.Anything,
			http.StatusOK,
		).Return(nil, expectedBody, nil)

		skalinAPI := &skalinAPI{api: mockApi}
		contact, err := skalinAPI.SaveContact(expectedContact)
		mockApi.AssertExpectations(t)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, expectedContact.RefId, contact.RefId)
		assert.Equal(t, expectedContact.LastName, contact.LastName)
		assert.Equal(t, expectedContact.Email, contact.Email)
	})

	t.Run("With error", func(t *testing.T) {
		mockApi := new(MockAPI)
		var expectedContact Contact
		err := json.Unmarshal([]byte(fakeContactData), &expectedContact)
		if !assert.NoError(t, err) {
			return
		}

		_expectedContact, err := json.Marshal(expectedContact)
		if !assert.NoError(t, err) {
			return
		}
		mockApi.On(
			"send",
			http.MethodPost,
			BuildUrl(SAVE_CONTACT_PATH),
			jsonContentType,
			mock.Anything,
			_expectedContact,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, fmt.Errorf("Status code != %v: %v", http.StatusInternalServerError, http.StatusOK))

		skalinAPI := &skalinAPI{api: mockApi}
		contact, err := skalinAPI.SaveContact(expectedContact)
		if !assert.Error(t, err) {
			return
		}
		assert.Nil(t, contact)
	})

	t.Run("Call API", func(t *testing.T) {
		if GetSkalinAppClientID() == "" || GetSkalinClientApiID() == "" || GetSkalinClientApiSecret() == "" || GetSkalinExistingCustomerRefIdForTest() == "" || GetSkalinContactCustomAttributeForTest() == "" {
			return
		}
		skalinApi, err := New(GetSkalinAppClientID(), GetSkalinClientApiID(), GetSkalinClientApiSecret())
		if !assert.NoError(t, err) {
			return
		}
		customerId := GetSkalinExistingCustomerRefIdForTest()
		contact := Contact{
			RefId:     "2",
			Customer:  &customerId,
			LastName:  "Ceci est un test de l'API (nom de famille)",
			FirstName: "Ceci est un test de l'API (prénom)",
			Email:     "contact+testapi@karnott.fr",
			Phone:     "0123456789",
			Tags:      []string{"tag1", "tag2", "tag3"},
			CustomAttributes: CustomAttributes{
				GetSkalinContactCustomAttributeForTest(): "Ceci est un test de l'API (attribut personnalisé 2)",
			},
		}
		contactSaved, err := skalinApi.SaveContact(contact)
		if !assert.NoError(t, err) {
			return
		}
		assert.NotEqual(t, "", contactSaved.Id)
		assert.Equal(t, contact.RefId, contactSaved.RefId)
		assert.Equal(t, contact.Email, contactSaved.Email)
		assert.Equal(t, contact.LastName, contactSaved.LastName)
		assert.Equal(t, contact.Phone, contactSaved.Phone)
	})
}

func TestCreateContactForCustomer(t *testing.T) {
	var fakeContactData = `{
		"id": "",
		"customerId": "1",
		"refId": "2",
		"firsName": "Mon super prenom", 
		"lastName": "Mon super nom de famille",
		"email": "Mon super nom de famille",
		"phone": "Mon super nom de famille",
		"tags": ["tag1"]
  }`
	var fakeContactResponse = `
	{
		"status":"success",
		"data": ` + fakeContactData + `
	}`

	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		expectedBody := []byte(fakeContactResponse)
		var expectedContact Contact
		err := json.Unmarshal([]byte(fakeContactData), &expectedContact)
		if !assert.NoError(t, err) {
			return
		}

		customerId := "1"
		expectedContact.CustomerId = nil // no need to send customerId in body
		_expectedContact, err := json.Marshal(expectedContact)
		if !assert.NoError(t, err) {
			return
		}
		mockApi.On(
			"send",
			http.MethodPost,
			BuildUrl(fmt.Sprintf(CREATE_CUSTOMER_CONTACT_PATH, customerId)),
			jsonContentType,
			mock.Anything,
			_expectedContact,
			mock.Anything,
			http.StatusOK,
		).Return(nil, expectedBody, nil)

		skalinAPI := &skalinAPI{api: mockApi}
		contact, err := skalinAPI.CreateContactForCustomer(expectedContact, customerId)
		mockApi.AssertExpectations(t)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, expectedContact.RefId, contact.RefId)
		assert.Equal(t, customerId, *contact.CustomerId)
		assert.Equal(t, expectedContact.LastName, contact.LastName)
		assert.Equal(t, expectedContact.Email, contact.Email)
	})

	t.Run("With error", func(t *testing.T) {
		mockApi := new(MockAPI)
		var expectedContact Contact
		err := json.Unmarshal([]byte(fakeContactData), &expectedContact)
		if !assert.NoError(t, err) {
			return
		}

		_expectedContact, err := json.Marshal(expectedContact)
		if !assert.NoError(t, err) {
			return
		}
		customerId := "1"
		mockApi.On(
			"send",
			http.MethodPost,
			BuildUrl(fmt.Sprintf(CREATE_CUSTOMER_CONTACT_PATH, customerId)),
			jsonContentType,
			mock.Anything,
			_expectedContact,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, fmt.Errorf("Status code != %v: %v", http.StatusInternalServerError, http.StatusOK))

		skalinAPI := &skalinAPI{api: mockApi}
		contact, err := skalinAPI.CreateContactForCustomer(expectedContact, customerId)
		if !assert.Error(t, err) {
			return
		}
		assert.Nil(t, contact)
	})

	t.Run("Call API", func(t *testing.T) {
		if GetSkalinAppClientID() == "" || GetSkalinClientApiID() == "" || GetSkalinClientApiSecret() == "" || GetSkalinExistingCustomerIdForTest() == "" || GetSkalinContactCustomAttributeForTest() == "" {
			return
		}
		skalinApi, err := New(GetSkalinAppClientID(), GetSkalinClientApiID(), GetSkalinClientApiSecret())
		if !assert.NoError(t, err) {
			return
		}
		customerId := GetSkalinExistingCustomerIdForTest()
		contact := Contact{
			RefId:     "3",
			LastName:  "Ceci est un test de l'API (nom de famille)",
			FirstName: "Ceci est un test de l'API (prénom)",
			Email:     "contact+testapi3@karnott.fr",
			Phone:     "0123456789",
			Tags:      []string{"tag1", "tag2", "tag3"},
			CustomAttributes: CustomAttributes{
				GetSkalinContactCustomAttributeForTest(): "Ceci est un test de l'API (attribut personnalisé 3)",
			},
		}
		contactSaved, err := skalinApi.CreateContactForCustomer(contact, customerId)
		if !assert.NoError(t, err) {
			return
		}
		assert.NotEqual(t, "", contactSaved.Id)
		assert.Equal(t, contact.RefId, contactSaved.RefId)
		assert.Equal(t, contact.Email, contactSaved.Email)
		assert.Equal(t, contact.LastName, contactSaved.LastName)
		assert.Equal(t, contact.Phone, contactSaved.Phone)
	})
}

func TestUpdateContact(t *testing.T) {
	var fakeContactData = `{
		"id": "12345",
		"refId": "2",
		"firsName": "Mon super prenom", 
		"lastName": "Mon super nom de famille",
		"email": "Mon super nom de famille",
		"phone": "Mon super nom de famille",
		"tags": ["tag1"]
  }`
	var fakeContactResponse = `
	{
		"status":"success",
	}`

	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		expectedBody := []byte(fakeContactResponse)
		var expectedContact Contact
		err := json.Unmarshal([]byte(fakeContactData), &expectedContact)
		if !assert.NoError(t, err) {
			return
		}

		_expectedContact, err := json.Marshal(expectedContact)
		if !assert.NoError(t, err) {
			return
		}
		mockApi.On(
			"send",
			http.MethodPatch,
			BuildUrl(fmt.Sprintf(UPDATE_CONTACT_PATH, expectedContact.Id)),
			jsonContentType,
			mock.Anything,
			_expectedContact,
			mock.Anything,
			http.StatusOK,
		).Return(nil, expectedBody, nil)

		skalinAPI := &skalinAPI{api: mockApi}
		contact, err := skalinAPI.UpdateContact(expectedContact)
		mockApi.AssertExpectations(t)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, expectedContact.RefId, contact.RefId)
		assert.Equal(t, expectedContact.LastName, contact.LastName)
		assert.Equal(t, expectedContact.Email, contact.Email)
	})

	t.Run("With error", func(t *testing.T) {
		mockApi := new(MockAPI)
		var expectedContact Contact
		err := json.Unmarshal([]byte(fakeContactData), &expectedContact)
		if !assert.NoError(t, err) {
			return
		}

		_expectedContact, err := json.Marshal(expectedContact)
		if !assert.NoError(t, err) {
			return
		}
		mockApi.On(
			"send",
			http.MethodPatch,
			BuildUrl(fmt.Sprintf(UPDATE_CONTACT_PATH, expectedContact.Id)),
			jsonContentType,
			mock.Anything,
			_expectedContact,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, fmt.Errorf("Status code != %v: %v", http.StatusInternalServerError, http.StatusOK))

		skalinAPI := &skalinAPI{api: mockApi}
		contact, err := skalinAPI.UpdateContact(expectedContact)
		if !assert.Error(t, err) {
			return
		}
		assert.Nil(t, contact)
	})

	t.Run("Call API", func(t *testing.T) {
		if GetSkalinAppClientID() == "" || GetSkalinClientApiID() == "" || GetSkalinClientApiSecret() == "" || GetSkalinExistingContactIdForTest() == "" && GetSkalinExistingCustomerRefIdForTest() == "" || GetSkalinContactCustomAttributeForTest() == "" {
			return
		}
		skalinApi, err := New(GetSkalinAppClientID(), GetSkalinClientApiID(), GetSkalinClientApiSecret())
		if !assert.NoError(t, err) {
			return
		}
		customerId := GetSkalinExistingCustomerRefIdForTest()
		contactId := GetSkalinExistingContactIdForTest()
		contact := Contact{
			Id:        contactId,
			Customer:  &customerId,
			LastName:  "Ceci est un test de l'API (nom de famille) 4",
			FirstName: "Ceci est un test de l'API (prénom) 5",
			Email:     "contact+testapi@karnott.fr",
			Phone:     "0123456789",
			Tags:      []string{"tag1", "tag2", "tag3"},
			CustomAttributes: CustomAttributes{
				GetSkalinContactCustomAttributeForTest(): "Ceci est un test de l'API (attribut personnalisé 2)",
			},
		}
		contactSaved, err := skalinApi.UpdateContact(contact)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, contact.Email, contactSaved.Email)
		assert.Equal(t, contact.FirstName, contactSaved.FirstName)
		assert.Equal(t, contact.LastName, contactSaved.LastName)
		assert.Equal(t, contact.Phone, contactSaved.Phone)
	})
}
