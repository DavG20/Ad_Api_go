package categoryRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/categoryService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Categories, error) {
	categoryData, err := categoryService.Get(id)
	if err != nil {
		return nil, err
	}

	category := utils.CategoryGraphqlConverter(categoryData)

	return category, err
}
