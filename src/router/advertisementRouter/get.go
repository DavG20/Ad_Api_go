package advertisementRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/advertisementService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Advertisements, error) {
	advertisementData, err := advertisementService.Get(id)
	if err != nil {
		return nil, err
	}

	advertisement := utils.AdvertisementGraphqlConverter(advertisementData)

	return advertisement, nil
}
