package bankService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
)

func Get(id string) (*db.BanksModel, error) {
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	bank, err := client.Banks.FindUnique(
		db.Banks.ID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return bank, nil
}
