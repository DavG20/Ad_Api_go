package withdrawalChannelLinkService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
)

func GetAll(filter WithdrawalChannelLinkFilter) ([]db.WithdrawalChannelLinksModel, error) {

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}
	parameter := []db.WithdrawalChannelLinksWhereParam{}

	if filter.ChannelBalanceID != nil {
		parameter = append(parameter, db.WithdrawalChannelLinks.ChannelBalanceID.Equals(*filter.ChannelBalanceID))
	}

	fetchWithdrawalChannelLinks := client.WithdrawalChannelLinks.FindMany(parameter[:]...).Take(int(limit))
	// TODO check this one , is this the best approach

	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchWithdrawalChannelLinks = client.WithdrawalChannelLinks.FindMany(parameter[:]...).Skip(1).Take(int(limit)).Cursor(
				db.WithdrawalChannelLinks.WithdrawalID.Cursor(string(*filter.Filter.After)),
			)
		}
		// TODO check this one , is this the best approach
		if filter.Filter.Before != nil {
			fetchWithdrawalChannelLinks = client.WithdrawalChannelLinks.FindMany(parameter[:]...).Skip(1).Take(int(limit)).Cursor(
				db.WithdrawalChannelLinks.WithdrawalID.Cursor(string(*filter.Filter.Before)),
			)
		}
	}
	withdrawalChannelLinks, err := fetchWithdrawalChannelLinks.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return withdrawalChannelLinks, nil

}
