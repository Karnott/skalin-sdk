package skalinsdk

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Skalin interface {
	GetContacts() ([]Contact, error)
	SaveContact(Contact) (*Contact, error)

	SaveCustomer(Customer) (*Customer, error)
}

type skalin struct {
	api API
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
	return &skalin{
		api: skalinApi.WithClientID(clientId).WithToken(data["access_token"].(string)),
	}, nil
}
