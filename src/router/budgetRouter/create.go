package budgetRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/budgetService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateBudgetInput) (*model.Budgets, error) {

	budgetData, err := budgetService.Create(
		budgetService.CreateInput{
			CampaignID:     input.CampaignID,
			InitialAmount:  input.InitialAmount,
			UsedAmount:     input.UsedAmount,
			RefundedAmount: input.RefundedAmount,
			Currency:       input.Currency,
		},
	)
	if err != nil {
		return nil, err
	}

	createdBudget := utils.BudgetGraphqlConverter(budgetData)

	return createdBudget, nil

}
