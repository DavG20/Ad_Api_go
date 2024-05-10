package withdrawalChannelLinkService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
)

func Get(withdrawalId string) (*db.WithdrawalChannelLinksModel, error) {
	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	withdrawalChannelLinks, err := client.WithdrawalChannelLinks.FindUnique(
		db.WithdrawalChannelLinks.WithdrawalID.Equals(withdrawalId)).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return withdrawalChannelLinks, nil
}
