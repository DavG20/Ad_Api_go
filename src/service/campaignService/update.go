package campaignService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/service"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.CampaignsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	campaign, _ := Get(input.ID)

	if campaign == nil {
		return nil, errors.New("invalid campaign id or channel doesn't exist")
	}

	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.CampaignsSetParam{}

	if input.Name != nil {
		nameExist, _ := GetCampaignByName(*input.Name)

		if nameExist != nil {
			return nil, errors.New("name is taken")
		}
		if *input.Name != campaign.Name {
			parameter = append(parameter, db.Campaigns.Name.Set(*input.Name))
		}

	}

	if input.Objective != nil {
		objective, _ := campaign.Objective()
		if *input.Objective != service.ObjectiveType(objective) {
			parameter = append(parameter, db.Campaigns.Objective.Set(db.ObjectiveType(*input.Objective)))

		}
	}
	if input.StartDate != nil {
		startDate, _ := campaign.StartDate()
		if *input.StartDate != startDate {
			parameter = append(parameter, db.Campaigns.StartDate.Set(*input.StartDate))

		}
	}

	updateChannel, err := client.Campaigns.FindUnique(
		db.Campaigns.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateChannel, nil
}
