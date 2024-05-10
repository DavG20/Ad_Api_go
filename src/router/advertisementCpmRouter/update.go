package advertisementCpmRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/advertisementCpmService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateAdvertisementCPMInput) (*model.AdvertisementCPMs, error) {

	updatedData, err := advertisementCpmService.Update(advertisementCpmService.UpdateInput{
		AdvertisementID: input.AdvertisementID,
		Rate:            input.Rate,
		ChannelShare:    input.ChannelShare,
		ProviderShare:   input.ProviderShare,
		TotalBudget:     input.TotalBudget,
		MaxLifeCycle:    input.MaxLifeCycle,
		RequiredViews:   input.RequiredViews,
	})
	if err != nil {
		return nil, err
	}

	updateAdvertisementCPM := utils.AdvertisementCPMGraphqlConverter(updatedData)
	return updateAdvertisementCPM, nil
}
