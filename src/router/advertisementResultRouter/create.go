package advertisementResultRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/advertisementResultService"
	"adtec/backend/src/utils"
	"context"
	"time"
)

func Create(ctx context.Context, input model.CreateAdvertisementResultInput) (*model.AdvertisementResults, error) {
	var startedAt time.Time
	date, _ := time.Parse("2006-01-02", input.StartedAt)
	startedAt = date

	advertisementResultData, err := advertisementResultService.Create(
		advertisementResultService.CreateInput{
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
		},
	)

	if err != nil {
		return nil, err
	}
	advertisementResult := utils.AdvertisementResultGraphqlConverter(advertisementResultData)
	return advertisementResult, nil
}
