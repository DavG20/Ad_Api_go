package advertisementResultService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/service/advertisementService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.AdvertisementResultsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	advertisement, err := advertisementService.Get(input.AdvertisementID)
	if err != nil || advertisement == nil {
		return nil, errors.New("advertisement  not found or invalid Advertisement id")
	}

	advertisementResult, _ := Get(input.AdvertisementID)
	if advertisementResult != nil {
		return nil, errors.New("advertisement id taken")
	}
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}
	parameter := []db.AdvertisementResultsSetParam{}

	if input.Budget != nil {
		parameter = append(parameter, db.AdvertisementResults.Budget.Set(*input.Budget))
	}

	if input.ProviderBudgetShare != nil {
		parameter = append(parameter, db.AdvertisementResults.ProviderBudgetShare.Set(*input.ProviderBudgetShare))
	}

	if input.ChannelBudgetShare != nil {
		parameter = append(parameter, db.AdvertisementResults.ChannelBudgetShare.Set(*input.ChannelBudgetShare))
	}

	if input.TotalClick != nil {
		parameter = append(parameter, db.AdvertisementResults.TotalClick.Set(*input.TotalClick))
	}

	if input.TotalForward != nil {
		parameter = append(parameter, db.AdvertisementResults.TotalForward.Set(*input.TotalForward))
	}

	if input.TotalHour != nil {
		parameter = append(parameter, db.AdvertisementResults.TotalHour.Set(*input.TotalHour))
	}

	if input.TotalReaction != nil {
		parameter = append(parameter, db.AdvertisementResults.TotalReaction.Set(*input.TotalReaction))
	}

	if input.TotalViews != nil {
		parameter = append(parameter, db.AdvertisementResults.TotalViews.Set(*input.TotalViews))
	}

	createdAdvertisementResult, err := client.AdvertisementResults.CreateOne(
		db.AdvertisementResults.StartedAt.Set(input.StartedAt),
		db.AdvertisementResults.Advertisement.Link(db.Advertisements.ID.Equals(input.AdvertisementID)),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdAdvertisementResult, nil
}
