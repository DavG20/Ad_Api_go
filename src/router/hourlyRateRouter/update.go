package hourlyRateRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/hourlyRateService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateHourlyRateInput) (*model.HourlyRates, error) {
	updatedHourlyRatedata, err := hourlyRateService.Update(
		hourlyRateService.UpdateInput{
			ID:              input.ID,
			Active:          input.Active,
			HourlyRate:      input.HourlyRate,
			MinHourlyVolume: input.MinHourlyVolume,
			MaxHourlyVolume: input.MaxHourlyVolume,
		},
	)

	if err != nil {
		return nil, err
	}

	hourlyRate := utils.HourlyRateGraphqlConverter(updatedHourlyRatedata)
	return hourlyRate, nil
}
