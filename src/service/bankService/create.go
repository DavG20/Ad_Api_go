package bankService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
	"log"
)

func Create(input CreateInput) (*db.BanksModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	bank, _ := GetByBankCode(input.BankCode)

	if bank != nil {
		return nil, errors.New("bank code already exists")
	}

	bank, _ = GetByBankName(input.BankName)

	if bank != nil {
		return nil, errors.New("bank name already exists")
	}

	createdBank, err := client.Banks.CreateOne(
		db.Banks.BankName.Set(input.BankName),
		db.Banks.BankCode.Set(input.BankCode),
	).Exec(ctx)

	if err != nil {
		log.Println("bank creation database side error", input.BankName, "error:", err)
		return nil, err
	}

	return createdBank, nil

}
