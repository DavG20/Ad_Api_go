package accountService

import (
	"adtec/backend/src/service"
	"time"
)

type CreateInput struct {
	UserID      string               `json:"userId" validate:"required"`
	UserName    string               `json:"userName" validate:"required,min=4"`
	Email       *string              `json:"email,omitempty" validate:"omitempty,email"`
	PhoneNumber *string              `json:"phoneNumber,omitempty" validate:"omitempty,startswith=+251"`
	FullName    *string              `json:"fullName,omitempty"`
	BirthDate   *time.Time           `json:"birthDate,omitempty" validate:"omitempty,ValidateTime"`
	AccountType *service.AccountType `json:"accountType,omitempty"`
}
type UpdateInput struct {
	ID          string               `json:"id" validate:"required"`
	UserName    *string              `json:"userName,omitempty" validate:"omitempty,min=4"`
	Email       *string              `json:"email,omitempty" validate:"omitempty,email"`
	PhoneNumber *string              `json:"phoneNumber,omitempty" validate:"omitempty,startswith=+251"`
	FullName    *string              `json:"fullName,omitempty"`
	BirthDate   *time.Time           `json:"birthDate,omitempty" validate:"omitempty,ValidateTime"`
	AccountType *service.AccountType `json:"accountType,omitempty"`
}
