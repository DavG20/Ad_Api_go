package channelCollectedAdLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelCollectedAdLinkService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, adPaymentId string) (*model.ChannelCollectedAdLinks, error) {
	channelCollectedAdLinkData, err := channelCollectedAdLinkService.Get(adPaymentId)

	if err != nil {
		return nil, err
	}

	channelCollectedAdLink := utils.ChannelCollectedAdLinkGraphqlConverter(channelCollectedAdLinkData)
	return channelCollectedAdLink, nil
}
