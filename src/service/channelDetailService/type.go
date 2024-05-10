package channelDetailService

import (
	"adtec/backend/src/service"
	"time"
)

type CreateInput struct {
	ChannelID       string     `json:"channelId" validate:"required"`
	SubCount        *int       `json:"subCount,omitempty" validate:"omitempty,gte=0"`
	AveragePostView *int       `json:"averagePostView,omitempty" validate:"omitempty,gte=0"`
	PostToSubRatio  *float64   `json:"postToSubRatio,omitempty" validate:"omitempty,gte=0"`
	PostFrequency   *float64   `json:"postFrequency,omitempty" validate:"omitempty,gte=0"`
	LastPostID      *string    `json:"lastPostId,omitempty"`
	LastPost        *time.Time `json:"lastPost,omitempty" validate:"omitempty,ValidateTime"`
	CollectedDate   *time.Time `json:"collectedDate,omitempty" validate:"omitempty,ValidateTime"`
}
type UpdateInput struct {
	ChannelID       string     `json:"channelId" validate:"required"`
	SubCount        *int       `json:"subCount,omitempty" validate:"omitempty,gte=0"`
	AveragePostView *int       `json:"averagePostView,omitempty" validate:"omitempty,gte=0"`
	PostToSubRatio  *float64   `json:"postToSubRatio,omitempty" validate:"omitempty,gte=0"`
	PostFrequency   *float64   `json:"postFrequency,omitempty" validate:"omitempty,gte=0"`
	LastPostID      *string    `json:"lastPostId,omitempty"`
	LastPost        *time.Time `json:"lastPost,omitempty" validate:"omitempty,ValidateTime"`
	CollectedDate   *time.Time `json:"collectedDate,omitempty" validate:"omitempty,ValidateTime"`
}
type FindManyInput struct {
	Filter *service.FilterData `json:"filter"`
}
