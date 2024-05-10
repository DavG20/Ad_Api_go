package fundingService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/service"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"encoding/json"
	"errors"
	"reflect"
)

func Update(input UpdateInput) (*db.FundingModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	funding, _ := Get(input.ID)

	if funding == nil {
		return nil, errors.New("invalid funding id or funding doesn't exist")
	}

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.FundingSetParam{}

	if input.Amount != nil {
		if *input.Amount != funding.Amount {
			parameter = append(parameter, db.Funding.Amount.Set(*input.Amount))
		}
	}
	if input.Currency != nil {
		if *input.Currency != funding.Currency {
			parameter = append(parameter, db.Funding.Currency.Set(*input.Currency))
		}
	}

	if input.Key != nil {
		if *input.Key != funding.Key {
			parameter = append(parameter, db.Funding.Key.Set(*input.Key))
		}
	}
	if input.Method != nil {
		method, _ := funding.Method()
		if *input.Method != method {
			parameter = append(parameter, db.Funding.Method.Set(*input.Method))
		}
	}
	if input.FundingTxRef != nil {
		fundingTxRef, _ := funding.FundingTxRef()
		if *input.FundingTxRef != fundingTxRef {
			parameter = append(parameter, db.Funding.FundingTxRef.Set(*input.FundingTxRef))
		}
	}
	if input.Reference != nil {
		reference, _ := funding.Reference()
		if *input.Reference != reference {
			parameter = append(parameter, db.Funding.Reference.Set(*input.Reference))
		}
	}
	if input.PaymentDate != nil {
		paymentDate, _ := funding.PaymentDate()
		if *input.PaymentDate != paymentDate {
			parameter = append(parameter, db.Funding.PaymentDate.Set(*input.PaymentDate))
		}
	}
	if input.RedirectURL != nil {
		redirectURL, _ := funding.RedirectURL()
		if *input.RedirectURL != redirectURL {
			parameter = append(parameter, db.Funding.RedirectURL.Set(*input.RedirectURL))
		}
	}
	if input.Extra != nil {
		newExtra, _ := json.Marshal(input.Extra)
		if !reflect.DeepEqual(funding.Extra, newExtra) {
			parameter = append(parameter, db.Funding.Extra.Set(newExtra))
		}
	}
	if input.Log != nil {
		newLog, _ := json.Marshal(input.Log)
		if !reflect.DeepEqual(funding.Log, newLog) {
			parameter = append(parameter, db.Funding.Log.Set(newLog))
		}
	}

	if input.Status != nil {

		if *input.Status != service.FundingStatus(funding.Status) {
			parameter = append(parameter, db.Funding.Status.Set(db.FundingStatus(*input.Status)))
		}

	}
	if input.Tax != nil {
		if *input.Tax != funding.Tax {
			parameter = append(parameter, db.Funding.Tax.Set(*input.Tax))
		}
	}

	updateFunding, err := client.Funding.FindUnique(
		db.Funding.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateFunding, nil
}
