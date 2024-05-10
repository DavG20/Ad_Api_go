package fundingService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Get(id string) (*db.FundingModel, error) {
	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	funding, err := client.Funding.FindUnique(
		db.Funding.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("funding not found")
	}
	return funding, nil
}
