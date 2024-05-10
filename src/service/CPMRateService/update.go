package CPMRateService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.CpmRatesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	cpmRate, _ := Get(input.ChannelID)

	if cpmRate == nil {
		return nil, errors.New("invalid ChannelId or CPMRate doesn't exist")
	}

	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.CpmRatesSetParam{}

	if input.Active != nil {
		parameter = append(parameter, db.CpmRates.Active.Set(*input.Active))
	}
	if input.Cpm != nil {
		parameter = append(parameter, db.CpmRates.Cpm.Set(*input.Cpm))
	}
	if input.MinCPMVolume != nil {
		parameter = append(parameter, db.CpmRates.MinCpmvolume.Set(*input.MinCPMVolume))
	}

	updateCPMRate, err := client.CpmRates.FindUnique(
		db.CpmRates.ChannelID.Equals(input.ChannelID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateCPMRate, nil
}
