package audienceService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.AudiencesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	audience, _ := Get(input.CampaignID)

	if audience == nil {
		return nil, errors.New("invalid audience id or audience doesn't exist")
	}

	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.AudiencesSetParam{}

	if input.Category != nil {
		category, _ := audience.Category()
		if *input.Category != category {
			parameter = append(parameter, db.Audiences.Category.Set(*input.Category))

		}
	}
	if input.Language != nil {
		language, _ := audience.Language()
		if *input.Language != language {
			parameter = append(parameter, db.Audiences.Language.Set(*input.Language))

		}
	}

	updateAudience, err := client.Audiences.FindUnique(
		db.Audiences.CampaignID.Equals(input.CampaignID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateAudience, nil
}
