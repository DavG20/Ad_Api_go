package channelCollectedAdLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelCollectedAdLinkService"
	"adtec/backend/src/utils/convertor"
	"context"
)

func Delete(ctx context.Context, input *model.ChannelCollectedAdLinkInput) (*model.DeletionResult, error) {

	deletionResult, err := channelCollectedAdLinkService.Delete(convertor.ChannelCollectedAdLinkInputConverter(input))

	return deletionResult, err

}
