package contentService

import "adtec/backend/src/service"

type CreateInput struct {
	CampaignID  string               `json:"campaignId" validate:"required"`
	ContentType *service.ContentType `json:"contentType,omitempty" ` //Todo: validate contentType
	Description *string              `json:"description,omitempty" validate:"omitempty,min=10,max=50"`
}
type UpdateInput struct {
	ID          string               `json:"id" validate:"required"`
	ContentType *service.ContentType `json:"contentType,omitempty"`
	Description *string              `json:"description,omitempty" validate:"omitempty,min=10,max=50"`
}
