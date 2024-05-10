package channelBalanceService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.ChannelBalancesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	channelBalance, _ := Get(input.ID)

	if channelBalance == nil {
		return nil, errors.New("invalid channelBalance id or ContentLink doesn't exist")
	}

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.ChannelBalancesSetParam{}

	if input.Amount != nil {
		if *input.Amount != channelBalance.Amount {
			parameter = append(parameter, db.ChannelBalances.Amount.Set(*input.Amount))

		}
	}
	if input.Currency != nil {
		if *input.Currency != channelBalance.Currency {
			parameter = append(parameter, db.ChannelBalances.Currency.Set(*input.Currency))

		}
	}

	UpdateChannelBalance, err := client.ChannelBalances.FindUnique(
		db.ChannelBalances.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return UpdateChannelBalance, nil
}
