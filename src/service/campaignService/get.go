package campaignService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Get(id string) (*db.CampaignsModel, error) {
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	campaign, err := client.Campaigns.FindUnique(
		db.Campaigns.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("campaign not found")
	}
	return campaign, nil
}
