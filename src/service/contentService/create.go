package contentService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/service/campaignService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.ContentsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	campaign, err := campaignService.Get(input.CampaignID)
	if err != nil || campaign == nil {
		return nil, errors.New("campaign  not found or invalid campaign id")
	}
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.ContentsSetParam{}

	if input.ContentType != nil {
		parameter = append(parameter, db.Contents.ContentType.Set(db.ContentType(*input.ContentType)))
	}
	if input.Description != nil {
		parameter = append(parameter, db.Contents.Description.Set(*input.Description))
	}

	createdContent, err := client.Contents.CreateOne(
		db.Contents.Campaign.Link(db.Campaigns.ID.Equals(input.CampaignID)),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdContent, nil
}
