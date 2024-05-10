package channelBalanceService

type CreateInput struct {
	AccountID string   `json:"accountId" validate:"required"`
	ChannelID string   `json:"channelId" validate:"required"`
	Amount    *float64 `json:"amount,omitempty" validate:"omitempty,gte=0"`
	Currency  *string  `json:"currency,omitempty" validate:"omitempty,iso4217"`
}
type UpdateInput struct {
	ID       string   `json:"id"`
	Amount   *float64 `json:"amount,omitempty" validate:"omitempty,gte=0"`
	Currency *string  `json:"currency,omitempty" validate:"omitempty,iso4217"`
}
