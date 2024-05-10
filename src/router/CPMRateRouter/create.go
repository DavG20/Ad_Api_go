package CPMRateRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/CPMRateService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateCPMRateInput) (*model.CpmRates, error) {
	cpmRateData, err := CPMRateService.Create(
		CPMRateService.CreateInput{
			ChannelID:    input.ChannelID,
			Active:       input.Active,
			Cpm:          input.Cpm,
			MinCPMVolume: input.MinCPMVolume,
		},
	)
	if err != nil {
		return nil, err
	}

	cpmRate := utils.CPMRateGraphqlConverter(cpmRateData)
	return cpmRate, nil

}
