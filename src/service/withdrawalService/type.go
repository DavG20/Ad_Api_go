package withdrawalService

import (
	"adtec/backend/src/service"
	"time"
)

type CreateInput struct {
	AccountID     string                 `json:"accountId" validate:"required"`
	AccountBankID string                 `json:"accountBankId" validate:"required"`
	TotalAmount   float64                `json:"totalAmount" validate:"required,gte=0"`
	Currency      *string                `json:"currency,omitempty" validate:"omitempty,iso4217"`
	FundingKey    *string                `json:"fundingKey,omitempty"`
	Method        *string                `json:"method,omitempty"`
	Reference     *string                `json:"reference,omitempty"`
	PaymentDate   *time.Time             `json:"paymentDate,omitempty" validate:"omitempty,ValidateTime"`
	Status        *service.FundingStatus `json:"status,omitempty"`
	Extra         interface{}            `json:"extra,omitempty"`
	Log           interface{}            `json:"log,omitempty"`
}
type UpdateInput struct {
	ID          string                 `json:"id" validate:"required"`
	TotalAmount *float64               `json:"totalAmount,omitempty" validate:"omitempty,gte=0"`
	Currency    *string                `json:"currency,omitempty" validate:"omitempty,iso4217"`
	FundingKey  *string                `json:"fundingKey,omitempty"`
	Method      *string                `json:"method,omitempty"`
	Reference   *string                `json:"reference,omitempty"`
	PaymentDate *time.Time             `json:"paymentDate,omitempty" validate:"omitempty,ValidateTime"`
	Status      *service.FundingStatus `json:"status,omitempty"`
	Extra       *interface{}           `json:"extra,omitempty"`
	Log         *interface{}           `json:"log,omitempty"`
}
type WithdrawalFilter struct {
	AccountID     *string                    `json:"accountId,omitempty"`
	AccountBankID *string                    `json:"accountBankId,omitempty"`
	Status        []service.FundingStatus    `json:"status,omitempty"`
	TotalAmount   *service.MinMaxFilterInput `json:"totalAmount,omitempty"`
	Currency      *string                    `json:"currency,omitempty"`
	Filter        *service.FilterData        `json:"filter,omitempty"`
}
