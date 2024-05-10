package fundingService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/enumConverter"
)

func GetAll(filter FundingFilter) ([]db.FundingModel, error) {

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	parameter := []db.FundingWhereParam{}

	if filter.Status != nil {
		parameter = append(parameter, db.Funding.Status.In(enumConverter.FundingStatusConverter(filter.Status)))
	}

	if filter.CompanyID != nil {
		parameter = append(parameter, db.Funding.CompanyID.Equals(*filter.CompanyID))
	}
	if filter.Currency != nil {
		parameter = append(parameter, db.Funding.Currency.Equals(*filter.Currency))
	}
	if filter.Amount != nil {
		if filter.Amount.Max != nil {
			parameter = append(parameter, db.Funding.Amount.Lte(*filter.Amount.Max))
		}
		parameter = append(parameter, db.Funding.Amount.Gte(filter.Amount.Min))

	}

	fetchFunding := client.Funding.FindMany(
		parameter[:]...,
	).Take(int(limit)).OrderBy(
		db.Funding.CreatedAt.Order(db.DESC),
	)
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchFunding = client.Funding.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.Funding.ID.Cursor(string(*filter.Filter.After)),
			).OrderBy(
				db.Funding.CreatedAt.Order(db.DESC),
			)
		}
		if filter.Filter.Before != nil {
			fetchFunding = client.Funding.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.Funding.ID.Cursor(string(*filter.Filter.Before)),
			).OrderBy(
				db.Funding.CreatedAt.Order(db.DESC),
			)
		}
	}
	funding, err := fetchFunding.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return funding, nil

}
