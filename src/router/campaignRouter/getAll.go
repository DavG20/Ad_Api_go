package campaignRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/campaignService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.CampaignFilter) ([]*model.Campaigns, error) {

	campaignData, err := campaignService.GetAll(convertor.CampaignFilterConverter(input))

	if err != nil {
		return nil, err
	}
	campaigns := []*model.Campaigns{}

	for _, campaign := range campaignData {
		campaign := utils.CampaignGraphqlConverter(campaign)
		campaigns = append(campaigns, campaign)
	}
	return campaigns, nil

}
