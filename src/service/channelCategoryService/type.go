package channelCategoryService

import "adtec/backend/src/service"

type ChannelCategoryInput struct {
	ChannelID string `json:"channelId" validate:"required"`
	Category  string `json:"category" validate:"required"`
}
type FindManyInput struct {
	ChannelId *string             `json:"channelId,omitempty"`
	Category  *string             `json:"category,omitempty"`
	Filter    *service.FilterData `json:"filter"`
}
