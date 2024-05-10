package companyService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Get(id string) (*db.CompaniesModel, error) {
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	content, err := client.Companies.FindUnique(
		db.Companies.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("company not found")
	}
	return content, nil
}
