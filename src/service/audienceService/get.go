package audienceService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
)

func Get(id string) (*db.AudiencesModel, error) {
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	campaign, err := client.Audiences.FindUnique(
		db.Audiences.CampaignID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return campaign, nil
}
