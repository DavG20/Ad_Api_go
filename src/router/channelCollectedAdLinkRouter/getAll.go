package channelCollectedAdLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelCollectedAdLinkService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.ChannelCollectedAdLinkFilter) ([]*model.ChannelCollectedAdLinks, error) {

	channelCollectedAdLinkData, err := channelCollectedAdLinkService.GetAll(convertor.ChannelCollectedAdLinkFilterConverter(input))

	if err != nil {
		return nil, err
	}
	channelCollectedAdLinks := []*model.ChannelCollectedAdLinks{}

	for _, channelCollectedAdLink := range channelCollectedAdLinkData {
		channelCollectedAdLink := utils.ChannelCollectedAdLinkGraphqlConverter(channelCollectedAdLink)
		channelCollectedAdLinks = append(channelCollectedAdLinks, channelCollectedAdLink)
	}
	return channelCollectedAdLinks, nil

}
