package locationRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/locationService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateLocationInput) (*model.Locations, error) {

	contentData, err := locationService.Create(
		locationService.CreateInput{
			Country:    input.Country,
			State:      input.State,
			City:       input.City,
			Address:    input.Address,
			PostalCode: input.PostalCode,
		},
	)
	if err != nil {
		return nil, err
	}
	createdLocation := utils.LocationGraphqlConverter(contentData)

	return createdLocation, nil
}
