package companyBankingService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
)

func GetByBankAccountAndId(bankAccount, bankId string) (*db.CompanyBankingsModel, error) {
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}
	companyBankings, err := client.CompanyBankings.FindUnique(
		db.CompanyBankings.UniqueBankAccountAndId(
			db.CompanyBankings.BankAccount.Equals(bankAccount),
			db.CompanyBankings.BankID.Equals(bankId),
		),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return companyBankings, err
}
