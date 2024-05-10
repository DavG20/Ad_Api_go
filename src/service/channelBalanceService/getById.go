package channelBalanceService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func GetById(id string) (*db.ChannelBalancesModel, error) {
	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}
	channelBalance, err := client.ChannelBalances.FindUnique(
		db.ChannelBalances.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("channelBalance not found")
	}

	return channelBalance, nil
}
