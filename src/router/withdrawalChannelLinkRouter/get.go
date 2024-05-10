package withdrawalChannelLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/withdrawalChannelLinkService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, withdrawalId string) (*model.WithdrawalChannelLinks, error) {
	withdrawalChannelLinkData, err := withdrawalChannelLinkService.Get(withdrawalId)

	if err != nil {
		return nil, err
	}

	withdrawalChannelLink := utils.WithdrawalChannelLinkGraphqlConverter(withdrawalChannelLinkData)
	return withdrawalChannelLink, nil
}
