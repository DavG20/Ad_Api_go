package companyMemberService

import "adtec/backend/src/service"

type CreateInput struct {
	AccountID string               `json:"accountId" validate:"required"`
	CompanyID string               `json:"companyId" validate:"required"`
	Role      *service.CompanyRole `json:"role,omitempty"` // Todo: validate role
}
type UpdateInput struct {
	ID   string               `json:"id"`
	Role *service.CompanyRole `json:"role,omitempty"`
}
type CompanyMemberFilter struct {
	CompanyID *string               `json:"companyId,omitempty"`
	AccountID *string               `json:"accountId,omitempty"`
	Role      []service.CompanyRole `json:"role,omitempty"`
	Filter    *service.FilterData   `json:"filter,omitempty"`
}
