package channelBalanceService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.ChannelBalancesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.ChannelBalancesSetParam{}

	if input.Currency != nil {
		parameter = append(parameter, db.ChannelBalances.Currency.Set(*input.Currency))
	}
	if input.Amount != nil {
		parameter = append(parameter, db.ChannelBalances.Amount.Set(*input.Amount))
	}

	createdChannelBalance, err := client.ChannelBalances.CreateOne(
		db.ChannelBalances.AccountID.Set(input.AccountID),
		db.ChannelBalances.ChannelID.Set(input.ChannelID),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdChannelBalance, nil
}
