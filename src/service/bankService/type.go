package bankService

import "adtec/backend/src/service"

type CreateInput struct {
	BankName string `json:"bankName" validate:"required"`
	BankCode string `json:"bankCode" validate:"required"`
}

type UpdateInput struct {
	ID       string  `json:"id" validate:"required"`
	BankName *string `json:"bankName,omitempty" validate:"omitempty"`
	BankCode *string `json:"bankCode,omitempty" validate:"omitempty"`
}
type FindManyInput struct {
	Filter *service.FilterData `json:"filter" validate:"omitempty"`
}
