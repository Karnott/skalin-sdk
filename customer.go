package skalinsdk

import (
	"encoding/json"
	"fmt"
	"time"
)

type Customer struct {
	Id               string           `json:"id,omitempty"`
	RefId            string           `json:"refId,omitempty"`
	Name             string           `json:"name,omitempty"`
	Stage            string           `json:"stage,omitempty"`
	Tags             []string         `json:"tags,omitempty"`
	LastActivityTs   *time.Time       `json:"lastActivityTs,omitempty"`
	CustomAttributes CustomAttributes `json:"-"`
}

type CustomerResponse struct {
	Status string   `json:"status"`
	Data   Customer `json:"data"`
}

// need custom MarshalJSON to merge custom attributes with contact
// see the doc of skalin
// for now, no need to create a custom UnmarshalJson because skalin API does not return custom attributes
func (c Customer) MarshalJSON() ([]byte, error) {
	type Alias Customer // prevent stack overflow
	if c.CustomAttributes == nil {
		return json.Marshal((Alias)(c))
	}
	customAttributes, err := json.Marshal(c.CustomAttributes)
	if err != nil {
		return nil, fmt.Errorf("error to marshal custom attributes: %v", err)
	}
	jsonContact, err := json.Marshal((Alias)(c)) // prevent stack overflow
	if err != nil {
		return nil, fmt.Errorf("error to marshal customer: %v", err)
	}
	var r map[string]interface{}
	err = json.Unmarshal(jsonContact, &r)
	if err != nil {
		return nil, fmt.Errorf("error to unmarshal customer: %v", err)
	}
	err = json.Unmarshal(customAttributes, &r)
	if err != nil {
		return nil, fmt.Errorf("error to unmarshal custom attributes: %v", err)
	}
	return json.Marshal(r)
}

const (
	SAVE_CUSTOMER_PATH = "/customers"
)

func (s *skalinAPI) SaveCustomer(customer Customer) (*Customer, error) {
	return save(s, SAVE_CUSTOMER_PATH, customer)
}

func (s *skalinAPI) GetCustomers(params *GetParams) ([]Customer, error) {
	return getEntities[[]Customer](s, SAVE_CUSTOMER_PATH, buildQueryParamsFromGetParams(params))
}
