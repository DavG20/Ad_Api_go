package channelRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Channels, error) {
	channelData, err := channelService.Get(id)
	if err != nil {
		return nil, err
	}

	channel := utils.ChannelGraphqlConverter(channelData)

	return channel, nil
}
