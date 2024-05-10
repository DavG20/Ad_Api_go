package advertisementCpmRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/advertisementCpmService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateAdvertisementCPMInput) (*model.AdvertisementCPMs, error) {

	advertisementCPMData, err := advertisementCpmService.Create(
		advertisementCpmService.CreateInput{
			AdvertisementID: input.AdvertisementID,
			Rate:            input.Rate,
			ChannelShare:    input.ChannelShare,
			ProviderShare:   input.ProviderShare,
			TotalBudget:     input.TotalBudget,
			MaxLifeCycle:    input.MaxLifeCycle,
			RequiredViews:   input.RequiredViews,
		},
	)
	if err != nil {
		return nil, err
	}

	createdAdvertisementCPM := utils.AdvertisementCPMGraphqlConverter(advertisementCPMData)

	return createdAdvertisementCPM, nil
}
