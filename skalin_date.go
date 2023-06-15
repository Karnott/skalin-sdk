package skalinsdk

import (
	"encoding/json"
	"strings"
	"time"
)

// need SkalinDate to unmarshal date from skalin API
// because format is "YYYY-MM-DD"
type SkalinDate time.Time

func (s *SkalinDate) UnmarshalJSON(b []byte) error {
	var stringTime = strings.Trim(string(b), "\"")
	t, err := time.Parse(time.DateOnly, stringTime)
	if err != nil {
		t, err = time.Parse(time.RFC3339, stringTime)
		if err != nil {
			return err
		}
	}
	*s = SkalinDate(t)
	return nil
}
func (s SkalinDate) MarshalJSON() ([]byte, error) {
	stringDate := time.Time(s).Format(time.DateOnly)
	return json.Marshal(stringDate)
}
