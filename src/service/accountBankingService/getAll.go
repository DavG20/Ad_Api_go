package accountBankingService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
)

func GetAll(filter FindManyInput) ([]db.AccountBankingsModel, error) {

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	parameter := []db.AccountBankingsWhereParam{}
	if filter.AccountId != nil {
		parameter = append(
			parameter,
			db.AccountBankings.Account.Where(db.Accounts.ID.Equals(*filter.AccountId)),
		)
	}
	if filter.BankId != nil {
		parameter = append(
			parameter,
			db.AccountBankings.Bank.Where(db.Banks.ID.Equals(*filter.BankId)),
		)
	}

	fetchAccountBanking := client.AccountBankings.FindMany(
		parameter[:]...,
	).Take(int(limit)).OrderBy(
		db.AccountBankings.CreatedAt.Order(db.DESC),
	)
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchAccountBanking = client.AccountBankings.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.AccountBankings.ID.Cursor(string(*filter.Filter.After)),
			).OrderBy(
				db.AccountBankings.CreatedAt.Order(db.DESC),
			)
		}
		if filter.Filter.Before != nil {
			fetchAccountBanking = client.AccountBankings.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.AccountBankings.ID.Cursor(string(*filter.Filter.Before)),
			).OrderBy(
				db.AccountBankings.CreatedAt.Order(db.DESC),
			)
		}
	}
	accountBankings, err := fetchAccountBanking.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return accountBankings, nil

}
