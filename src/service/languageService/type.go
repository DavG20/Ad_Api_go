package languageService

import "adtec/backend/src/service"

type CreateInput struct {
	Language string `json:"language" validate:"required,min=2,max=20"`
}
type FindManyInput struct {
	Filter *service.FilterData `json:"filter"`
}
