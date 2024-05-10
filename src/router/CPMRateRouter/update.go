package CPMRateRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/CPMRateService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateCPMRateInput) (*model.CpmRates, error) {
	updatedCPMRatedata, err := CPMRateService.Update(
		CPMRateService.UpdateInput{
			ChannelID:    input.ChannelID,
			Active:       input.Active,
			Cpm:          input.Cpm,
			MinCPMVolume: input.MinCPMVolume,
		},
	)

	if err != nil {
		return nil, err
	}

	cpmRate := utils.CPMRateGraphqlConverter(updatedCPMRatedata)
	return cpmRate, nil
}
