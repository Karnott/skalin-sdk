package skalinsdk

import (
	"net/http"
	"net/url"

	"github.com/stretchr/testify/mock"
)

type MockAPI struct {
	mock.Mock
}

func (m *MockAPI) PostData(url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error) {
	return m.send(http.MethodPost, url, contentType, extraHeaders, body, queryParams, expectedStatusCode)
}
func (m *MockAPI) PutData(url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error) {
	return m.send(http.MethodPut, url, contentType, extraHeaders, body, queryParams, expectedStatusCode)
}
func (m *MockAPI) PatchData(url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error) {
	return m.send(http.MethodPatch, url, contentType, extraHeaders, body, queryParams, expectedStatusCode)
}
func (m *MockAPI) GetData(url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error) {
	return m.send(http.MethodGet, url, contentType, extraHeaders, body, queryParams, expectedStatusCode)
}
func (m *MockAPI) GetLogger() *CustomLog {
	return Log
}

func (m *MockAPI) send(method, url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error) {
	args := m.Called(method, url, contentType, extraHeaders, body, queryParams, expectedStatusCode)
	arg0 := args.Get(0)
	var resp *http.Response
	if arg0 != nil {
		resp = arg0.(*http.Response)
	}
	arg1 := args.Get(1)
	var bodyResp []byte = nil
	if arg1 != nil {
		bodyResp = arg1.([]byte)
	}
	return resp, bodyResp, args.Error(2)
}

func (m *MockAPI) WithToken(_ string) API {
	return m
}
