package hourlyRateService

type CreateInput struct {
	ChannelID       string   `json:"channelId" validate:"required"`
	Active          *bool    `json:"active,omitempty" validate:"omitempty,boolean"`
	HourlyRate      *float64 `json:"hourlyRate,omitempty" validate:"omitempty,gte=0"`
	MinHourlyVolume *float64 `json:"minHourlyVolume,omitempty" validate:"omitempty,gte=0"`
	MaxHourlyVolume *float64 `json:"maxHourlyVolume,omitempty" validate:"omitempty,gte=0"`
}

type UpdateInput struct {
	ID              string   `json:"id" validate:"required"`
	Active          *bool    `json:"active,omitempty" validate:"omitempty,boolean"`
	HourlyRate      *float64 `json:"hourlyRate,omitempty" validate:"omitempty,gte=0"`
	MinHourlyVolume *float64 `json:"minHourlyVolume,omitempty" validate:"omitempty,gte=0"`
	MaxHourlyVolume *float64 `json:"maxHourlyVolume,omitempty" validate:"omitempty,gte=0"`
}
