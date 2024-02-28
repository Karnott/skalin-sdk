package skalinsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
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
	UPDATE_CONTACT_PATH          = "/contacts/%v"
	CREATE_CUSTOMER_CONTACT_PATH = "/customers/%v/contacts"
)

// because in skalin API, many contact can have the same refId,
// only the first match will be updated (if the refId already exists)
func (s *skalinAPI) SaveContact(contact Contact) (*Contact, error) {
	return save(s, SAVE_CONTACT_PATH, contact)
}

func (s *skalinAPI) UpdateContact(contact Contact) (*Contact, error) {
	if contact.Id == "" {
		return nil, fmt.Errorf("contact id is empty")
	}
	// for now the API does not return the updated contact
	err := update(s, fmt.Sprintf(UPDATE_CONTACT_PATH, contact.Id), contact)
	if err != nil {
		return nil, err
	}
	return &contact, nil
}

func (s *skalinAPI) GetContacts(params *GetParams) ([]Contact, error) {
	return getEntities[[]Contact](s, SAVE_CONTACT_PATH, buildQueryParamsFromGetParams(params))
}

func (s *skalinAPI) CreateContactForCustomer(contact Contact, customerId string) (*Contact, error) {
	return save(s, fmt.Sprintf(CREATE_CUSTOMER_CONTACT_PATH, customerId), contact)
}

func (s *skalinAPI) DeleteContact(contact Contact) error {
	if contact.Id == "" {
		return fmt.Errorf("contact id is empty")
	}
	// for now the API does not return the updated agreement
	r, b, err := s.api.DeleteData(BuildUrl(fmt.Sprintf(UPDATE_CONTACT_PATH, contact.Id)), "", nil, nil, nil, http.StatusOK)
	s.api.GetLogger().Infof("Delete contact %+v: %v", r, string(b))
	if err != nil {
		return err
	}
	return nil
}
