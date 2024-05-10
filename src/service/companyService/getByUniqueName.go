package companyService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func GetByUniqueName(uniqueName string) (*db.CompaniesModel, error) {
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	content, err := client.Companies.FindUnique(
		db.Companies.UniqueName.Equals(uniqueName),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("location not found")
	}
	return content, nil
}
