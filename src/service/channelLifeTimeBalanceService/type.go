package channelLifeTimeBalanceService

type CreateInput struct {
	AccountID   string  `json:"accountId" validate:"required"`
	ChannelID   string  `json:"channelId" validate:"required"`
	TotalAmount float64 `json:"totalAmount" validate:"required,gte=0"`
	Currency    *string `json:"currency,omitempty" validate:"omitempty,iso4217"`
	TotalHour   float64 `json:"totalHour" validate:"required,gte=0"`
	TotalAd     float64 `json:"totalAd" validate:"required,gte=0"`
}
type UpdateInput struct {
	ID          string   `json:"id" validate:"required"`
	TotalAmount *float64 `json:"totalAmount,omitempty" validate:"omitempty,gte=0"`
	Currency    *string  `json:"currency,omitempty" validate:"omitempty,iso4217"`
	TotalHour   *float64 `json:"totalHour,omitempty" validate:"omitempty,gte=0"`
	TotalAd     *float64 `json:"totalAd,omitempty" validate:"omitempty,gte=0"`
}
