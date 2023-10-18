package skalinsdk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetTags(t *testing.T) {
	var fakeTagsData = `[
  {
    "id": "1",
		"name": "tag 1",
		"type": "CUSTOM", 
		"entity": "CONTACT",
		"color": "#000000"
  }
  ]`
	var fakeTagsResponse = `
	{
		"status":"success",
		"data": ` + fakeTagsData + `
	}`

	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		expectedBody := []byte(fakeTagsResponse)
		_expectedTags := []byte(fakeTagsData)
		var expectedTags []Tag
		err := json.Unmarshal(_expectedTags, &expectedTags)
		if !assert.NoError(t, err) {
			return
		}

		mockApi.On(
			"send",
			http.MethodGet,
			BuildUrl(GET_TAGS),
			jsonContentType,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			http.StatusOK,
		).Return(nil, expectedBody, nil)

		skalinAPI := &skalinAPI{api: mockApi}
		tags, err := skalinAPI.GetTags(nil)
		mockApi.AssertExpectations(t)
		if !assert.NoError(t, err) || !assert.Equal(t, len(expectedTags), len(tags)) {
			return
		}
		expectedTag := expectedTags[0]
		tag := tags[0]
		assert.Equal(t, expectedTag.Id, tag.Id)
		assert.Equal(t, expectedTag.Name, tag.Name)
	})

	t.Run("With error", func(t *testing.T) {
		mockApi := new(MockAPI)
		mockApi.On(
			"send",
			http.MethodGet,
			BuildUrl(GET_TAGS),
			jsonContentType,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, fmt.Errorf("Status code != %v: %v", http.StatusInternalServerError, http.StatusOK))

		skalinAPI := &skalinAPI{api: mockApi}
		tags, err := skalinAPI.GetTags(nil)
		mockApi.AssertExpectations(t)
		if !assert.Error(t, err) {
			return
		}
		assert.Nil(t, tags)
	})

	t.Run("Call API", func(t *testing.T) {
		if GetSkalinAppClientID() == "" || GetSkalinClientApiID() == "" || GetSkalinClientApiSecret() == "" {
			return
		}
		skalinApi, err := New(GetSkalinAppClientID(), GetSkalinClientApiID(), GetSkalinClientApiSecret())
		if !assert.NoError(t, err) {
			return
		}
		tags, err := skalinApi.GetTags(nil)
		if !assert.NoError(t, err) {
			return
		}
		if len(tags) == 0 {
			return
		}
		for _, tag := range tags {
			assert.NotEqual(t, "", tag.Id)
			assert.NotEqual(t, "", tag.Name)
		}
	})
}

func TestGetTagByID(t *testing.T) {
	var fakeTagData = `{
    "id": "1",
		"name": "tag 1",
		"type": "CUSTOM", 
		"entity": "CONTACT",
		"color": "#000000"
  }`
	var fakeTagsResponse = `
	{
		"status":"success",
		"data": ` + fakeTagData + `
	}`

	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		expectedBody := []byte(fakeTagsResponse)
		_expectedTags := []byte(fakeTagData)
		var expectedTag Tag
		err := json.Unmarshal(_expectedTags, &expectedTag)
		if !assert.NoError(t, err) {
			return
		}

		mockApi.On(
			"send",
			http.MethodGet,
			BuildUrl(fmt.Sprintf(GET_TAG_BY_ID, expectedTag.Id)),
			jsonContentType,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			http.StatusOK,
		).Return(nil, expectedBody, nil)

		skalinAPI := &skalinAPI{api: mockApi}
		tag, err := skalinAPI.GetTagByID(expectedTag.Id)
		mockApi.AssertExpectations(t)
		if !assert.NoError(t, err) || !assert.NotNil(t, tag) {
			return
		}
		assert.Equal(t, expectedTag.Id, tag.Id)
		assert.Equal(t, expectedTag.Name, tag.Name)
	})

	t.Run("With error", func(t *testing.T) {
		mockApi := new(MockAPI)
		mockApi.On(
			"send",
			http.MethodGet,
			BuildUrl(fmt.Sprintf(GET_TAG_BY_ID, "id")),
			jsonContentType,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, fmt.Errorf("Status code != %v: %v", http.StatusInternalServerError, http.StatusOK))

		skalinAPI := &skalinAPI{api: mockApi}
		tag, err := skalinAPI.GetTagByID("id")
		mockApi.AssertExpectations(t)
		if !assert.Error(t, err) {
			return
		}
		assert.Nil(t, tag)
	})

	t.Run("Call API", func(t *testing.T) {
		if GetSkalinAppClientID() == "" || GetSkalinClientApiID() == "" || GetSkalinClientApiSecret() == "" {
			return
		}
		skalinApi, err := New(GetSkalinAppClientID(), GetSkalinClientApiID(), GetSkalinClientApiSecret())
		if !assert.NoError(t, err) {
			return
		}
		tags, err := skalinApi.GetTags(nil)
		if !assert.NoError(t, err) {
			return
		}
		if len(tags) == 0 {
			return
		}
		for _, tag := range tags {
			assert.NotEqual(t, "", tag.Id)
			assert.NotEqual(t, "", tag.Name)
			tagFromAPI, err := skalinApi.GetTagByID(tag.Id)
			if !assert.NoError(t, err) {
				return
			}
			assert.Equal(t, tag.Id, tagFromAPI.Id)
			assert.Equal(t, tag.Name, tagFromAPI.Name)
			break
		}
	})
}
