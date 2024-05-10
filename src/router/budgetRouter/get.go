package budgetRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/budgetService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Budgets, error) {
	budgetData, err := budgetService.Get(id)

	if err != nil {
		return nil, err
	}

	budget := utils.BudgetGraphqlConverter(budgetData)

	return budget, nil
}
