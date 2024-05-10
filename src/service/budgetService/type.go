package budgetService

type CreateInput struct {
	CampaignID     string   `json:"campaignId" validate:"required"`
	InitialAmount  float64  `json:"initialAmount,omitempty" validate:"omitempty, gte=0"`
	UsedAmount     *float64 `json:"usedAmount,omitempty" validate:"omitempty, gte=0,ltefield=InitialAmount"`
	RefundedAmount *float64 `json:"refundedAmount,omitempty" validate:"omitempty, gte=0,ltefield=InitialAmount"`
	Currency       *string  `json:"currency,omitempty"`
}

type UpdateInput struct {
	CampaignID     string   `json:"campaignId" validate:"required"`
	InitialAmount  *float64 `json:"initialAmount,omitempty" validate:"omitempty, gte=0"`
	UsedAmount     *float64 `json:"usedAmount,omitempty" validate:"omitempty, gte=0,ltefield=InitialAmount"`
	RefundedAmount *float64 `json:"refundedAmount,omitempty" validate:"omitempty, gte=0,ltefield=InitialAmount"`
	Currency       *string  `json:"currency,omitempty" validate:"omitempty,iso4217"`
}
