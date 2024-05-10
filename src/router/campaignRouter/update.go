package campaignRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/campaignService"
	"adtec/backend/src/utils"
	"context"
	"time"
)

func Update(ctx context.Context, input model.UpdateCampaignInput) (*model.Campaigns, error) {
	var startDate *time.Time

	if input.StartDate != nil {
		date, _ := time.Parse("2006-01-02", *input.StartDate)
		startDate = &date
	}

	updatedData, err := campaignService.Update(campaignService.UpdateInput{
		ID:        input.ID,
		Name:      input.Name,
		Objective: (*service.ObjectiveType)(input.Objective),
		StartDate: startDate,
	})
	if err != nil {
		return nil, err
	}

	updatedCampaign := utils.CampaignGraphqlConverter(updatedData)
	return updatedCampaign, nil
}
