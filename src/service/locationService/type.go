package locationService

type CreateInput struct {
	Country    string  `json:"country" validate:"required,iso3166_1_alpha3"`
	State      string  `json:"state" validate:"required"`
	City       string  `json:"city" validate:"required"`
	Address    *string `json:"address,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
}
type UpdateInput struct {
	ID         string  `json:"id" validate:"required"`
	Country    *string `json:"country,omitempty" validate:"omitempty,iso3166_1_alpha3"`
	State      *string `json:"state,omitempty"`
	City       *string `json:"city,omitempty"`
	Address    *string `json:"address,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
}
