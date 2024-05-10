package campaignService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.CampaignsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	campaignByName, _ := GetCampaignByName(input.Name)

	if campaignByName != nil {
		return nil, errors.New("name Taken")
	}
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.CampaignsSetParam{}

	if input.Objective != nil {
		parameter = append(parameter, db.Campaigns.Objective.Set(db.ObjectiveType(*input.Objective)))
	}
	if input.StartDate != nil {
		parameter = append(parameter, db.Campaigns.StartDate.Set(*input.StartDate))
	}

	createdCampaign, err := client.Campaigns.CreateOne(
		db.Campaigns.CompanyID.Set(input.CompanyID),
		db.Campaigns.Name.Set(input.Name),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdCampaign, nil
}
