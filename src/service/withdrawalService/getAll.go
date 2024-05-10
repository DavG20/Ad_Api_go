package withdrawalService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/enumConverter"
)

func GetAll(filter WithdrawalFilter) ([]db.WithdrawalsModel, error) {

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	parameter := []db.WithdrawalsWhereParam{}

	if filter.Status != nil {
		parameter = append(parameter, db.Withdrawals.Status.In(enumConverter.FundingStatusConverter(filter.Status)))
	}

	if filter.AccountID != nil {
		parameter = append(parameter, db.Withdrawals.AccountID.Equals(*filter.AccountID))
	}
	if filter.AccountBankID != nil {
		parameter = append(parameter, db.Withdrawals.AccountBankID.Equals(*filter.AccountBankID))
	}
	if filter.Currency != nil {
		parameter = append(parameter, db.Withdrawals.Currency.Equals(*filter.Currency))
	}
	if filter.TotalAmount != nil {
		if filter.TotalAmount.Max != nil {
			parameter = append(parameter, db.Withdrawals.TotalAmount.Lte(*filter.TotalAmount.Max))
		}
		parameter = append(parameter, db.Withdrawals.TotalAmount.Gte(filter.TotalAmount.Min))

	}

	fetchWithdrawals := client.Withdrawals.FindMany(
		parameter[:]...,
	).Take(int(limit)).OrderBy(
		db.Withdrawals.CreatedAt.Order(db.DESC),
	)
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchWithdrawals = client.Withdrawals.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.Withdrawals.ID.Cursor(string(*filter.Filter.After)),
			).OrderBy(
				db.Withdrawals.CreatedAt.Order(db.DESC),
			)
		}
		if filter.Filter.Before != nil {
			fetchWithdrawals = client.Withdrawals.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.Withdrawals.ID.Cursor(string(*filter.Filter.Before)),
			).OrderBy(
				db.Withdrawals.CreatedAt.Order(db.DESC),
			)
		}
	}
	withdrawals, err := fetchWithdrawals.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return withdrawals, nil

}
