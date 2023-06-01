package skalinsdk

import (
	"time"
)

type Customer struct {
	Id             string     `json:"id,omitempty"`
	RefId          string     `json:"refId,omitempty"`
	Name           string     `json:"name,omitempty"`
	Stage          string     `json:"stage,omitempty"`
	Tags           []string   `json:"tags,omitempty"`
	LastActivityTs *time.Time `json:"lastActivityTs,omitempty"`
}

type CustomerResponse struct {
	Status string   `json:"status"`
	Data   Customer `json:"data"`
}

const (
	SAVE_CUSTOMER_PATH = "/customers"
)

func (s *skalin) SaveCustomer(customer Customer) (*Customer, error) {
	return save(s, SAVE_CUSTOMER_PATH, customer)
}
