package locationService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Get(id string) (*db.LocationsModel, error) {
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	location, err := client.Locations.FindUnique(
		db.Locations.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("location not found")
	}
	return location, nil
}
