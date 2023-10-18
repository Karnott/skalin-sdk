package skalinsdk

import "fmt"

type Tag struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Entity string `json:"entity"`
	Color  string `json:"color"`
}

const (
	GET_TAGS      = "/tags"
	GET_TAG_BY_ID = "/tags/%v"
)

func (s *skalinAPI) GetTags(params *GetParams) ([]Tag, error) {
	return getEntities[[]Tag](s, GET_TAGS, buildQueryParamsFromGetParams(params))
}

func (s *skalinAPI) GetTagByID(id string) (*Tag, error) {
	return getEntity[Tag](s, fmt.Sprintf(GET_TAG_BY_ID, id), nil)
}
