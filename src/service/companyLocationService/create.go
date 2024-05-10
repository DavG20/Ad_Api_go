package companyLocationService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/service/companyService"
	"adtec/backend/src/service/locationService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.CompanyLocationsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	company, _ := companyService.Get(input.CompanyID)

	if company == nil {
		return nil, errors.New("company not found")
	}
	location, _ := locationService.Get(input.LocationID)

	if location == nil {
		return nil, errors.New("location not found")
	}

	companyLocation, _ := Get(input.CompanyID, input.LocationID)
	if companyLocation != nil {
		return nil, errors.New("companyLocation already exist")
	}

	createdCompanyLocation, err := client.CompanyLocations.CreateOne(
		db.CompanyLocations.Company.Link(db.Companies.ID.Equals(input.CompanyID)),
		db.CompanyLocations.Location.Link(db.Locations.ID.Equals(input.LocationID)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdCompanyLocation, nil
}
