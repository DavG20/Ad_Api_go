package channelLifeTimeBalanceRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelLifeTimeBalanceService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, channelId string) (*model.ChannelLifeTimeBalances, error) {
	channelLifeTimeBalanceData, err := channelLifeTimeBalanceService.Get(channelId)
	if err != nil {
		return nil, err
	}

	channelLifeTimeBalance := utils.ChannelLifeTimeBalanceGraphqlConverter(channelLifeTimeBalanceData)

	return channelLifeTimeBalance, nil
}
