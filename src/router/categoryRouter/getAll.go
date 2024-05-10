package categoryRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/categoryService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.CategoryFilter) ([]*model.Categories, error) {

	categoriesData, err := categoryService.GetAll(convertor.CategoryFilterConverter(input))

	if err != nil {
		return nil, err
	}

	categories := []*model.Categories{}

	for _, category := range categoriesData {
		category := utils.CategoryGraphqlConverter(category)
		categories = append(categories, category)
	}

	return categories, nil

}
