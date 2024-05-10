package campaignRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/campaignService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Campaigns, error) {
	campaignData, err := campaignService.Get(id)
	if err != nil {
		return nil, err
	}

	campaign := utils.CampaignGraphqlConverter(campaignData)

	return campaign, nil
}
