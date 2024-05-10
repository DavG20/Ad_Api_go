package categoryRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/categoryService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateCategoryInput) (*model.Categories, error) {
	createdCategory, err := categoryService.Create(
		categoryService.CreateInput{
			Category: input.Category,
		},
	)
	if err != nil {
		return nil, err
	}

	category := utils.CategoryGraphqlConverter(createdCategory)
	return category, nil
}
