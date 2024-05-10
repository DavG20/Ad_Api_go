package campaignRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/campaignService"
	"adtec/backend/src/utils"
	"context"
	"time"
)

func Create(ctx context.Context, input model.CreateCampaignInput) (*model.Campaigns, error) {
	var campaignStartAt *time.Time
	if input.StartDate != nil {
		date, _ := time.Parse("2006-01-02", *input.StartDate)
		campaignStartAt = &date
	}
	campaignData, err := campaignService.Create(
		campaignService.CreateInput{
			CompanyID: input.CompanyID,
			Name:      input.Name,
			Objective: (*service.ObjectiveType)(input.Objective),
			StartDate: campaignStartAt,
		},
	)
	if err != nil {
		return nil, err
	}

	createdCampaign := utils.CampaignGraphqlConverter(campaignData)

	return createdCampaign, nil
}
