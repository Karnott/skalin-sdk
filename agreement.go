package skalinsdk

import (
	"fmt"
	"net/http"
)

type Agreement struct {
	Id               string      `json:"id,omitempty"`
	CustomerId       *string     `json:"customerId,omitempty"` // correspond to the customer Id
	Customer         *string     `json:"customer,omitempty"`   // correspond to the customer refId
	RefId            string      `json:"refId,omitempty"`
	StartDate        *SkalinDate `json:"startDate,omitempty"`   // need to be at format `YYYY-MM-DD`
	EndDate          *SkalinDate `json:"endDate,omitempty"`     // need to be at format `YYYY-MM-DD`
	RenewalDate      *SkalinDate `json:"renewalDate,omitempty"` // need to be at format `YYYY-MM-DD`
	AutoRenew        bool        `json:"autoRenew,omitempty"`
	Engagement       *int        `json:"engagement,omitempty"` // need pointer because engagement value can be 0
	EngagementPeriod string      `json:"engagementPeriod,omitempty"`
	Notice           *int        `json:"notice,omitempty"` // need pointer because engagement value can be 0
	NoticePeriod     string      `json:"noticePeriod,omitempty"`
	Plan             string      `json:"plan,omitempty"`
	Type             string      `json:"type,omitempty"`
	Mrr              *int        `json:"mrr,omitempty"`
	Fee              *int        `json:"fee,omitempty"`
}

const (
	SAVE_AGREEMENT_PATH            = "/agreements"
	UPDATE_AGREEMENT_PATH          = "/agreements/%v"
	CREATE_CUSTOMER_AGREEMENT_PATH = "/customers/%v/agreements"
)

// because in skalin API, many agreement can have the same refId,
// only the first match will be updated (if the refId already exists)
func (s *skalinAPI) SaveAgreement(agreement Agreement) (*Agreement, error) {
	return save(s, SAVE_AGREEMENT_PATH, agreement)
}

func (s *skalinAPI) UpdateAgreement(agreement Agreement) (*Agreement, error) {
	if agreement.Id == "" {
		return nil, fmt.Errorf("agreement id is empty")
	}
	// for now the API does not return the updated agreement
	err := update(s, fmt.Sprintf(UPDATE_AGREEMENT_PATH, agreement.Id), agreement)
	if err != nil {
		return nil, err
	}
	return &agreement, nil
}

func (s *skalinAPI) GetAgreements(params *GetParams) ([]Agreement, error) {
	return getEntities[[]Agreement](s, SAVE_AGREEMENT_PATH, buildQueryParamsFromGetParams(params))
}

func (s *skalinAPI) CreateAgreementForCustomer(agreement Agreement, customerId string) (*Agreement, error) {
	return save(s, fmt.Sprintf(CREATE_CUSTOMER_AGREEMENT_PATH, customerId), agreement)
}

func (s *skalinAPI) DeleteAgreement(agreement Agreement) error {
	if agreement.Id == "" {
		return fmt.Errorf("agreement id is empty")
	}
	// for now the API does not return the updated agreement
	r, b, err := s.api.DeleteData(BuildUrl(fmt.Sprintf(UPDATE_AGREEMENT_PATH, agreement.Id)), "", nil, nil, nil, http.StatusOK)
	s.api.GetLogger().Infof("Delete agreement %+v: %v", r, string(b))
	if err != nil {
		return err
	}
	return nil
}
