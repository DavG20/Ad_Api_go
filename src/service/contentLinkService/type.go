package contentLinkService

type CreateInput struct {
	ContentID string  `json:"contentId" validate:"required"`
	Link      *string `json:"link,omitempty" validate:"omitempty,url"`
	Title     *string `json:"title,omitempty" validate:"omitempty,min=4,max=50"`
}
type UpdateInput struct {
	ContentID string  `json:"contentId" validate:"required"`
	Link      *string `json:"link,omitempty" validate:"omitempty,url"`
	Title     *string `json:"title,omitempty" validate:"omitempty,min=4,max=50"`
}
