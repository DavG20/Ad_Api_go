package categoryService

import "adtec/backend/src/service"

type CreateInput struct {
	Category string `json:"category" validate:"required"`
}
type FindManyInput struct {
	Filter *service.FilterData `json:"filter"`
}
