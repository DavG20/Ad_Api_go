package CPMRateService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
)

func Get(channelId string) (*db.CpmRatesModel, error) {
	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	cpmRate, err := client.CpmRates.FindUnique(
		db.CpmRates.ChannelID.Equals(channelId),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return cpmRate, nil

}
