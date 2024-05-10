package withdrawalService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/service"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"encoding/json"
	"errors"
	"reflect"
)

func Update(input UpdateInput) (*db.WithdrawalsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	withdrawal, _ := Get(input.ID)

	if withdrawal == nil {
		return nil, errors.New("invalid withdrawal id or withdrawal doesn't exist")
	}

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.WithdrawalsSetParam{}

	if input.TotalAmount != nil {
		if *input.TotalAmount != withdrawal.TotalAmount {
			parameter = append(parameter, db.Withdrawals.TotalAmount.Set(*input.TotalAmount))
		}
	}
	if input.Currency != nil {
		if *input.Currency != withdrawal.Currency {
			parameter = append(parameter, db.Withdrawals.Currency.Set(*input.Currency))
		}
	}

	if input.FundingKey != nil {
		fundingKey, _ := withdrawal.FundingKey()
		if *input.FundingKey != fundingKey {
			parameter = append(parameter, db.Withdrawals.FundingKey.Set(*input.FundingKey))
		}
	}
	if input.Method != nil {
		method, _ := withdrawal.Method()
		if *input.Method != method {
			parameter = append(parameter, db.Withdrawals.Method.Set(*input.Method))
		}
	}
	if input.Reference != nil {
		reference, _ := withdrawal.Reference()
		if *input.Reference != reference {
			parameter = append(parameter, db.Withdrawals.Reference.Set(*input.Reference))
		}
	}
	if input.PaymentDate != nil {
		paymentDate, _ := withdrawal.PaymentDate()
		if *input.PaymentDate != paymentDate {
			parameter = append(parameter, db.Withdrawals.PaymentDate.Set(*input.PaymentDate))
		}
	}

	if input.Status != nil {

		if *input.Status != service.FundingStatus(withdrawal.Status) {
			parameter = append(parameter, db.Withdrawals.Status.Set(db.FundingStatus(*input.Status)))
		}

	}

	if input.Extra != nil {
		newExtra, _ := json.Marshal(input.Extra)
		if !reflect.DeepEqual(withdrawal.Extra, newExtra) {
			parameter = append(parameter, db.Withdrawals.Extra.Set(newExtra))
		}
	}
	if input.Log != nil {
		newLog, _ := json.Marshal(input.Log)
		if !reflect.DeepEqual(withdrawal.Log, newLog) {
			parameter = append(parameter, db.Withdrawals.Log.Set(newLog))
		}
	}

	updateWithdrawal, err := client.Withdrawals.FindUnique(
		db.Withdrawals.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateWithdrawal, nil
}
