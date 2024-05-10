package channelCollectedAdLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelCollectedAdLinkService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.ChannelCollectedAdLinkInput) (*model.ChannelCollectedAdLinks, error) {
	createdChannelCollectedAdLink, err := channelCollectedAdLinkService.Create(
		channelCollectedAdLinkService.ChannelCollectedAdLinkInput{
			ChannelBalanceID: input.ChannelBalanceID,
			AdPaymentID:      input.AdPaymentID,
		},
	)
	if err != nil {
		return nil, err
	}

	channelCollectedAdLink := utils.ChannelCollectedAdLinkGraphqlConverter(createdChannelCollectedAdLink)
	return channelCollectedAdLink, nil
}
