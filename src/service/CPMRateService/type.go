package CPMRateService

type CreateInput struct {
	ChannelID    string   `json:"channelId" validate:"required"`
	Active       *bool    `json:"active,omitempty" validate:"omitempty,boolean"`
	Cpm          *float64 `json:"CPM,omitempty" validate:"omitempty,gte=0"`
	MinCPMVolume *float64 `json:"minCPMVolume,omitempty" validate:"omitempty,gte=0"`
}

type UpdateInput struct {
	ChannelID    string   `json:"channelId" validate:"required"`
	Active       *bool    `json:"active,omitempty" validate:"omitempty,boolean"`
	Cpm          *float64 `json:"CPM,omitempty" validate:"omitempty,gte=0"`
	MinCPMVolume *float64 `json:"minCPMVolume,omitempty" validate:"omitempty,gte=0"`
}
