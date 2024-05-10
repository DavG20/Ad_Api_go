package channelBalanceRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelBalanceService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateChannelBalanceInput) (*model.ChannelBalances, error) {

	updatedData, err := channelBalanceService.Update(channelBalanceService.UpdateInput{
		ID:       input.ID,
		Amount:   input.Amount,
		Currency: input.Currency,
	})
	if err != nil {
		return nil, err
	}

	updatedChannelBalance := utils.ChannelBalanceGraphqlConverter(updatedData)
	return updatedChannelBalance, nil
}
