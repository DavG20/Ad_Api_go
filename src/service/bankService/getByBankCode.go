package bankService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
)

func GetByBankCode(bankCode string) (*db.BanksModel, error) {
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	bank, err := client.Banks.FindUnique(
		db.Banks.BankCode.Equals(bankCode),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return bank, nil
}
