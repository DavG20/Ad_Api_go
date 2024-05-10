package channelBalanceRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelBalanceService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, channelId string) (*model.ChannelBalances, error) {
	channelBalanceData, err := channelBalanceService.Get(channelId)
	if err != nil {
		return nil, err
	}

	channelBalance := utils.ChannelBalanceGraphqlConverter(channelBalanceData)

	return channelBalance, nil
}
