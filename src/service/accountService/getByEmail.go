package accountService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
)

func GetByEmail(email string) (*db.AccountsModel, error) {
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	account, err := client.Accounts.FindUnique(
		db.Accounts.Email.Equals(email),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return account, nil

}
