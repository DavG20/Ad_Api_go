package fundingService

import (
	"adtec/backend/src/service"
	"time"
)

type CreateInput struct {
	CompanyID    string                 `json:"companyId" validate:"required"`
	Amount       float64                `json:"amount" validate:"required,gte=0"`
	Currency     string                 `json:"currency" validate:"required,iso4217"`
	Key          string                 `json:"key" validate:"required"`
	Method       *string                `json:"method,omitempty" `
	FundingTxRef *string                `json:"fundingTxRef,omitempty"`
	Reference    *string                `json:"reference,omitempty"`
	PaymentDate  *time.Time             `json:"paymentDate,omitempty"`
	RedirectURL  *string                `json:"redirectUrl,omitempty"`
	Status       *service.FundingStatus `json:"status,omitempty"` //TODO: add validation
	Extra        interface{}            `json:"extra"`
	Log          interface{}            `json:"log"`
	Tax          string                 `json:"tax" validate:"required"`
}
type UpdateInput struct {
	ID           string                 `json:"id" validate:"required"`
	Amount       *float64               `json:"amount,omitempty" validate:"omitempty,gte=0"`
	Currency     *string                `json:"currency,omitempty" validate:"omitempty,iso4217"`
	Key          *string                `json:"key,omitempty"`
	Method       *string                `json:"method,omitempty"`
	FundingTxRef *string                `json:"fundingTxRef,omitempty"`
	Reference    *string                `json:"reference,omitempty"`
	PaymentDate  *time.Time             `json:"paymentDate,omitempty"`
	RedirectURL  *string                `json:"redirectUrl,omitempty"`
	Status       *service.FundingStatus `json:"status,omitempty"`
	Extra        *interface{}           `json:"extra,omitempty"`
	Log          *interface{}           `json:"log,omitempty"`
	Tax          *string                `json:"tax,omitempty"`
}
type FundingFilter struct {
	CompanyID *string                    `json:"companyId,omitempty"`
	Status    []service.FundingStatus    `json:"status,omitempty"`
	Amount    *service.MinMaxFilterInput `json:"amount,omitempty"`
	Currency  *string                    `json:"currency,omitempty"`
	Filter    *service.FilterData        `json:"filter,omitempty"`
}
