package accountBankingService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
)

func Get(id string) (*db.AccountBankingsModel, error) {

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	accountBanking, err := client.AccountBankings.FindUnique(
		db.AccountBankings.ID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return accountBanking, nil

}
