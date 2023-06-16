package skalinsdk

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Skalin interface {
	GetContacts(*GetParams) ([]Contact, error)
	SaveContact(Contact) (*Contact, error)
	UpdateContact(Contact) (*Contact, error)
	CreateContactForCustomer(Contact, string) (*Contact, error)

	GetCustomers(*GetParams) ([]Customer, error)
	SaveCustomer(Customer) (*Customer, error)

	GetAgreements(*GetParams) ([]Agreement, error)
	SaveAgreement(Agreement) (*Agreement, error)
	UpdateAgreement(Agreement) (*Agreement, error)
	CreateAgreementForCustomer(Agreement, string) (*Agreement, error)

	SetLogger(logger logrus.FieldLogger)
}

type skalin struct {
	api API
}

func (s *skalin) SetLogger(logger logrus.FieldLogger) {
	s.api.SetLogger(logger)
}

func New(clientId, clientApiId, clientApiSecret string) (Skalin, error) {
	//format string with parameter
	body := fmt.Sprintf(`{"client_id": "%s", "client_secret": "%s", "grant_type": "client_credentials", "audience":"https://api.skalin.io/"}`, clientApiId, clientApiSecret)
	skalinApi := new(SkalinAPI)
	response, responseBytes, err := skalinApi.PostData(SKALIN_AUTH_URL, "application/json", nil, []byte(body), nil, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("error=%s; httpCode=%d", err, response.StatusCode)
	}
	var data map[string]interface{}
	err = json.Unmarshal(responseBytes, &data)
	if err != nil {
		return nil, fmt.Errorf("error=%s; httpCode=%d", err, response.StatusCode)
	}

	logrus.Infof("%s", data)
	skalin := &skalin{
		api: skalinApi.WithClientID(clientId).WithToken(data["access_token"].(string)),
	}
	// set default logger (but can be replace by another one)
	skalin.SetLogger(Log)
	return skalin, nil
}
