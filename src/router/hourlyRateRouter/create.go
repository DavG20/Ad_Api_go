package hourlyRateRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/hourlyRateService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateHourlyRateInput) (*model.HourlyRates, error) {
	hourlyRateData, err := hourlyRateService.Create(
		hourlyRateService.CreateInput{
			ChannelID:       input.ChannelID,
			Active:          input.Active,
			HourlyRate:      input.HourlyRate,
			MinHourlyVolume: input.MinHourlyVolume,
			MaxHourlyVolume: input.MaxHourlyVolume,
		},
	)
	if err != nil {
		return nil, err
	}

	hourlyRate := utils.HourlyRateGraphqlConverter(hourlyRateData)
	return hourlyRate, nil

}
