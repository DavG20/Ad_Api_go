package withdrawalService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Get(id string) (*db.WithdrawalsModel, error) {
	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	withdrawal, err := client.Withdrawals.FindUnique(
		db.Withdrawals.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("withdrawals not found")
	}
	return withdrawal, nil
}
