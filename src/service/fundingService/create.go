package fundingService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"encoding/json"
	"errors"
)

func Create(input CreateInput) (*db.FundingModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.FundingSetParam{}

	if input.Method != nil {
		parameter = append(parameter, db.Funding.Method.Set(*input.Method))
	}

	if input.FundingTxRef != nil {
		parameter = append(parameter, db.Funding.FundingTxRef.Set(*input.FundingTxRef))
	}
	if input.Reference != nil {
		parameter = append(parameter, db.Funding.Reference.Set(*input.Reference))
	}

	if input.PaymentDate != nil {
		parameter = append(parameter, db.Funding.PaymentDate.Set(*input.PaymentDate))
	}
	if input.RedirectURL != nil {
		parameter = append(parameter, db.Funding.RedirectURL.Set(*input.RedirectURL))
	}
	if input.Status != nil {
		parameter = append(parameter, db.Funding.Status.Set(db.FundingStatus(*input.Status)))
	}

	if input.Extra != nil {
		newExtra, _ := json.Marshal(input.Extra)
		parameter = append(parameter, db.Funding.Extra.Set(newExtra))
	}

	if input.Log != nil {
		newLog, _ := json.Marshal(input.Log)
		parameter = append(parameter, db.Funding.Log.Set(newLog))
	}
	createdFunding, err := client.Funding.CreateOne(
		db.Funding.CompanyID.Set(input.CompanyID),
		db.Funding.Amount.Set(input.Amount),
		db.Funding.Key.Set(input.Key),
		db.Funding.Tax.Set(input.Tax),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdFunding, nil
}
