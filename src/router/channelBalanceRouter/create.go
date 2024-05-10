package channelBalanceRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/service/channelBalanceService"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/utils"
	"context"
	"errors"
)

func Create(ctx context.Context, input model.CreateChannelBalanceInput) (*model.ChannelBalances, error) {
	account, err := accountService.Get(input.AccountID)
	if err != nil || account == nil {
		return nil, errors.New("account not found or invalid account id")
	}

	channel, err := channelService.Get(input.ChannelID)
	if channel == nil || err != nil {
		return nil, errors.New("channel not found ")
	}
	createData, err := channelBalanceService.Create(
		channelBalanceService.CreateInput{
			AccountID: input.AccountID,
			ChannelID: input.ChannelID,
			Amount:    input.Amount,
			Currency:  input.Currency,
		},
	)
	if err != nil {
		return nil, err
	}

	createdChannelBalance := utils.ChannelBalanceGraphqlConverter(createData)

	return createdChannelBalance, nil
}
