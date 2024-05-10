package campaignService

import (
	"adtec/backend/src/service"
	"time"
)

type CreateInput struct {
	CompanyID string                 `json:"companyId" validate:"required"`
	Name      string                 `json:"name" validate:"required,min=3"`
	Objective *service.ObjectiveType `json:"objective,omitempty" validate:"oneof=Reach Impression Clicks"`
	StartDate *time.Time             `json:"startDate,omitempty"`
}
type UpdateInput struct {
	ID        string                 `json:"id" validate:"required"`
	Name      *string                `json:"name,omitempty" validate:"omitempty,min=3"`
	Objective *service.ObjectiveType `json:"objective,omitempty" validate:"oneof=Reach Impression Clicks"`
	StartDate *time.Time             `json:"startDate,omitempty"`
}
type CampaignFilter struct {
	Name                       *string                    `json:"name,omitempty"`
	CompanyID                  *string                    `json:"companyId,omitempty"`
	Objective                  []service.ObjectiveType    `json:"objective,omitempty"`
	Categories                 []string                   `json:"categories,omitempty"`
	Languages                  []string                   `json:"languages,omitempty"`
	InitialBudget              *service.MinMaxFilterInput `json:"initialBudget,omitempty"`
	UsedAmount                 *service.MinMaxFilterInput `json:"usedAmount,omitempty"`
	ContentType                []service.ContentType      `json:"contentType,omitempty"`
	AdvertisementStatus        []service.AdStatus         `json:"advertisementStatus,omitempty"`
	AdvertisementRate          *service.MinMaxFilterInput `json:"advertisementRate,omitempty"`
	AdvertisementRequiredViews *service.MinMaxFilterInput `json:"advertisementRequiredViews,omitempty"`
	AdvertisementViews         *service.MinMaxFilterInput `json:"advertisementViews,omitempty"`
	AdvertisementHours         *service.MinMaxFilterInput `json:"advertisementHours,omitempty"`
	Filter                     *service.FilterData        `json:"filter,omitempty"`
}
