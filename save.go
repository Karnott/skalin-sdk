package skalinsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type EntitiesGeneric interface {
	Customer | Contact | Agreement
}

type EntitySlice[T EntitiesGeneric] interface {
	~[]T // permit to define a core type to use `make` and `append` on generic (https://go.dev/ref/spec#Core_types)
}
type ResponseMetadata interface {
	PaginationMetadata
}

type PaginationMetadata struct {
	Pagination struct {
		Size  int `json:"size"`
		Page  int `json:"page"`
		Total int `json:"total"`
	} `json:"pagination"`
}

type GenericResponse[T EntitySlice[V] | EntitiesGeneric, V EntitiesGeneric, U ResponseMetadata] struct {
	Status   string `json:"status"`
	Data     T      `json:"data"`
	Metadata U      `json:"metadata,omitempty"`
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

func save[T EntitiesGeneric](s *skalin, path string, entity T) (*T, error) {
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
	var jsonResp GenericResponse[T, T, PaginationMetadata]
	err = json.Unmarshal(bodyResp, &jsonResp)
	if err != nil {
		return nil, fmt.Errorf("error to unmarshal entity for save [%v] response: %w", path, err)
	}
	return &jsonResp.Data, nil
}

func update[T EntitiesGeneric](s *skalin, path string, entity T) error {
	url := BuildUrl(path)
	jsonEntity, err := json.Marshal(entity)
	if err != nil {
		return err
	}
	_, _, err = s.api.PatchData(
		url,
		jsonContentType,
		nil,
		jsonEntity,
		nil,
		http.StatusOK,
	)

	// no need to unmarshal response
	// because it returns only `status: success`
	return err
}

func getEntitiesWithMetadata[T EntitySlice[V], V EntitiesGeneric, U PaginationMetadata](s *skalin, path string, queryParams *url.Values) (*GenericResponse[T, V, U], error) {
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
	var jsonResp GenericResponse[T, V, U]
	err = json.Unmarshal(bodyResp, &jsonResp)
	if err != nil {
		return nil, fmt.Errorf("error to unmarshal entity for get [%v] response: %w", path, err)
	}
	return &jsonResp, nil
}

func getEntities[T EntitySlice[V], V EntitiesGeneric](s *skalin, path string, queryParams *url.Values) (T, error) {
	isGetAll := false
	data := make(T, 0)
	if queryParams == nil {
		queryParams = &url.Values{}
	}
	for !isGetAll {
		jsonResp, err := getEntitiesWithMetadata[T](s, path, queryParams)
		if err != nil {
			return nil, err
		}
		data = append(data, jsonResp.Data...)
		page := jsonResp.Metadata.Pagination.Page
		queryParams.Set("page", strconv.Itoa(page+1))
		// need to continue to get data until the total is reached
		if len(data) >= jsonResp.Metadata.Pagination.Total {
			isGetAll = true
		}
	}
	return data, nil
}
