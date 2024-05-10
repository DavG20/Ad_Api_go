package channelCollectedAdLinkService

import "adtec/backend/src/service"

type ChannelCollectedAdLinkInput struct {
	ChannelBalanceID string `json:"channelBalanceId" validate:"required"`
	AdPaymentID      string `json:"adPaymentId" validate:"required"`
}
type ChannelCollectedAdLinkFilter struct {
	ChannelBalanceID *string             `json:"channelBalanceId,omitempty"`
	Filter           *service.FilterData `json:"filter"`
}
