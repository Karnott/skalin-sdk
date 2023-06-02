package skalinsdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

type API interface {
	PutData(url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error)
	PostData(url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error)
	GetData(url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error)
	send(method, url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error)
	WithToken(token string) API
	GetLogger() *CustomLog
}

type SkalinAPI struct {
	clientID *string
	token    *string
	logger   *CustomLog
}

func (a *SkalinAPI) SetLogger(logger logrus.FieldLogger) {
	a.logger = &CustomLog{
		logger,
	}
}

func (a SkalinAPI) GetLogger() *CustomLog {
	if a.logger != nil {
		return a.logger
	}
	return Log
}

func (a SkalinAPI) WithToken(token string) API {
	a.token = &token
	return a
}
func (a SkalinAPI) WithClientID(clientID string) API {
	a.clientID = &clientID
	return a
}

func (a SkalinAPI) doRequest(method, queryUrl, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewBuffer(body)
	}
	req, err := http.NewRequest(method, queryUrl, bodyReader)
	if err != nil {
		return nil, err
	}
	for key, values := range extraHeaders {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	if a.token != nil {
		req.Header["Authorization"] = []string{"Bearer " + *a.token}
	}
	req.Header["Accept-Language"] = []string{"fr"}

	if queryParams == nil {
		queryParams = &url.Values{}
	}
	if a.clientID != nil {
		queryParams.Set("clientId", *a.clientID)
	}
	req.URL.RawQuery = queryParams.Encode()
	return http.DefaultClient.Do(req)
}

func (a SkalinAPI) send(method, url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error) {
	logFields := logrus.Fields{
		"method":             method,
		"url":                url,
		"content-type":       contentType,
		"expectedStatusCode": expectedStatusCode,
	}
	if queryParams != nil {
		logFields["params"] = queryParams.Encode()
	}
	a.GetLogger().WithFields(logFields).Infof("call skalin API")
	res, err := a.doRequest(
		method,
		url,
		contentType,
		extraHeaders,
		body,
		queryParams,
	)
	if err != nil {
		a.GetLogger().WithFields(logFields).Errorf("error to call skalin API: %v", err)
		return nil, nil, err
	}
	bodyResp, err := a.readResp(res, expectedStatusCode)
	if err != nil {
		if body != nil {
			logFields["body"] = string(body)
		}
		logFields["bodyResponse"] = string(body)
		logFields["responseStatusCode"] = res.StatusCode
		a.GetLogger().WithFields(logFields).Errorf("error to read response from skalin API: %v", err)
	}
	return res, bodyResp, err
}

func (a SkalinAPI) PostData(url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error) {
	return a.send(
		http.MethodPost,
		url,
		contentType,
		extraHeaders,
		body,
		queryParams,
		expectedStatusCode,
	)
}

func (a SkalinAPI) PutData(url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error) {
	return a.send(
		http.MethodPut,
		url,
		contentType,
		extraHeaders,
		body,
		queryParams,
		expectedStatusCode,
	)
}

func (a SkalinAPI) GetData(url, contentType string, extraHeaders map[string][]string, body []byte, queryParams *url.Values, expectedStatusCode int) (*http.Response, []byte, error) {
	return a.send(
		http.MethodGet,
		url,
		contentType,
		extraHeaders,
		body,
		queryParams,
		expectedStatusCode,
	)
}

func (a SkalinAPI) readResp(res *http.Response, expectedStatusCode int) ([]byte, error) {
	if res == nil {
		return nil, nil
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != expectedStatusCode {
		errorMessage := a.extractErrorMessage(body)
		if errorMessage == nil {
			errorMessage = fmt.Errorf("status code != %v: %v", expectedStatusCode, res.StatusCode)
		}
		return body, errorMessage
	}
	return body, err
}

type SkalinResponseError struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func (a SkalinAPI) extractErrorMessage(body []byte) error {
	if len(body) == 0 {
		return ErrUndefined
	}

	var responseErr SkalinResponseError
	err := json.Unmarshal(body, &responseErr)
	if err != nil {
		a.logger.Infof("error to unmarshal skalin response error: %v", err)
		return errors.New(string(body))
	}
	if responseErr.Message != "" {
		return errors.New(responseErr.Message)
	}
	return fmt.Errorf("error to call skalin API with code: %v", responseErr.Code)
}
