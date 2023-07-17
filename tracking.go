package skalinsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-playground/validator/v10"
)

type HitAction string

const HitActionUserIdendity HitAction = "ui"
const HitActionEvent HitAction = "ev"

type HitIdentity struct {
	ID    *string `json:"id,omitempty" validate:"required_without=Email"`
	Email *string `json:"email,omitempty" validate:"required_without=ID,omitempty,email"`
}

type HitEvent struct {
	Name      string `json:"name" validate:"required"`
	EventName string `json:"event_name" validate:"required"`
}

type HitTrack struct {
	Action        HitAction   `validate:"required"`
	VisitorID     string      `validate:"required,len=16"`
	VisitID       string      `validate:"required,len=16"`
	Identity      HitIdentity `validate:"required"`
	Event         *HitEvent   `validate:"required_if=Action ev"` // mandatory if action is event
	EventID       *string     `validate:"omitempty,len=16"`
	CustomerID    *string
	Ts            *time.Time
	URL           *string
	CustomHeaders map[string][]string
}

func (a skalinTracker) Hit(ht HitTrack) (*http.Response, []byte, error) {
	validator := validator.New()
	err := validator.Struct(ht)
	if err != nil {
		return nil, nil, err
	}

	if a.api.GetClientID() == nil {
		return nil, nil, fmt.Errorf("client_id is not set")
	}

	data := url.Values{}
	data.Set("rec", "1")
	data.Set("action", string(ht.Action))
	data.Set("visitor_id", ht.VisitorID)
	data.Set("visit_id", ht.VisitID)

	data.Set("client_id", *a.api.GetClientID())

	i, err := json.Marshal(ht.Identity)
	if err != nil {
		return nil, nil, err
	}
	data.Set("identity", string(i))

	if ht.Event != nil {
		e, err := json.Marshal(ht.Event)
		if err != nil {
			return nil, nil, err
		}
		data.Set("event", string(e))
	}

	if ht.CustomerID != nil {
		data.Set("customer_id", *ht.CustomerID)
	}

	if ht.EventID != nil {
		data.Set("event_id", *ht.EventID)
	}

	if ht.Ts != nil {
		data.Set("localtime", ht.Ts.Local().Format(time.TimeOnly))
	}

	if ht.URL != nil {
		data.Set("url", *ht.URL)
		if ht.Action == HitActionEvent && ht.Ts != nil {
			data.Set("ts", ht.Ts.UTC().Format("2006-01-02T15:04:05"))
		}
	}

	return a.api.PostData(
		SKALIN_HIT_URL,
		formURLEncodedContentType,
		ht.CustomHeaders,
		nil,
		&data,
		http.StatusOK,
	)
}
