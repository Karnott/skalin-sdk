package skalinsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type GenericResponse[T Customer | Contact | []Customer | []Contact] struct {
	Status string `json:"status"`
	Data   T      `json:"data"`
}

type GetParams struct {
	Page    *int
	Size    *int
	Sort    *string
	Filters map[string]interface{}
}

func mapFiltersToString(filters map[string]interface{}) string {
	stringFilters := make([]string, 0)
	for key, value := range filters {
		stringFilters = append(stringFilters, fmt.Sprintf("%v:%v", key, value))
	}
	return strings.Join(stringFilters, ",")
}

func buildQueryParamsFromGetParams(params *GetParams) *url.Values {
	if params == nil {
		return nil
	}
	queryParams := &url.Values{}
	if params.Page != nil {
		queryParams.Add("page", fmt.Sprintf("%v", *params.Page))
	}
	if params.Size != nil {
		queryParams.Add("size", fmt.Sprintf("%v", *params.Size))
	}
	if params.Sort != nil {
		queryParams.Add("sort", fmt.Sprintf("%v", *params.Sort))
	}
	if params.Filters != nil {
		queryParams.Add("filters", mapFiltersToString(params.Filters))
	}
	return queryParams
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

func update[T Customer | Contact](s *skalin, path string, entity T) (*T, error) {
	url := BuildUrl(path)
	jsonEntity, err := json.Marshal(entity)
	if err != nil {
		return nil, err
	}
	_, bodyResp, err := s.api.PatchData(
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
		return nil, fmt.Errorf("error to unmarshal entity for update [%v] response: %w", path, err)
	}
	return &jsonResp.Data, nil
}

func getEntities[T []Customer | []Contact](s *skalin, path string, queryParams *url.Values) (T, error) {
	url := BuildUrl(path)
	_, bodyResp, err := s.api.GetData(
		url,
		jsonContentType,
		nil,
		nil,
		queryParams,
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
