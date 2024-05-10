package accountBankingService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
)

func GetByBankAccountAndId(bankAccount, bankId string) (*db.AccountBankingsModel, error) {
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}
	accountBanking, err := client.AccountBankings.FindUnique(
		db.AccountBankings.UniqueBankAccountAndId(
			db.AccountBankings.BankAccount.Equals(bankAccount),
			db.AccountBankings.BankID.Equals(bankId),
		),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return accountBanking, err
}
