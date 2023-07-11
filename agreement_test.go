package skalinsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAgreements(t *testing.T) {
	var fakeAgreementsData = `[
  {
    "id": "1",
		"refId": "2",
		"startDate": "2020-01-01", 
		"engagement": 1,
		"engagementPeriod": "YEAR",
		"notice": 1,
		"noticePeriod": "MONTH",
		"plan": "PLAN",
		"type": "INITIAL",
		"mrr": 100,
		"fee": 10,
		"autoRenew": true
  }
  ]`
	var fakeAgreementsResponse = `
	{
		"status":"success",
		"data": ` + fakeAgreementsData + `
	}`

	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		expectedBody := []byte(fakeAgreementsResponse)
		_expectedAgreements := []byte(fakeAgreementsData)
		var expectedAgreements []Agreement
		err := json.Unmarshal(_expectedAgreements, &expectedAgreements)
		if !assert.NoError(t, err) {
			return
		}

		mockApi.On(
			"send",
			http.MethodGet,
			BuildUrl(SAVE_AGREEMENT_PATH),
			jsonContentType,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			http.StatusOK,
		).Return(nil, expectedBody, nil)

		skalinAPI := &skalinAPI{api: mockApi}
		agreements, err := skalinAPI.GetAgreements(nil)
		mockApi.AssertExpectations(t)
		if !assert.NoError(t, err) || !assert.Equal(t, len(expectedAgreements), len(agreements)) {
			return
		}
		expectedAgreement := expectedAgreements[0]
		agreement := agreements[0]
		assert.Equal(t, expectedAgreement.Id, agreement.Id)
	})

	t.Run("With error", func(t *testing.T) {
		mockApi := new(MockAPI)
		mockApi.On(
			"send",
			http.MethodGet,
			BuildUrl(SAVE_AGREEMENT_PATH),
			jsonContentType,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, fmt.Errorf("Status code != %v: %v", http.StatusInternalServerError, http.StatusOK))

		skalinAPI := &skalinAPI{api: mockApi}
		agreements, err := skalinAPI.GetAgreements(nil)
		mockApi.AssertExpectations(t)
		if !assert.Error(t, err) {
			return
		}
		assert.Nil(t, agreements)
	})

	t.Run("Call API", func(t *testing.T) {
		if GetSkalinAppClientID() == "" || GetSkalinClientApiID() == "" || GetSkalinClientApiSecret() == "" {
			return
		}
		skalinApi, err := New(GetSkalinAppClientID(), GetSkalinClientApiID(), GetSkalinClientApiSecret())
		if !assert.NoError(t, err) {
			return
		}
		agreements, err := skalinApi.GetAgreements(nil)
		if !assert.NoError(t, err) {
			return
		}
		if len(agreements) == 0 {
			return
		}
		for _, agreement := range agreements {
			assert.NotEqual(t, "", agreement.Id)
		}
	})
}

func TestSaveAgreement(t *testing.T) {
	var fakeAgreementData = `{
		"id": "",
		"refId": "2",
		"startDate": "2020-01-01", 
		"engagement": 1,
		"engagementPeriod": "YEAR",
		"notice": 1,
		"noticePeriod": "MONTH",
		"plan": "PLAN",
		"type": "INITIAL",
		"mr": 100,
		"fee": 10,
		"autoRenew": true
  }`
	var fakeAgreementResponse = `
	{
		"status":"success",
		"data": ` + fakeAgreementData + `
	}`

	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		expectedBody := []byte(fakeAgreementResponse)
		var expectedAgreement Agreement
		err := json.Unmarshal([]byte(fakeAgreementData), &expectedAgreement)
		if !assert.NoError(t, err) {
			return
		}

		_expectedAgreement, err := json.Marshal(expectedAgreement)
		if !assert.NoError(t, err) {
			return
		}
		mockApi.On(
			"send",
			http.MethodPost,
			BuildUrl(SAVE_AGREEMENT_PATH),
			jsonContentType,
			mock.Anything,
			_expectedAgreement,
			mock.Anything,
			http.StatusOK,
		).Return(nil, expectedBody, nil)

		skalinAPI := &skalinAPI{api: mockApi}
		agreement, err := skalinAPI.SaveAgreement(expectedAgreement)
		mockApi.AssertExpectations(t)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, expectedAgreement.RefId, agreement.RefId)
		assert.Equal(t, expectedAgreement.Notice, agreement.Notice)
		assert.Equal(t, expectedAgreement.NoticePeriod, agreement.NoticePeriod)
		assert.Equal(t, expectedAgreement.StartDate, agreement.StartDate)
		assert.Equal(t, expectedAgreement.Engagement, agreement.Engagement)
		assert.Equal(t, expectedAgreement.EngagementPeriod, agreement.EngagementPeriod)
		assert.Equal(t, expectedAgreement.Plan, agreement.Plan)
		assert.Equal(t, expectedAgreement.Mrr, agreement.Mrr)
		assert.Equal(t, expectedAgreement.Fee, agreement.Fee)
	})

	t.Run("With error", func(t *testing.T) {
		mockApi := new(MockAPI)
		var expectedAgreement Agreement
		err := json.Unmarshal([]byte(fakeAgreementData), &expectedAgreement)
		if !assert.NoError(t, err) {
			return
		}

		_expectedAgreement, err := json.Marshal(expectedAgreement)
		if !assert.NoError(t, err) {
			return
		}
		mockApi.On(
			"send",
			http.MethodPost,
			BuildUrl(SAVE_AGREEMENT_PATH),
			jsonContentType,
			mock.Anything,
			_expectedAgreement,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, fmt.Errorf("Status code != %v: %v", http.StatusInternalServerError, http.StatusOK))

		skalinAPI := &skalinAPI{api: mockApi}
		agreement, err := skalinAPI.SaveAgreement(expectedAgreement)
		if !assert.Error(t, err) {
			return
		}
		assert.Nil(t, agreement)
	})

}

func TestCreateAgreementForCustomer(t *testing.T) {
	var fakeAgreementData = `{
		"id": "",
		"customerId": "1",
		"refId": "2",
		"startDate": "2020-01-01", 
		"engagement": 1,
		"engagementPeriod": "YEAR",
		"notice": 1,
		"noticePeriod": "MONTH",
		"plan": "PLAN",
		"type": "INITIAL",
		"mr": 100,
		"fee": 10,
		"autoRenew": true
  }`
	var fakeAgreementResponse = `
	{
		"status":"success",
		"data": ` + fakeAgreementData + `
	}`

	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		expectedBody := []byte(fakeAgreementResponse)
		var expectedAgreement Agreement
		err := json.Unmarshal([]byte(fakeAgreementData), &expectedAgreement)
		if !assert.NoError(t, err) {
			return
		}

		customerId := "1"
		expectedAgreement.CustomerId = nil // no need to send customerId in body
		_expectedAgreement, err := json.Marshal(expectedAgreement)
		if !assert.NoError(t, err) {
			return
		}
		mockApi.On(
			"send",
			http.MethodPost,
			BuildUrl(fmt.Sprintf(CREATE_CUSTOMER_AGREEMENT_PATH, customerId)),
			jsonContentType,
			mock.Anything,
			_expectedAgreement,
			mock.Anything,
			http.StatusOK,
		).Return(nil, expectedBody, nil)

		skalinAPI := &skalinAPI{api: mockApi}
		agreement, err := skalinAPI.CreateAgreementForCustomer(expectedAgreement, customerId)
		mockApi.AssertExpectations(t)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, customerId, *agreement.CustomerId)
		assert.Equal(t, expectedAgreement.RefId, agreement.RefId)
		assert.Equal(t, expectedAgreement.Notice, agreement.Notice)
		assert.Equal(t, expectedAgreement.NoticePeriod, agreement.NoticePeriod)
		assert.Equal(t, expectedAgreement.StartDate, agreement.StartDate)
		assert.Equal(t, expectedAgreement.Engagement, agreement.Engagement)
		assert.Equal(t, expectedAgreement.EngagementPeriod, agreement.EngagementPeriod)
		assert.Equal(t, expectedAgreement.Plan, agreement.Plan)
		assert.Equal(t, expectedAgreement.Mrr, agreement.Mrr)
		assert.Equal(t, expectedAgreement.Fee, agreement.Fee)
	})

	t.Run("With error", func(t *testing.T) {
		mockApi := new(MockAPI)
		var expectedAgreement Agreement
		err := json.Unmarshal([]byte(fakeAgreementData), &expectedAgreement)
		if !assert.NoError(t, err) {
			return
		}

		_expectedAgreement, err := json.Marshal(expectedAgreement)
		if !assert.NoError(t, err) {
			return
		}
		customerId := "1"
		mockApi.On(
			"send",
			http.MethodPost,
			BuildUrl(fmt.Sprintf(CREATE_CUSTOMER_AGREEMENT_PATH, customerId)),
			jsonContentType,
			mock.Anything,
			_expectedAgreement,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, fmt.Errorf("Status code != %v: %v", http.StatusInternalServerError, http.StatusOK))

		skalinAPI := &skalinAPI{api: mockApi}
		agreement, err := skalinAPI.CreateAgreementForCustomer(expectedAgreement, customerId)
		if !assert.Error(t, err) {
			return
		}
		assert.Nil(t, agreement)
	})
}

func TestUpdateAgreement(t *testing.T) {
	var fakeAgreementData = `{
		"id": "12345",
		"refId": "2",
		"startDate": "2020-01-01", 
		"engagement": 1,
		"engagementPeriod": "YEAR",
		"notice": 1,
		"noticePeriod": "MONTH",
		"plan": "PLAN",
		"type": "INITIAL",
		"mr": 100,
		"fee": 10,
		"autoRenew": true
  }`
	var fakeAgreementResponse = `
	{
		"status":"success",
	}`

	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		expectedBody := []byte(fakeAgreementResponse)
		var expectedAgreement Agreement
		err := json.Unmarshal([]byte(fakeAgreementData), &expectedAgreement)
		if !assert.NoError(t, err) {
			return
		}

		_expectedAgreement, err := json.Marshal(expectedAgreement)
		if !assert.NoError(t, err) {
			return
		}
		mockApi.On(
			"send",
			http.MethodPatch,
			BuildUrl(fmt.Sprintf(UPDATE_AGREEMENT_PATH, expectedAgreement.Id)),
			jsonContentType,
			mock.Anything,
			_expectedAgreement,
			mock.Anything,
			http.StatusOK,
		).Return(nil, expectedBody, nil)

		skalinAPI := &skalinAPI{api: mockApi}
		agreement, err := skalinAPI.UpdateAgreement(expectedAgreement)
		mockApi.AssertExpectations(t)
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, expectedAgreement.RefId, agreement.RefId)
		assert.Equal(t, expectedAgreement.Engagement, agreement.Engagement)
		assert.Equal(t, expectedAgreement.EngagementPeriod, agreement.EngagementPeriod)
	})

	t.Run("With error", func(t *testing.T) {
		mockApi := new(MockAPI)
		var expectedAgreement Agreement
		err := json.Unmarshal([]byte(fakeAgreementData), &expectedAgreement)
		if !assert.NoError(t, err) {
			return
		}

		_expectedAgreement, err := json.Marshal(expectedAgreement)
		if !assert.NoError(t, err) {
			return
		}
		mockApi.On(
			"send",
			http.MethodPatch,
			BuildUrl(fmt.Sprintf(UPDATE_AGREEMENT_PATH, expectedAgreement.Id)),
			jsonContentType,
			mock.Anything,
			_expectedAgreement,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, fmt.Errorf("Status code != %v: %v", http.StatusInternalServerError, http.StatusOK))

		skalinAPI := &skalinAPI{api: mockApi}
		agreement, err := skalinAPI.UpdateAgreement(expectedAgreement)
		if !assert.Error(t, err) {
			return
		}
		assert.Nil(t, agreement)
	})
}
