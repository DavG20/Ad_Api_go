package companyLocationService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
)

func Get(companyId, locationId string) (*db.CompanyLocationsModel, error) {
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	companyLocation, err := client.CompanyLocations.FindUnique(
		db.CompanyLocations.UniqueCompanyLocationID(
			db.CompanyLocations.CompanyID.Equals(companyId),
			db.CompanyLocations.LocationID.Equals(locationId),
		),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return companyLocation, nil
}
