package channelLifeTimeBalanceRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/service/channelLifeTimeBalanceService"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/utils"
	"context"
	"errors"
)

func Create(ctx context.Context, input model.CreateChannelLifeTimeBalanceInput) (*model.ChannelLifeTimeBalances, error) {
	account, err := accountService.Get(input.AccountID)
	if err != nil || account == nil {
		return nil, errors.New("account not found or invalid account id")
	}

	channel, err := channelService.Get(input.ChannelID)
	if channel == nil || err != nil {
		return nil, errors.New("channel not found ")
	}
	createData, err := channelLifeTimeBalanceService.Create(
		channelLifeTimeBalanceService.CreateInput{
			AccountID:   input.AccountID,
			ChannelID:   input.ChannelID,
			TotalAmount: input.TotalAmount,
			Currency:    input.Currency,
			TotalHour:   input.TotalHour,
			TotalAd:     input.TotalAd,
		},
	)
	if err != nil {
		return nil, err
	}

	createdChannelLifeTimeBalance := utils.ChannelLifeTimeBalanceGraphqlConverter(createData)

	return createdChannelLifeTimeBalance, nil
}
