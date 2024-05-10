package hourlyRateService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.HourlyRatesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	hourlyRate, _ := Get(input.ID)

	if hourlyRate == nil {
		return nil, errors.New("invalide HourlyRate id or HourlyRate doesn't exist")
	}

	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.HourlyRatesSetParam{}

	if input.Active != nil {

		if *input.Active != hourlyRate.Active {
			parameter = append(parameter, db.HourlyRates.Active.Set(*input.Active))
		}

	}
	if input.HourlyRate != nil {
		if *input.HourlyRate != hourlyRate.HourlyRate {
			parameter = append(parameter, db.HourlyRates.HourlyRate.Set(*input.HourlyRate))

		}
	}
	if input.MinHourlyVolume != nil {
		if *input.MinHourlyVolume != hourlyRate.MaxHourlyVolume {
			parameter = append(parameter, db.HourlyRates.MinHourlyVolume.Set(*input.MinHourlyVolume))

		}
	}
	if input.MaxHourlyVolume != nil {
		if *input.MaxHourlyVolume != hourlyRate.MaxHourlyVolume {
			parameter = append(parameter, db.HourlyRates.MaxHourlyVolume.Set(*input.MaxHourlyVolume))
		}
	}

	updateHourlyRate, err := client.HourlyRates.FindUnique(
		db.HourlyRates.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateHourlyRate, nil
}
