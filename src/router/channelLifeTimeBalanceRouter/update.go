package channelLifeTimeBalanceRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelLifeTimeBalanceService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateChannelLifeTimeBalanceInput) (*model.ChannelLifeTimeBalances, error) {

	updatedData, err := channelLifeTimeBalanceService.Update(channelLifeTimeBalanceService.UpdateInput{
		ID:          input.ID,
		TotalAmount: input.TotalAmount,
		Currency:    input.Currency,
		TotalHour:   input.TotalHour,
		TotalAd:     input.TotalAd,
	})
	if err != nil {
		return nil, err
	}

	updatedChannelLifeTimeBalance := utils.ChannelLifeTimeBalanceGraphqlConverter(updatedData)
	return updatedChannelLifeTimeBalance, nil
}
