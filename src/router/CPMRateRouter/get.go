package CPMRateRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/CPMRateService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, channelId string) (*model.CpmRates, error) {
	cpmRatedata, err := CPMRateService.Get(channelId)

	if err != nil || cpmRatedata == nil {
		return nil, err
	}

	cpmRate := utils.CPMRateGraphqlConverter(cpmRatedata)

	return cpmRate, nil

}
