package channelCollectedAdLinkService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
)

func GetByChannelIdAndAdPaymentId(input ChannelCollectedAdLinkInput) (*db.ChannelCollectedAdLinksModel, error) {
	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	channelCollectedAdLink, err := client.ChannelCollectedAdLinks.FindUnique(
		db.ChannelCollectedAdLinks.ID(
			db.ChannelCollectedAdLinks.ChannelBalanceID.Equals(input.ChannelBalanceID),
			db.ChannelCollectedAdLinks.AdPaymentID.Equals(input.AdPaymentID),
		)).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return channelCollectedAdLink, nil
}
