package advertisementCpmService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.AdvertisementCpmsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}
	
	advertisementCPM, _ := Get(input.AdvertisementID)

	if advertisementCPM == nil {
		return nil, errors.New("invalid AdvertisementCpm id or AdvertisementCpm doesn't exist")
	}

	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.AdvertisementCpmsSetParam{}

	if input.Rate != nil {
		if *input.Rate != advertisementCPM.Rate {
			parameter = append(parameter, db.AdvertisementCpms.Rate.Set(*input.Rate))

		}
	}
	if input.ProviderShare != nil {
		if *input.ProviderShare != advertisementCPM.ProviderShare {
			parameter = append(parameter, db.AdvertisementCpms.ProviderShare.Set(*input.ProviderShare))

		}
	}
	if input.ChannelShare != nil {
		if *input.ChannelShare != advertisementCPM.ChannelShare {
			parameter = append(parameter, db.AdvertisementCpms.ChannelShare.Set(*input.ChannelShare))

		}
	}
	if input.TotalBudget != nil {
		if *input.TotalBudget != advertisementCPM.TotalBudget {
			parameter = append(parameter, db.AdvertisementCpms.TotalBudget.Set(*input.TotalBudget))

		}
	}
	if input.MaxLifeCycle != nil {
		if *input.MaxLifeCycle != advertisementCPM.MaxLifeCycle {
			parameter = append(parameter, db.AdvertisementCpms.MaxLifeCycle.Set(*input.MaxLifeCycle))

		}
	}
	if input.RequiredViews != nil {
		if *input.RequiredViews != advertisementCPM.RequiredViews {
			parameter = append(parameter, db.AdvertisementCpms.RequiredViews.Set(*input.RequiredViews))

		}
	}

	updateAdvertisementCPM, err := client.AdvertisementCpms.FindUnique(
		db.AdvertisementCpms.AdvertisementID.Equals(input.AdvertisementID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateAdvertisementCPM, nil
}
