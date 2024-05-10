package companyBankingService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/service/bankService"
	"adtec/backend/src/service/companyService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
	"log"
)

func Create(input CreateInput) (*db.CompanyBankingsModel, error) {

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
	company, _ := companyService.Get(input.CompanyID)

	if company == nil {
		return nil, errors.New("invalid company id")
	}

	//check bankname and bankaccount

	companyBanking, _ := GetByBankAccountAndId(input.BankAccount, input.BankId)

	if companyBanking != nil {
		return nil, errors.New("bankAccount and bankId already registered")
	}

	createdCompanyBanking, err := client.CompanyBankings.CreateOne(
		db.CompanyBankings.FullNameOnBank.Set(input.FullNameOnBank),
		db.CompanyBankings.BankAccount.Set(input.BankAccount),
		db.CompanyBankings.Bank.Link(db.Banks.ID.Equals(input.BankId)),
		db.CompanyBankings.Company.Link(db.Companies.ID.Equals(input.CompanyID)),
		db.CompanyBankings.Currency.Set(input.Currency),
	).Exec(ctx)

	if err != nil {
		log.Println("companyBank creation database side error", input.FullNameOnBank, "error:", err)
		return nil, err
	}

	return createdCompanyBanking, nil

}
