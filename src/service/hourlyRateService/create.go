package hourlyRateService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.HourlyRatesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.HourlyRatesSetParam{}

	if input.Active != nil {
		parameter = append(parameter, db.HourlyRates.Active.Set(*input.Active))
	}
	if input.HourlyRate != nil {
		parameter = append(parameter, db.HourlyRates.HourlyRate.Set(*input.HourlyRate))
	}
	if input.MinHourlyVolume != nil {
		parameter = append(parameter, db.HourlyRates.MinHourlyVolume.Set(*input.MinHourlyVolume))
	}
	if input.MaxHourlyVolume != nil {
		parameter = append(parameter, db.HourlyRates.MaxHourlyVolume.Set(*input.MaxHourlyVolume))
	}

	createdHourlyRate, err := client.HourlyRates.CreateOne(
		db.HourlyRates.Channel.Link(db.Channels.ID.Equals(input.ChannelID)),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdHourlyRate, nil
}
