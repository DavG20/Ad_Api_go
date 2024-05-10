package advertisementResultRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/advertisementResultService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.AdvertisementResults, error) {
	advertisementResultData, err := advertisementResultService.Get(id)

	if err != nil || advertisementResultData == nil {
		return nil, err
	}
	advertisementResult := utils.AdvertisementResultGraphqlConverter(advertisementResultData)

	return advertisementResult, nil
}
