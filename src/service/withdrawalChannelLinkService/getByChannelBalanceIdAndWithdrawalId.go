package withdrawalChannelLinkService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
)

func GetByChannelBalanceIdAndWithdrawalId(input WithdrawalChannelLinkInput) (*db.WithdrawalChannelLinksModel, error) {
	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	withdrawalChannelLink, err := client.WithdrawalChannelLinks.FindUnique(
		db.WithdrawalChannelLinks.ID(
			db.WithdrawalChannelLinks.WithdrawalID.Equals(input.WithdrawalID),
			db.WithdrawalChannelLinks.ChannelBalanceID.Equals(input.ChannelBalanceID),
		)).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return withdrawalChannelLink, nil
}
