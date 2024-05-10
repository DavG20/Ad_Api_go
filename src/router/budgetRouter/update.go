package budgetRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/budgetService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateBudgetInput) (*model.Budgets, error) {

	updatedData, err := budgetService.Update(budgetService.UpdateInput{
		CampaignID:     input.CampaignID,
		InitialAmount:  input.InitialAmount,
		UsedAmount:     input.UsedAmount,
		RefundedAmount: input.RefundedAmount,
		Currency:       input.Currency,
	})
	if err != nil {
		return nil, err
	}

	updatedBudget := utils.BudgetGraphqlConverter(updatedData)
	return updatedBudget, nil
}
