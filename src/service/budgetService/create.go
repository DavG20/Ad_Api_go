package budgetService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.BudgetsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	budget, _ := Get(input.CampaignID)

	if budget != nil {
		return nil, errors.New("campaign id is already registered")
	}

	client, ctx, err := database.GetCampaignClient()
	if err != nil {
		return nil, err
	}
	parameter := []db.BudgetsSetParam{}

	if input.UsedAmount != nil {
		parameter = append(parameter, db.Budgets.UsedAmount.Set(*input.UsedAmount))
	}
	if input.RefundedAmount != nil {
		parameter = append(parameter, db.Budgets.RefundedAmount.Set(*input.RefundedAmount))
	}
	if input.Currency != nil {
		parameter = append(parameter, db.Budgets.Currency.Set(*input.Currency))
	}
	createdAudience, err := client.Budgets.CreateOne(
		db.Budgets.InitialAmount.Set(input.InitialAmount),
		db.Budgets.Campaign.Link(db.Campaigns.ID.Equals(input.CampaignID)),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdAudience, nil
}
