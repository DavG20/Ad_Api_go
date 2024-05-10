package bankService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
)

func GetAll(filter FindManyInput) ([]db.BanksModel, error) {

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	fetchBank := client.Banks.FindMany().Take(int(limit)).OrderBy(
		db.Banks.CreatedAt.Order(db.DESC),
	)
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchBank = client.Banks.FindMany().Skip(1).Take(int(limit)).Cursor(
				db.Banks.ID.Cursor(string(*filter.Filter.After)),
			).OrderBy(
				db.Banks.CreatedAt.Order(db.DESC),
			)
		}
		if filter.Filter.Before != nil {
			fetchBank = client.Banks.FindMany().Skip(1).Take(int(limit)).Cursor(
				db.Banks.ID.Cursor(string(*filter.Filter.Before)),
			).OrderBy(
				db.Banks.CreatedAt.Order(db.DESC),
			)
		}
	}
	banks, err := fetchBank.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return banks, nil

}
