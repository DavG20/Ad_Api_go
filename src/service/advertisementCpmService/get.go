package advertisementCpmService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Get(id string) (*db.AdvertisementCpmsModel, error) {
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	advertisementCPM, err := client.AdvertisementCpms.FindUnique(
		db.AdvertisementCpms.AdvertisementID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("content not found")
	}
	return advertisementCPM, nil
}
