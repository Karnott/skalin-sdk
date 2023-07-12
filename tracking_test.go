package skalinsdk

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func sPtr(s string) *string { return &s }

func TestHit(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		mockApi := new(MockAPI)
		mockApi.On(
			"send",
			http.MethodPost,
			SKALIN_HIT_URL,
			formURLEncodedContentType,
			mock.Anything,
			mock.Anything,
			mock.Anything,
			http.StatusOK,
		).Return(nil, nil, nil)

		skalinTracker := &skalinTracker{api: mockApi}
		_, _, err := skalinTracker.Hit(HitTrack{
			Action:    HitActionEvent,
			VisitorID: "1234567890123456",
			VisitID:   "1234567890123456",
			Identity: HitIdentity{
				ID: sPtr("test"),
			},
			Event: &HitEvent{
				Name:      "test",
				EventName: "test",
			},
		})

		assert.NoError(t, err)
	})
}
