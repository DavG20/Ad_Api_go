package advertisementResultRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/advertisementResultService"
	"adtec/backend/src/utils"
	"context"
	"time"
)

func Update(ctx context.Context, input model.UpdateAdvertisementResultInput) (*model.AdvertisementResults, error) {
	var startedAt *time.Time

	if input.StartedAt != nil {
		date, _ := time.Parse("2006-01-02", *input.StartedAt)
		startedAt = &date
	}

	updatedData, err := advertisementResultService.Update(advertisementResultService.UpdateInput{
		AdvertisementID:     input.AdvertisementID,
		StartedAt:           startedAt,
		Budget:              input.Budget,
		ProviderBudgetShare: input.ProviderBudgetShare,
		ChannelBudgetShare:  input.ChannelBudgetShare,
		TotalHour:           input.TotalHour,
		TotalClick:          input.TotalClick,
		TotalViews:          input.TotalViews,
		TotalForward:        input.TotalForward,
		TotalReaction:       input.TotalReaction,
	})
	if err != nil {
		return nil, err
	}

	updatedAdvertisementResult := utils.AdvertisementResultGraphqlConverter(updatedData)
	return updatedAdvertisementResult, nil
}
