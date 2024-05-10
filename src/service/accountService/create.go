package accountService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
	"log"
)

func Create(input CreateInput) (*db.AccountsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}
	account, _ := GetByUserID(input.UserID)

	if account != nil {
		return nil, errors.New("UserId already Registered")
	}

	userNameExist, _ := GetByUserName(input.UserName)

	if userNameExist != nil {
		return nil, errors.New("UserName is Taken")
	}

	parameter := []db.AccountsSetParam{}

	if input.Email != nil {
		emailExist, _ := GetByEmail(*input.Email)

		if emailExist != nil {
			return nil, errors.New("email is taken")
		}
		parameter = append(parameter, db.Accounts.Email.Set(*input.Email))

	}
	if input.PhoneNumber != nil {
		phoneExist, _ := GetByPhoneNumber(*input.PhoneNumber)

		if phoneExist != nil {
			return nil, errors.New("phone is taken")
		}
		parameter = append(parameter, db.Accounts.PhoneNumber.Set(*input.PhoneNumber))
	}

	if input.FullName != nil {
		parameter = append(parameter, db.Accounts.FullName.Set(*input.FullName))
	}

	if input.BirthDate != nil {
		parameter = append(parameter, db.Accounts.BirthDate.Set(*input.BirthDate))
	}
	if input.AccountType != nil {
		parameter = append(parameter, db.Accounts.AccountType.Set(db.AccountType(*input.AccountType)))
	}
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	createdAccount, err := client.Accounts.CreateOne(
		db.Accounts.UserID.Set(input.UserID),
		db.Accounts.UserName.Set(input.UserName),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		log.Println("Account creation database side error", input.UserName)
		return nil, err
	}
	return createdAccount, nil

}
