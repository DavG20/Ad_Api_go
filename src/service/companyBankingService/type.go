package companyBankingService

import (
	"adtec/backend/src/service"
)

type CreateInput struct {
	CompanyID      string `json:"companyId" validate:"required"`
	BankId         string `json:"bankId" validate:"required"`
	FullNameOnBank string `json:"fullNameOnBank" validate:"required,min=4"`
	BankAccount    string `json:"bankAccount" validate:"required,min=10"` //Todo: validate bank account , I don't know the format
	Currency       string `json:"currency,omitempty" validate:"required,iso4217"`
}

type FindManyInput struct {
	CompanyID *string             `json:"companyId,omitempty"`
	Filter    *service.FilterData `json:"filter"`
	BankId    *string             `json:"bankId"`
}
