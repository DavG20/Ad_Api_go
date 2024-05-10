package channelLifeTimeBalanceService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.ChannelLifeTimeBalancesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.ChannelLifeTimeBalancesSetParam{}

	if input.Currency != nil {
		parameter = append(parameter, db.ChannelLifeTimeBalances.Currency.Set(*input.Currency))
	}

	createdChannelLifeTimeBalance, err := client.ChannelLifeTimeBalances.CreateOne(
		db.ChannelLifeTimeBalances.AccountID.Set(input.AccountID),
		db.ChannelLifeTimeBalances.ChannelID.Set(input.ChannelID),
		db.ChannelLifeTimeBalances.TotalAmount.Set(input.TotalAmount),
		db.ChannelLifeTimeBalances.TotalHour.Set(input.TotalHour),
		db.ChannelLifeTimeBalances.TotalAd.Set(input.TotalAd),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdChannelLifeTimeBalance, nil
}
