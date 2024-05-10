package withdrawalChannelLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/withdrawalChannelLinkService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetByChannelIdAndAdPaymentId(ctx context.Context, input *model.WithdrawalChannelLinkInput) (*model.WithdrawalChannelLinks, error) {
	withdrawalChannelLinkData, err := withdrawalChannelLinkService.GetByChannelBalanceIdAndWithdrawalId(convertor.WithdrawalChannelLinkInputConverter(input))

	if err != nil {
		return nil, err
	}

	withdrawalChannelLink := utils.WithdrawalChannelLinkGraphqlConverter(withdrawalChannelLinkData)
	return withdrawalChannelLink, nil
}
