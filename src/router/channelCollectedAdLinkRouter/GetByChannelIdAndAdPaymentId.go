package channelCollectedAdLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelCollectedAdLinkService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetByChannelIdAndAdPaymentId(ctx context.Context, input *model.ChannelCollectedAdLinkInput) (*model.ChannelCollectedAdLinks, error) {
	channelCollectedAdLinkData, err := channelCollectedAdLinkService.GetByChannelIdAndAdPaymentId(convertor.ChannelCollectedAdLinkInputConverter(input))

	if err != nil {
		return nil, err
	}

	channelCollectedAdLink := utils.ChannelCollectedAdLinkGraphqlConverter(channelCollectedAdLinkData)
	return channelCollectedAdLink, nil
}
