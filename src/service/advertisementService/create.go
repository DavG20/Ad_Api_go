package advertisementService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/service/campaignService"
	"adtec/backend/src/service/contentService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.AdvertisementsModel, error) {
	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	campaign, _ := campaignService.Get(input.CampaignID)

	if campaign == nil {
		return nil, errors.New("invalid campaign Id ")
	}
	content, _ := contentService.Get(input.ContentID)

	if content == nil {
		return nil, errors.New("invalid content Id ")
	}
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.AdvertisementsSetParam{}

	if input.Status != nil {
		parameter = append(parameter, db.Advertisements.Status.Set(db.AdStatus(*input.Status)))
	}
	if input.MessageID != nil {
		parameter = append(parameter, db.Advertisements.MessageID.Set(*input.MessageID))
	}

	createdAdvertisement, err := client.Advertisements.CreateOne(
		db.Advertisements.ChannelID.Set(input.ChannelID),
		db.Advertisements.Content.Link(db.Contents.ID.Equals(input.ContentID)),
		db.Advertisements.Campaign.Link(db.Campaigns.ID.Equals(input.CampaignID)),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdAdvertisement, nil
}
