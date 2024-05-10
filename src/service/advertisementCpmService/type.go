package advertisementCpmService

type CreateInput struct {
	AdvertisementID string  `json:"advertisementId" validate:"required"`
	Rate            float64 `json:"Rate" validate:"required"`
	ChannelShare    float64 `json:"channelShare" validate:"required"`
	ProviderShare   float64 `json:"providerShare" validate:"required"`
	TotalBudget     float64 `json:"totalBudget" validate:"required"`
	MaxLifeCycle    int     `json:"maxLifeCycle" validate:"required"`
	RequiredViews   int     `json:"requiredViews" validate:"required"`
}
type UpdateInput struct {
	AdvertisementID string   `json:"advertisementId" validate:"required"`
	Rate            *float64 `json:"Rate,omitempty"`
	ChannelShare    *float64 `json:"channelShare,omitempty"`
	ProviderShare   *float64 `json:"providerShare,omitempty"`
	TotalBudget     *float64 `json:"totalBudget,omitempty"`
	MaxLifeCycle    *int     `json:"maxLifeCycle,omitempty"`
	RequiredViews   *int     `json:"requiredViews,omitempty"`
}
