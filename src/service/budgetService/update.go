package budgetService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.BudgetsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	budget, _ := Get(input.CampaignID)

	if budget == nil {
		return nil, errors.New("invalid budget id or budget doesn't exist")
	}

	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.BudgetsSetParam{}

	if input.InitialAmount != nil {
		if *input.InitialAmount != budget.InitialAmount {
			parameter = append(parameter, db.Budgets.InitialAmount.Set(*input.InitialAmount))

		}
	}
	if input.UsedAmount != nil {
		if *input.UsedAmount != budget.UsedAmount {
			parameter = append(parameter, db.Budgets.UsedAmount.Set(*input.UsedAmount))

		}
	}
	if input.RefundedAmount != nil {
		refundedAmount, _ := budget.RefundedAmount()
		if *input.RefundedAmount != refundedAmount {
			parameter = append(parameter, db.Budgets.RefundedAmount.Set(*input.RefundedAmount))

		}
	}
	if input.Currency != nil {
		if *input.Currency != budget.Currency {
			parameter = append(parameter, db.Budgets.Currency.Set(*input.Currency))

		}
	}

	updateBudget, err := client.Budgets.FindUnique(
		db.Budgets.CampaignID.Equals(input.CampaignID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateBudget, nil
}
