package skalinsdk

import (
	"encoding/json"
	"fmt"
	"time"
)

type CustomAttributes map[string]interface{}

type Contact struct {
	Id               string           `json:"id,omitempty"`
	CustomerId       *string          `json:"customerId,omitempty"` // correspond to the customer Id
	Customer         *string          `json:"customer,omitempty"`   // correspond to the customer refId
	RefId            string           `json:"refId,omitempty"`
	Email            string           `json:"email,omitempty"`
	FirstName        string           `json:"firstName,omitempty"`
	LastName         string           `json:"lastName,omitempty"`
	Phone            string           `json:"phone,omitempty"`
	Tags             []string         `json:"tags,omitempty"`
	LastActivityTs   *time.Time       `json:"lastActivityTs,omitempty"`
	CustomAttributes CustomAttributes `json:"-"`
}

// need custom MarshalJSON to merge custom attributes with contact
// see the doc of skalin
// for now, no need to create a custom UnmarshalJson because skalin API does not return custom attributes
func (c Contact) MarshalJSON() ([]byte, error) {
	type Alias Contact // prevent stack overflow
	if c.CustomAttributes == nil {
		return json.Marshal((Alias)(c))
	}
	customAttributes, err := json.Marshal(c.CustomAttributes)
	if err != nil {
		return nil, fmt.Errorf("error to marshal custom attributes: %v", err)
	}
	jsonContact, err := json.Marshal((Alias)(c)) // prevent stack overflow
	if err != nil {
		return nil, fmt.Errorf("error to marshal contact: %v", err)
	}
	var r map[string]interface{}
	err = json.Unmarshal(jsonContact, &r)
	if err != nil {
		return nil, fmt.Errorf("error to unmarshal contact: %v", err)
	}
	err = json.Unmarshal(customAttributes, &r)
	if err != nil {
		return nil, fmt.Errorf("error to unmarshal custom attributes: %v", err)
	}
	return json.Marshal(r)
}

const (
	SAVE_CONTACT_PATH            = "/contacts"
	CREATE_CUSTOMER_CONTACT_PATH = "/customers/%v/contacts"
)

func (s *skalin) SaveContact(contact Contact) (*Contact, error) {
	return save(s, SAVE_CONTACT_PATH, contact)
}

func (s *skalin) GetContacts(params *GetParams) ([]Contact, error) {
	return getEntities[[]Contact](s, SAVE_CONTACT_PATH, buildQueryParamsFromGetParams(params))
}

func (s *skalin) CreateContactForCustomer(contact Contact, customerId string) (*Contact, error) {
	return save(s, fmt.Sprintf(CREATE_CUSTOMER_CONTACT_PATH, customerId), contact)
}
