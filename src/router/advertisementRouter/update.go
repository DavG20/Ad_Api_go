package advertisementRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/advertisementService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateAdvertisementInput) (*model.Advertisements, error) {

	updatedData, err := advertisementService.Update(advertisementService.UpdateInput{
		ID:     input.ID,
		Status: (*service.AdStatus)(input.Status),
	})
	if err != nil {
		return nil, err
	}

	updatedAdvertisement := utils.AdvertisementGraphqlConverter(updatedData)
	return updatedAdvertisement, nil
}
