package hourlyRateRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/hourlyRateService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.HourlyRates, error) {
	hourlyRatedata, err := hourlyRateService.Get(id)

	if err != nil || hourlyRatedata == nil {
		return nil, err
	}

	hourlyRate := utils.HourlyRateGraphqlConverter(hourlyRatedata)

	return hourlyRate, nil

}
