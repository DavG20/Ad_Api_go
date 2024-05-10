package withdrawalChannelLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/withdrawalChannelLinkService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.WithdrawalChannelLinkInput) (*model.WithdrawalChannelLinks, error) {
	createdWithdrawalChannelLink, err := withdrawalChannelLinkService.Create(
		withdrawalChannelLinkService.WithdrawalChannelLinkInput{
			ChannelBalanceID: input.ChannelBalanceID,
			WithdrawalID:     input.WithdrawalID},
	)
	if err != nil {
		return nil, err
	}

	withdrawalChannelLink := utils.WithdrawalChannelLinkGraphqlConverter(createdWithdrawalChannelLink)
	return withdrawalChannelLink, nil
}
