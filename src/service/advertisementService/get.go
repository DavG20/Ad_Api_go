package advertisementService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Get(id string) (*db.AdvertisementsModel, error) {
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	advertisement, err := client.Advertisements.FindUnique(
		db.Advertisements.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("advertisements not found")
	}
	return advertisement, nil
}
