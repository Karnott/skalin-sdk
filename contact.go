package skalinsdk

import (
	"time"
)

type Contact struct {
	Id             string     `json:"id,omitempty"`
	CustomerId     *string    `json:"customerId,omitempty"` // correspond to the customer Id
	Customer       *string    `json:"customer,omitempty"`   // correspond to the customer refId
	RefId          string     `json:"refId,omitempty"`
	Email          string     `json:"email,omitempty"`
	FirstName      string     `json:"firstName,omitempty"`
	LastName       string     `json:"lastName,omitempty"`
	Phone          string     `json:"phone,omitempty"`
	Tags           []string   `json:"tags,omitempty"`
	LastActivityTs *time.Time `json:"lastActivityTs,omitempty"`
}

const (
	SAVE_CONTACT_PATH = "/contacts"
)

func (s *skalin) SaveContact(contact Contact) (*Contact, error) {
	return save(s, SAVE_CONTACT_PATH, contact)
}

func (s *skalin) GetContacts() ([]Contact, error) {
	return getEntities[[]Contact](s, SAVE_CONTACT_PATH)
}
