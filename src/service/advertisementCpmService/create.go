package advertisementCpmService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/service/advertisementService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.AdvertisementCpmsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	advertisement, err := advertisementService.Get(input.AdvertisementID)
	if err != nil || advertisement == nil {
		return nil, errors.New("advertisement  not found or invalid advertisement id")
	}
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	createdAdvertisementCPM, err := client.AdvertisementCpms.CreateOne(
		db.AdvertisementCpms.Rate.Set(input.Rate),
		db.AdvertisementCpms.ChannelShare.Set(input.ChannelShare),
		db.AdvertisementCpms.ProviderShare.Set(input.ProviderShare),
		db.AdvertisementCpms.TotalBudget.Set(input.TotalBudget),
		db.AdvertisementCpms.MaxLifeCycle.Set(input.MaxLifeCycle),
		db.AdvertisementCpms.RequiredViews.Set(input.RequiredViews),
		db.AdvertisementCpms.Advertisement.Link(db.Advertisements.ID.Equals(input.AdvertisementID)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdAdvertisementCPM, nil
}
