package channelDetailRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelDetailService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, channelId string) (*model.ChannelDetails, error) {
	channelDetailData, err := channelDetailService.Get(channelId)
	if err != nil {
		return nil, err
	}
	channelDetail := utils.ChannelDetailGraphqlConverter(channelDetailData)

	return channelDetail, nil

}
