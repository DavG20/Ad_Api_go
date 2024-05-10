package adPaymentService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.AdPaymentsModel, error) {
	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}
	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.AdPaymentsSetParam{}

	if input.Amount != nil {
		parameter = append(parameter, db.AdPayments.Amount.Set(*input.Amount))
	}

	if input.Currency != nil {
		parameter = append(parameter, db.AdPayments.Currency.Set(*input.Currency))
	}
	if input.CompletionTime != nil {
		parameter = append(parameter, db.AdPayments.CompletionTime.Set(*input.CompletionTime))
	}

	createdAdPayment, err := client.AdPayments.CreateOne(
		db.AdPayments.AdvertisementID.Set(input.AdvertisementID),
		db.AdPayments.AccountID.Set(input.AccountID),
		db.AdPayments.CampaignID.Set(input.CampaignID),
		db.AdPayments.ChannelID.Set(input.ChannelID),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdAdPayment, nil
}
