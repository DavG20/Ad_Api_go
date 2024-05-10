package campaignService

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Delete(id string) (*model.DeletionResult, error) {
	client, ctx, err := database.GetCampaignClient()

	deletionResult := &model.DeletionResult{
		Success: true,
		Message: "campaign deleted Successfully",
	}

	if err != nil {
		deletionResult.Message = "internal error of db"
		deletionResult.Success = false
		return deletionResult, err
	}
	campaign, _ := Get(id)

	if campaign == nil {
		deletionResult.Success = false
		deletionResult.Message = "Channel not Found"
		return deletionResult, errors.New("channel not Found")
	}

	_, err = client.Campaigns.FindUnique(
		db.Campaigns.ID.Equals(id),
	).Delete().Exec(ctx)
	if err != nil {
		return nil, err
	}

	return deletionResult, nil

}
