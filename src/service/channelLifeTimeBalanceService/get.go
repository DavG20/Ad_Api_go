package channelLifeTimeBalanceService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Get(id string) (*db.ChannelLifeTimeBalancesModel, error) {
	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}
	channelLifeTimeBalance, err := client.ChannelLifeTimeBalances.FindUnique(
		db.ChannelLifeTimeBalances.ChannelID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("channelLifeTimeBalance not found")
	}

	return channelLifeTimeBalance, nil
}
