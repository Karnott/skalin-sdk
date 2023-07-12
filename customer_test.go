package skalinsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCustomCustomerMarshaller(t *testing.T) {
	customAttribute := CustomAttributes{
		"customAttribute1": "customValue1",
		"customAttribute2": "customValue2",
	}
	customer := &Customer{
		CustomAttributes: customAttribute,
		Name:             "Mon super nom",
	}
	customerBytes, err := json.Marshal(customer)
	if !assert.NoError(t, err) {
		return
	}
	var result map[string]interface{}
	err = json.Unmarshal(customerBytes, &result)
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, customer.Name, result["name"])
	assert.Equal(t, customAttribute["customAttribute1"], result["customAttribute1"])
	assert.Equal(t, customAttribute["customAttribute2"], result["customAttribute2"])
}

func TestSaveCustomer(t *testing.T) {
	var fakeCustomerData = `{
		"id": "",
		"refId": "2",
		"name": "Mon super nom", 
		"stage": "Mon super stage",
		"tags": ["tag1"]
  }`
	var fakeCustomerResponse = `
	{
		"status":"success",
		"data": ` + fakeCustomerData + `
	}`

	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		expectedBody := []byte(fakeCustomerResponse)
		var expectedCustomer Customer
		err := json.Unmarshal([]byte(fakeCustomerData), &expectedCustomer)
		if !assert.NoError(t, err) {
			return
		}

		_expectedCustomer, err := json.Marshal(expectedCustomer)
		if !assert.NoError(t, err) {
			return
		}
		mockApi.On(
			"send",
			http.MethodPost,
			BuildUrl(SAVE_CUSTOMER_PATH),
			jsonContentType,
			mock.Anything,
			_expectedCustomer,
			mock.Anything,
			http.StatusOK,
		).Return(nil, expectedBody, nil)

		skalinAPI := &skalinAPI{api: mockApi}
		customer, err := skalinAPI.SaveCustomer(expectedCustomer)
		mockApi.AssertExpectations(t)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, expectedCustomer.RefId, customer.RefId)
		assert.Equal(t, expectedCustomer.Name, customer.Name)
		assert.Equal(t, expectedCustomer.Stage, customer.Stage)
		assert.Equal(t, expectedCustomer.Tags[0], customer.Tags[0])
	})

	t.Run("With error", func(t *testing.T) {
		mockApi := new(MockAPI)
		var expectedCustomer Customer
		err := json.Unmarshal([]byte(fakeCustomerData), &expectedCustomer)
		if !assert.NoError(t, err) {
			return
		}

		_expectedCustomer, err := json.Marshal(expectedCustomer)
		if !assert.NoError(t, err) {
			return
		}
		mockApi.On(
			"send",
			http.MethodPost,
			BuildUrl(SAVE_CUSTOMER_PATH),
			jsonContentType,
			mock.Anything,
			_expectedCustomer,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, fmt.Errorf("Status code != %v: %v", http.StatusInternalServerError, http.StatusOK))

		skalinAPI := &skalinAPI{api: mockApi}
		customer, err := skalinAPI.SaveCustomer(expectedCustomer)
		if !assert.Error(t, err) {
			return
		}
		assert.Nil(t, customer)
	})

	t.Run("Call API", func(t *testing.T) {
		if GetSkalinAppClientID() == "" || GetSkalinClientApiID() == "" || GetSkalinClientApiSecret() == "" || GetSkalinCustomerCustomAttributeForTest() == "" {
			return
		}
		skalinApi, err := New(GetSkalinAppClientID(), GetSkalinClientApiID(), GetSkalinClientApiSecret())
		if !assert.NoError(t, err) {
			return
		}
		customer := Customer{
			RefId: "2",
			Name:  "Ceci est un test de l'API (nom du customer)",
			Stage: "ceci est un test de l'api (stage)",
			Tags:  []string{"tag1", "tag2", "tag3"},
			CustomAttributes: CustomAttributes{
				GetSkalinCustomerCustomAttributeForTest(): "Ceci est un test de l'API (attribut personnalisé 2)",
			},
		}
		customerSaved, err := skalinApi.SaveCustomer(customer)
		if !assert.NoError(t, err) {
			return
		}
		assert.NotEqual(t, "", customerSaved.Id)
		assert.Equal(t, customer.RefId, customerSaved.RefId)
		assert.Equal(t, customer.Name, customerSaved.Name)
		assert.Equal(t, customer.Stage, customerSaved.Stage)
		assert.Equal(t, len(customer.Tags), len(customerSaved.Tags))
	})
}

func TestGetCustomers(t *testing.T) {
	var fakeCustomersData = `[
  {
    "id": "1",
		"refId": "2",
		"name": "Mon super nom", 
		"stage": "Mon super stage"
  }
  ]`
	var fakeCustomersResponse = `
	{
		"status":"success",
		"data": ` + fakeCustomersData + `
	}`

	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		expectedBody := []byte(fakeCustomersResponse)
		_expectedCustomers := []byte(fakeCustomersData)
		var expectedCustomers []Customer
		err := json.Unmarshal(_expectedCustomers, &expectedCustomers)
		if !assert.NoError(t, err) {
			return
		}

		mockApi.On(
			"send",
			http.MethodGet,
			BuildUrl(SAVE_CUSTOMER_PATH),
			jsonContentType,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			http.StatusOK,
		).Return(nil, expectedBody, nil)

		skalinAPI := &skalinAPI{api: mockApi}
		customers, err := skalinAPI.GetCustomers(nil)
		mockApi.AssertExpectations(t)
		if !assert.NoError(t, err) || !assert.Equal(t, len(expectedCustomers), len(customers)) {
			return
		}
		expectedCustomer := expectedCustomers[0]
		customer := customers[0]
		assert.Equal(t, expectedCustomer.Id, customer.Id)
	})

	t.Run("With error", func(t *testing.T) {
		mockApi := new(MockAPI)
		mockApi.On(
			"send",
			http.MethodGet,
			BuildUrl(SAVE_CUSTOMER_PATH),
			jsonContentType,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, fmt.Errorf("Status code != %v: %v", http.StatusInternalServerError, http.StatusOK))

		skalinAPI := &skalinAPI{api: mockApi}
		customers, err := skalinAPI.GetCustomers(nil)
		mockApi.AssertExpectations(t)
		if !assert.Error(t, err) {
			return
		}
		assert.Nil(t, customers)
	})

	t.Run("Call API", func(t *testing.T) {
		if GetSkalinAppClientID() == "" || GetSkalinClientApiID() == "" || GetSkalinClientApiSecret() == "" {
			return
		}
		skalinApi, err := New(GetSkalinAppClientID(), GetSkalinClientApiID(), GetSkalinClientApiSecret())
		if !assert.NoError(t, err) {
			return
		}
		customers, err := skalinApi.GetCustomers(nil)
		if !assert.NoError(t, err) {
			return
		}
		if len(customers) == 0 {
			return
		}
		for _, customer := range customers {
			assert.NotEqual(t, "", customer.Id)
		}
	})
}
