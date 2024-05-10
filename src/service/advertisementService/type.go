package advertisementService

import "adtec/backend/src/service"

type CreateInput struct {
	CampaignID string            `json:"campaignId" validate:"required"`
	ContentID  string            `json:"contentId" validate:"required"`
	ChannelID  string            `json:"channelId" validate:"required"`
	MessageID  *int              `json:"messageId,omitempty" `
	Status     *service.AdStatus `json:"status,omitempty" validate:"oneof=Pending Running Closed Rejected"`
}
type UpdateInput struct {
	ID     string            `json:"id" validate:"required"`
	Status *service.AdStatus `json:"status,omitempty" validate:"oneof=Pending Running Closed Rejected"`
}
type AdvertisementFilter struct {
	ChannelID          *string                    `json:"channelId,omitempty"`
	CompanyID          *string                    `json:"companyId,omitempty"`
	Status             []service.AdStatus         `json:"status,omitempty"`
	Rate               *service.MinMaxFilterInput `json:"rate,omitempty"`
	ChannelBudgetShare *service.MinMaxFilterInput `json:"channelBudgetShare,omitempty"`
	RequiredViews      *service.MinMaxFilterInput `json:"requiredViews,omitempty"`
	Views              *service.MinMaxFilterInput `json:"Views,omitempty"`
	Hours              *service.MinMaxFilterInput `json:"Hours,omitempty"`
	Filter             *service.FilterData        `json:"filter,omitempty"`
}
