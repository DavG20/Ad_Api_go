package bankService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
)

func GetByBankName(bankName string) (*db.BanksModel, error) {
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	bank, err := client.Banks.FindUnique(
		db.Banks.BankName.Equals(bankName),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return bank, nil
}
