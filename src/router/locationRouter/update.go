package locationRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/locationService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateLocationInput) (*model.Locations, error) {

	updatedData, err := locationService.Update(locationService.UpdateInput{
		ID:         input.ID,
		Country:    input.Country,
		State:      input.State,
		City:       input.City,
		Address:    input.Address,
		PostalCode: input.PostalCode,
	})
	if err != nil {
		return nil, err
	}

	updatedLocation := utils.LocationGraphqlConverter(updatedData)
	return updatedLocation, nil
}
