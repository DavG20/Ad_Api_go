package channelDetailRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelDetailService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, channelDetailFilter *model.ChannelDetailFilter) ([]*model.ChannelDetails, error) {

	channelDetailData, err := channelDetailService.GetAll(convertor.ChannelDetailFilterConverter(channelDetailFilter))

	if err != nil {
		return nil, err
	}
	channelDetails := []*model.ChannelDetails{}

	for _, channelDetail := range channelDetailData {
		channelDetail := utils.ChannelDetailGraphqlConverter(channelDetail)
		channelDetails = append(channelDetails, channelDetail)

	}
	return channelDetails, nil
}
