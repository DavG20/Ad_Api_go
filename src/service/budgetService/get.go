package budgetService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
)

func Get(id string) (*db.BudgetsModel, error) {
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	budget, err := client.Budgets.FindUnique(
		db.Budgets.CampaignID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return budget, nil
}
