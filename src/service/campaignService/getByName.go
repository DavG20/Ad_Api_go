package campaignService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
)

func GetCampaignByName(name string) (*db.CampaignsModel, error) {
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	campaign, err := client.Campaigns.FindUnique(
		db.Campaigns.Name.Equals(name),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return campaign, nil
}
