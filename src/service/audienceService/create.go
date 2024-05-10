package audienceService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.AudiencesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	audience, _ := Get(input.CampaignID)

	if audience != nil {
		return nil, errors.New("campaign id is already registered")
	}

	client, ctx, err := database.GetCampaignClient()
	if err != nil {
		return nil, err
	}
	parameter := []db.AudiencesSetParam{}
	if input.Category != nil {
		parameter = append(parameter, db.Audiences.Category.Set(*input.Category))
	}
	if input.Language != nil {
		parameter = append(parameter, db.Audiences.Language.Set(*input.Language))
	}
	createdAudience, err := client.Audiences.CreateOne(
		db.Audiences.Campaign.Link(db.Campaigns.ID.Equals(input.CampaignID)),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdAudience, nil
}
