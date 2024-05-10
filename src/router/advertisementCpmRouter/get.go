package advertisementCpmRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/advertisementCpmService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.AdvertisementCPMs, error) {
	advertisementCPMData, err := advertisementCpmService.Get(id)
	if err != nil {
		return nil, err
	}

	advertisementCPM := utils.AdvertisementCPMGraphqlConverter(advertisementCPMData)

	return advertisementCPM, nil
}
