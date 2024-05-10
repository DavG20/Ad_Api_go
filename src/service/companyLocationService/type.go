package companyLocationService

type CreateInput struct {
	CompanyID  string `json:"companyId" validate:"required"`
	LocationID string `json:"locationId" validate:"required"`
}
