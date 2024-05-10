package accountBankingService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/service/bankService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
	"log"
)

func Create(input CreateInput) (*db.AccountBankingsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	client, ctx, err := database.GetAccountClient()
	if err != nil {
		return nil, err
	}

	bank, _ := bankService.Get(input.BankId)

	if bank == nil {
		return nil, errors.New("invalid bank id")
	}
	account, _ := accountService.Get(input.AccountId)

	if account == nil {
		return nil, errors.New("invalid account id")
	}

	//check bankName and bankAccount

	accountBanking, _ := GetByBankAccountAndId(input.BankAccount, input.BankId)

	if accountBanking != nil {
		log.Println("bankAccount and bankName already registered")
		return nil, errors.New("bankAccount and bankId already registered")
	}
	createdAccount, err := client.AccountBankings.CreateOne(
		db.AccountBankings.Bank.Link(db.Banks.ID.Equals(input.BankId)),
		db.AccountBankings.FullNameOnBank.Set(input.FullNameOnBank),
		db.AccountBankings.BankAccount.Set(input.BankAccount),
		db.AccountBankings.Currency.Set(input.Currency),
		db.AccountBankings.Account.Link(db.Accounts.ID.Equals(input.AccountId)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdAccount, nil

}
