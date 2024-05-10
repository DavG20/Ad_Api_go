package withdrawalService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"encoding/json"
	"errors"
)

func Create(input CreateInput) (*db.WithdrawalsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}
	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.WithdrawalsSetParam{}

	if input.Currency != nil {
		parameter = append(parameter, db.Withdrawals.Currency.Set(*input.Currency))
	}
	if input.FundingKey != nil {
		parameter = append(parameter, db.Withdrawals.FundingKey.Set(*input.FundingKey))
	}

	if input.Method != nil {
		parameter = append(parameter, db.Withdrawals.Method.Set(*input.Method))
	}

	if input.Reference != nil {
		parameter = append(parameter, db.Withdrawals.Reference.Set(*input.Reference))
	}

	if input.PaymentDate != nil {
		parameter = append(parameter, db.Withdrawals.PaymentDate.Set(*input.PaymentDate))
	}
	if input.Status != nil {
		parameter = append(parameter, db.Withdrawals.Status.Set(db.FundingStatus(*input.Status)))
	}

	if input.Extra != nil {
		newExtra, _ := json.Marshal(input.Extra)
		parameter = append(parameter, db.Withdrawals.Extra.Set(newExtra))
	}

	if input.Log != nil {
		newLog, _ := json.Marshal(input.Log)
		parameter = append(parameter, db.Withdrawals.Log.Set(newLog))
	}
	createdWithdrawals, err := client.Withdrawals.CreateOne(
		db.Withdrawals.AccountID.Set(input.AccountID),
		db.Withdrawals.AccountBankID.Set(input.AccountBankID),
		db.Withdrawals.TotalAmount.Set(input.TotalAmount),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdWithdrawals, nil
}
