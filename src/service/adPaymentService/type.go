package adPaymentService

import (
	"adtec/backend/src/service"
	"time"
)

type CreateInput struct {
	AdvertisementID string     `json:"advertisementId" validate:"required"`
	AccountID       string     `json:"accountId" validate:"required"`
	CampaignID      string     `json:"campaignId" validate:"required"`
	ChannelID       string     `json:"channelId" validate:"required"`
	Amount          *float64   `json:"amount,omitempty" validate:"omitempty, gte=0"`
	Currency        *string    `json:"currency,omitempty" validate:"omitempty,iso4217"`
	CompletionTime  *time.Time `json:"completionTime,omitempty" validate:"omitempty,ValidateTime"`
}
type UpdateInput struct {
	ID             string     `json:"id" validate:"required"`
	Amount         *float64   `json:"amount,omitempty" validate:"omitempty,gte=0"`
	Currency       *string    `json:"currency,omitempty" validate:"omitempty,iso4217"`
	CompletionTime *time.Time `json:"completionTime,omitempty" validate:"omitempty,ValidateTime"`
}
type AdPaymentFilter struct {
	AdvertisementID *string                    `json:"advertisementId,omitempty"`
	CampaignID      *string                    `json:"campaignId,omitempty"`
	ChannelID       *string                    `json:"channelId,omitempty"`
	Amount          *service.MinMaxFilterInput `json:"amount,omitempty"`
	Currency        *string                    `json:"currency,omitempty"`
	Filter          *service.FilterData        `json:"filter,omitempty"`
}
