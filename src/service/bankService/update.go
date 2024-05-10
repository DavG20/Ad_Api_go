package bankService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.BanksModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}
	banks, _ := Get(input.ID)

	if banks == nil {
		return nil, errors.New("invalid bank ID")
	}

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.BanksSetParam{}
	if input.BankCode != nil {
		if *input.BankCode != banks.BankCode {
			parameter = append(parameter, db.Banks.BankCode.Set(*input.BankCode))
		}
	}
	if input.BankName != nil {
		if *input.BankName != banks.BankName {
			parameter = append(parameter, db.Banks.BankName.Set(*input.BankName))
		}
	}

	updatedBank, err := client.Banks.FindUnique(
		db.Banks.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return updatedBank, nil

}
