package channelCollectedAdLinkService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
)

func GetAll(filter ChannelCollectedAdLinkFilter) ([]db.ChannelCollectedAdLinksModel, error) {

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}
	parameter := []db.ChannelCollectedAdLinksWhereParam{}

	if filter.ChannelBalanceID != nil {
		parameter = append(parameter, db.ChannelCollectedAdLinks.ChannelBalanceID.Equals(*filter.ChannelBalanceID))
	}

	fetchChannelCollectedAdLinks := client.ChannelCollectedAdLinks.FindMany(parameter[:]...).Take(int(limit))
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchChannelCollectedAdLinks = client.ChannelCollectedAdLinks.FindMany(parameter[:]...).Skip(1).Take(int(limit)).Cursor(
				db.ChannelCollectedAdLinks.AdPaymentID.Cursor(string(*filter.Filter.After)),
			)
		}
		// TODO check this one , is this the best approach
		if filter.Filter.Before != nil {
			fetchChannelCollectedAdLinks = client.ChannelCollectedAdLinks.FindMany(parameter[:]...).Skip(1).Take(int(limit)).Cursor(
				db.ChannelCollectedAdLinks.AdPaymentID.Cursor(string(*filter.Filter.Before)),
			)
		}
	}
	channelCollectedAdLinks, err := fetchChannelCollectedAdLinks.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return channelCollectedAdLinks, nil

}
