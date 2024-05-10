package withdrawalChannelLinkService

import "adtec/backend/src/service"

type WithdrawalChannelLinkInput struct {
	WithdrawalID     string `json:"withdrawalId" validate:"required"`
	ChannelBalanceID string `json:"channelBalanceId" validate:"required"`
}
type WithdrawalChannelLinkFilter struct {
	ChannelBalanceID *string             `json:"channelBalanceId,omitempty"`
	Filter           *service.FilterData `json:"filter,omitempty"`
}
