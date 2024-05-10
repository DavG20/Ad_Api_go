package companyBankingService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
)

func Get(id string) (*db.CompanyBankingsModel, error) {

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	companyBanking, err := client.CompanyBankings.FindUnique(
		db.CompanyBankings.ID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return companyBanking, nil

}
