package channelCategoryRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelCategoryService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.ChannelCategoryInput) (*model.ChannelCategories, error) {
	createdChannelCategory, err := channelCategoryService.Create(
		channelCategoryService.ChannelCategoryInput{
			ChannelID: input.ChannelID,
			Category:  input.Category,
		},
	)
	if err != nil {
		return nil, err
	}

	channelCategory := utils.ChannelCategoryGraphqlConverter(createdChannelCategory)
	return channelCategory, nil
}
