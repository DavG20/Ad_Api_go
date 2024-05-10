package locationService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.LocationsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.LocationsSetParam{}

	if input.Address != nil {
		parameter = append(parameter, db.Locations.Address.Set(*input.Address))
	}
	if input.PostalCode != nil {
		parameter = append(parameter, db.Locations.PostalCode.Set(*input.PostalCode))
	}

	createdLocation, err := client.Locations.CreateOne(
		db.Locations.Country.Set(input.Country),
		db.Locations.State.Set(input.State),
		db.Locations.City.Set(input.City),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdLocation, nil
}
