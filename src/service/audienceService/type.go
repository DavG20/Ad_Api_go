package audienceService

type CreateInput struct {
	CampaignID string  `json:"campaignId" validate:"required"`
	Category   *string `json:"category,omitempty"`
	Language   *string `json:"language,omitempty"`
}
type UpdateInput struct {
	CampaignID string  `json:"campaignId" validate:"required"`
	Category   *string `json:"category,omitempty"`
	Language   *string `json:"language,omitempty"`
}
