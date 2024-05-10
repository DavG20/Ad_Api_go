package hourlyRateService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
)

func Get(id string) (*db.HourlyRatesModel, error) {
	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	hourlyRate, err := client.HourlyRates.FindUnique(
		db.HourlyRates.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return hourlyRate, nil
}
