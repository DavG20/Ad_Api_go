package advertisementResultService

import "time"

type CreateInput struct {
	AdvertisementID     string    `json:"advertisementId" validate:"required"`
	StartedAt           time.Time `json:"startedAt" validate:"required,ValidateTime"`
	Budget              *float64  `json:"budget,omitempty" validate:"omitempty,gte=0"`
	ProviderBudgetShare *float64  `json:"providerBudgetShare,omitempty"`
	ChannelBudgetShare  *float64  `json:"channelBudgetShare,omitempty"`
	TotalHour           *int      `json:"totalHour,omitempty"`
	TotalClick          *int      `json:"totalClick,omitempty"`
	TotalViews          *int      `json:"totalViews,omitempty"`
	TotalForward        *int      `json:"totalForward,omitempty"`
	TotalReaction       *int      `json:"totalReaction,omitempty"`
}

type UpdateInput struct {
	AdvertisementID     string     `json:"advertisementId" validate:"required"`
	StartedAt           *time.Time `json:"startedAt,omitempty" validate:"omitempty,ValidateTime"`
	Budget              *float64   `json:"budget,omitempty" validate:"omitempty,gte=0"`
	ProviderBudgetShare *float64   `json:"providerBudgetShare,omitempty"`
	ChannelBudgetShare  *float64   `json:"channelBudgetShare,omitempty"`
	TotalHour           *int       `json:"totalHour,omitempty"`
	TotalClick          *int       `json:"totalClick,omitempty"`
	TotalViews          *int       `json:"totalViews,omitempty"`
	TotalForward        *int       `json:"totalForward,omitempty"`
	TotalReaction       *int       `json:"totalReaction,omitempty"`
}
