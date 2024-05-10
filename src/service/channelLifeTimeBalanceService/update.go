package channelLifeTimeBalanceService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.ChannelLifeTimeBalancesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	channelLifeTimeBalance, _ := Get(input.ID)

	if channelLifeTimeBalance == nil {
		return nil, errors.New("invalid channelLifeTimeBalance id or ContentLink doesn't exist")
	}

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.ChannelLifeTimeBalancesSetParam{}

	if input.TotalAmount != nil {
		if *input.TotalAmount != channelLifeTimeBalance.TotalAmount {
			parameter = append(parameter, db.ChannelLifeTimeBalances.TotalAmount.Set(*input.TotalAmount))

		}
	}
	if input.Currency != nil {
		if *input.Currency != channelLifeTimeBalance.Currency {
			parameter = append(parameter, db.ChannelLifeTimeBalances.Currency.Set(*input.Currency))

		}
	}
	if input.TotalHour != nil {
		if *input.TotalHour != channelLifeTimeBalance.TotalHour {
			parameter = append(parameter, db.ChannelLifeTimeBalances.TotalHour.Set(*input.TotalHour))

		}
	}
	if input.TotalAd != nil {
		if *input.TotalAd != channelLifeTimeBalance.TotalAd {
			parameter = append(parameter, db.ChannelLifeTimeBalances.TotalAd.Set(*input.TotalAd))

		}
	}

	UpdateChannelLifeTimeBalance, err := client.ChannelLifeTimeBalances.FindUnique(
		db.ChannelLifeTimeBalances.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return UpdateChannelLifeTimeBalance, nil
}
