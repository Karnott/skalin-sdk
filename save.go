package skalinsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GenericResponse[T Customer | Contact | []Customer | []Contact] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

func save[T Customer | Contact](s *skalin, path string, entity T) (*T, error) {
	url := BuildUrl(path)
	jsonEntity, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}
	_, bodyResp, err := s.api.PostData(
		url,
		jsonContentType,
		nil,
		jsonEntity,
		nil,
		http.StatusOK,
	)

	if err != nil {
		return nil, err
	}
	var jsonResp GenericResponse[T]
	err = json.Unmarshal(bodyResp, &jsonResp)
	if err != nil {
		return nil, fmt.Errorf("error to unmarshal entity for save [%v] response: %w", path, err)
	}
	return &jsonResp.Data, nil
}

func getEntities[T []Customer | []Contact](s *skalin, path string) (T, error) {
	url := BuildUrl(path)
	_, bodyResp, err := s.api.GetData(
		url,
		jsonContentType,
		nil,
		nil,
		nil,
		http.StatusOK,
	)

	if err != nil {
		return nil, err
	}
	var jsonResp GenericResponse[T]
	err = json.Unmarshal(bodyResp, &jsonResp)
	if err != nil {
		return nil, fmt.Errorf("error to unmarshal entity for get [%v] response: %w", path, err)
	}
	return jsonResp.Data, nil
}
