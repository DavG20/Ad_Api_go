package advertisementResultService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.AdvertisementResultsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	advertisementResult, _ := Get(input.AdvertisementID)

	if advertisementResult == nil {
		return nil, errors.New("invalid AdvertisementResult id or AdvertisementResult doesn't exist")
	}

	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.AdvertisementResultsSetParam{}

	if input.StartedAt != nil {
		if *input.StartedAt != advertisementResult.StartedAt {
			parameter = append(parameter, db.AdvertisementResults.StartedAt.Set(*input.StartedAt))
		}
	}
	if input.Budget != nil {
		if *input.Budget != advertisementResult.Budget {
			parameter = append(parameter, db.AdvertisementResults.Budget.Set(*input.Budget))
		}
	}

	if input.ProviderBudgetShare != nil {
		if *input.ProviderBudgetShare != advertisementResult.ProviderBudgetShare {
			parameter = append(parameter, db.AdvertisementResults.ProviderBudgetShare.Set(*input.ProviderBudgetShare))
		}
	}

	if input.ChannelBudgetShare != nil {
		if *input.ChannelBudgetShare != advertisementResult.ChannelBudgetShare {
			parameter = append(parameter, db.AdvertisementResults.ChannelBudgetShare.Set(*input.ChannelBudgetShare))
		}
	}

	if input.TotalClick != nil {
		parameter = append(parameter, db.AdvertisementResults.TotalClick.Set(*input.TotalClick))
	}

	if input.TotalForward != nil {
		if *input.TotalForward != advertisementResult.TotalForward {
			parameter = append(parameter, db.AdvertisementResults.TotalForward.Set(*input.TotalForward))
		}
	}

	if input.TotalHour != nil {
		if *input.TotalHour != advertisementResult.TotalHour {
			parameter = append(parameter, db.AdvertisementResults.TotalHour.Set(*input.TotalHour))
		}
	}

	if input.TotalReaction != nil {
		if *input.TotalReaction != advertisementResult.TotalReaction {
			parameter = append(parameter, db.AdvertisementResults.TotalReaction.Set(*input.TotalReaction))
		}
	}

	if input.TotalViews != nil {
		if *input.TotalViews != advertisementResult.TotalViews {
			parameter = append(parameter, db.AdvertisementResults.TotalViews.Set(*input.TotalViews))
		}
	}

	updateContent, err := client.AdvertisementResults.FindUnique(
		db.AdvertisementResults.AdvertisementID.Equals(input.AdvertisementID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateContent, nil
}
