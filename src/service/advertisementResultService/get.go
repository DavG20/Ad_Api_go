package advertisementResultService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
)

func Get(id string) (*db.AdvertisementResultsModel, error) {
	client, ctx, err := database.GetCampaignClient()
	if err != nil {
		return nil, err
	}

	creadAdvertisementResult, err := client.AdvertisementResults.FindUnique(
		db.AdvertisementResults.AdvertisementID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return creadAdvertisementResult, nil
}
