package withdrawalChannelLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/withdrawalChannelLinkService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.WithdrawalChannelLinkFilter) ([]*model.WithdrawalChannelLinks, error) {

	withdrawalChannelLinkData, err := withdrawalChannelLinkService.GetAll(convertor.WithdrawalChannelLinkFilterConverter(input))

	if err != nil {
		return nil, err
	}
	withdrawalChannelLinks := []*model.WithdrawalChannelLinks{}

	for _, withdrawalChannelLink := range withdrawalChannelLinkData {
		withdrawalChannelLink := utils.WithdrawalChannelLinkGraphqlConverter(withdrawalChannelLink)
		withdrawalChannelLinks = append(withdrawalChannelLinks, withdrawalChannelLink)
	}
	return withdrawalChannelLinks, nil

}
