package advertisementService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/service"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.AdvertisementsModel, error) {
	
	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	advertisement, _ := Get(input.ID)

	if advertisement == nil {
		return nil, errors.New("invalid advertisement id or advertisement doesn't exist")
	}

	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.AdvertisementsSetParam{}

	if input.Status != nil {

		if *input.Status != service.AdStatus(advertisement.Status) {
			parameter = append(parameter, db.Advertisements.Status.Set(db.AdStatus(*input.Status)))
		}

	}

	updateAdvertisement, err := client.Advertisements.FindUnique(
		db.Advertisements.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateAdvertisement, nil
}
