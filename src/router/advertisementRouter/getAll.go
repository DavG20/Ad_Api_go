package advertisementRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/advertisementService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.AdvertisementFilter) ([]*model.Advertisements, error) {

	advertisementData, err := advertisementService.GetAll(convertor.AdvertisementFilterConverter(input))

	if err != nil {
		return nil, err
	}
	advertisements := []*model.Advertisements{}

	for _, advertisement := range advertisementData {
		advertisement := utils.AdvertisementGraphqlConverter(advertisement)
		advertisements = append(advertisements, advertisement)
	}
	return advertisements, nil

}
