package channelContentLanguageService

import "adtec/backend/src/service"

type ChannelContentLanguageInput struct {
	ChannelID string `json:"channelId" validate:"required"`
	Language  string `json:"language" validate:"required"`
}
type FindManyInput struct {
	ChannelId *string             `json:"channelId"`
	Language  *string             `json:"language"`
	Filter    *service.FilterData `json:"filter"`
}
