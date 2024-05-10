package companyService

type CreateInput struct {
	CreatorID  string  `json:"creatorId" validate:"required"`
	Name       string  `json:"name" validate:"required,min=3,max=50"`
	UniqueName string  `json:"uniqueName" validate:"required,min=3,max=20"` // todo: check the validation of unique name
	TinNumber  *string `json:"tinNumber,omitempty" `
	VatNumber  *string `json:"vatNumber,omitempty"`
}
type UpdateInput struct {
	ID         string  `json:"id" validate:"required"`
	Name       *string `json:"name,omitempty" validate:"omitempty,min=3,max=50"`
	UniqueName *string `json:"uniqueName,omitempty" validate:"omitempty,min=3,max=20"` // todo: unique
	TinNumber  *string `json:"tinNumber,omitempty"`
	VatNumber  *string `json:"vatNumber,omitempty"`
}
