package locationService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.LocationsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	location, err := Get(input.ID)

	if location == nil {
		return nil, err
	}

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.LocationsSetParam{}

	if input.Country != nil {
		if *input.Country != location.Country {
			parameter = append(parameter, db.Locations.Country.Set(*input.Country))

		}
	}
	if input.City != nil {
		if *input.City != location.City {
			parameter = append(parameter, db.Locations.City.Set(*input.City))

		}
	}
	if input.State != nil {
		if *input.State != location.State {
			parameter = append(parameter, db.Locations.State.Set(*input.State))

		}
	}
	if input.Address != nil {
		address, _ := location.Address()
		if *input.Address != address {
			parameter = append(parameter, db.Locations.Address.Set(*input.Address))

		}
	}
	if input.PostalCode != nil {
		postalCode, _ := location.PostalCode()
		if *input.PostalCode != postalCode {
			parameter = append(parameter, db.Locations.PostalCode.Set(*input.PostalCode))

		}
	}

	updateLocation, err := client.Locations.FindUnique(
		db.Locations.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateLocation, nil
}
