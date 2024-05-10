package channelRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.ChannelFilter) ([]*model.Channels, error) {

	channelData, err := channelService.GetAll(convertor.ChannelFilterConverter(input))

	if err != nil {
		return nil, err
	}
	channels := []*model.Channels{}

	for _, channel := range channelData {
		channel := utils.ChannelGraphqlConverter(channel)
		channels = append(channels, channel)
	}
	return channels, nil

}
