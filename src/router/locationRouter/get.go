package locationRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/locationService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Locations, error) {
	locationData, err := locationService.Get(id)
	if err != nil {
		return nil, err
	}

	location := utils.LocationGraphqlConverter(locationData)

	return location, nil
}
