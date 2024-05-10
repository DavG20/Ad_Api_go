package adPaymentService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
)

func GetAll(filter AdPaymentFilter) ([]db.AdPaymentsModel, error) {

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	parameter := []db.AdPaymentsWhereParam{}

	if filter.CampaignID != nil {
		parameter = append(parameter, db.AdPayments.CampaignID.Equals(*filter.CampaignID))
	}
	if filter.AdvertisementID != nil {
		parameter = append(parameter, db.AdPayments.AdvertisementID.Equals(*filter.AdvertisementID))
	}
	if filter.ChannelID != nil {
		parameter = append(parameter, db.AdPayments.ChannelID.Equals(*filter.ChannelID))
	}
	if filter.Currency != nil {
		parameter = append(parameter, db.AdPayments.Currency.Equals(*filter.Currency))
	}
	if filter.Amount != nil {
		if filter.Amount.Max != nil {
			parameter = append(parameter, db.AdPayments.Amount.Lte(*filter.Amount.Max))
		}
		parameter = append(parameter, db.AdPayments.Amount.Gte(filter.Amount.Min))

	}

	fetchAdPayments := client.AdPayments.FindMany(
		parameter[:]...,
	).Take(int(limit)).OrderBy(
		db.AdPayments.CreatedAt.Order(db.DESC),
	)
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchAdPayments = client.AdPayments.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.AdPayments.ID.Cursor(string(*filter.Filter.After)),
			).OrderBy(
				db.AdPayments.CreatedAt.Order(db.DESC),
			)
		}
		if filter.Filter.Before != nil {
			fetchAdPayments = client.AdPayments.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.AdPayments.ID.Cursor(string(*filter.Filter.Before)),
			).OrderBy(
				db.AdPayments.CreatedAt.Order(db.DESC),
			)
		}
	}
	adPayments, err := fetchAdPayments.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return adPayments, nil

}
