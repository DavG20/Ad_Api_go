package accountBankingService

import (
	"adtec/backend/src/service"
)

type CreateInput struct {
	AccountId      string `json:"accountId" validate:"required"`
	BankId         string `json:"bankId" validate:"required"`
	FullNameOnBank string `json:"fullNameOnBank" validate:"required"`
	BankAccount    string `json:"bankAccount" validate:"required"`
	Currency       string `json:"currency,omitempty" validate:"required,iso4217"`
}

type FindManyInput struct {
	AccountId *string             `json:"accountId"`
	Filter    *service.FilterData `json:"filter"`
	BankId    *string             `json:"bankId"`
}
