package channelCategoryRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelCategoryService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.ChannelCategoryFilter) ([]*model.ChannelCategories, error) {

	channelCategoriesData, err := channelCategoryService.GetAll(convertor.ChannelCategoryFilterConverter(input))

	if err != nil {
		return nil, err
	}
	channelCategories := []*model.ChannelCategories{}

	for _, channelCategory := range channelCategoriesData {
		channelCategory := utils.ChannelCategoryGraphqlConverter(channelCategory)
		channelCategories = append(channelCategories, channelCategory)
	}
	return channelCategories, nil

}
