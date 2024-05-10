package channelCollectedAdLinkService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
)

func Get(adPaymentId string) (*db.ChannelCollectedAdLinksModel, error) {
	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	channelCollectedAdLink, err := client.ChannelCollectedAdLinks.FindUnique(
		db.ChannelCollectedAdLinks.AdPaymentID.Equals(adPaymentId)).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return channelCollectedAdLink, nil
}
