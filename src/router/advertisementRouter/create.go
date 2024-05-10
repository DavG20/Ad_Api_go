package advertisementRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/advertisementService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateAdvertisementInput) (*model.Advertisements, error) {

	advertisementData, err := advertisementService.Create(
		advertisementService.CreateInput{
			CampaignID: input.CampaignID,
			ContentID:  input.ContentID,
			ChannelID:  input.ChannelID,
			MessageID:  input.MessageID,
			Status:     (*service.AdStatus)(input.Status),
		},
	)
	if err != nil {
		return nil, err
	}

	createdAdvertisement := utils.AdvertisementGraphqlConverter(advertisementData)

	return createdAdvertisement, nil
}
