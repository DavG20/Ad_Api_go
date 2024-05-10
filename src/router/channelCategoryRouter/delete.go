package channelCategoryRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelCategoryService"
	"adtec/backend/src/utils/convertor"
	"context"
)

func Delete(ctx context.Context, input *model.ChannelCategoryInput) (*model.DeletionResult, error) {

	deletionResult, err := channelCategoryService.Delete(convertor.ChannelCategoryInputConverter(input))

	return deletionResult, err

}
