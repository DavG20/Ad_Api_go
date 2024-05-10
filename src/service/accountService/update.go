package accountService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/service"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.AccountsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}
	

	account, _ := Get(input.ID)

	if account == nil {
		return nil, errors.New("account not found")
	}

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.AccountsSetParam{}

	if input.PhoneNumber != nil {
		phoneExist, _ := GetByPhoneNumber(*input.PhoneNumber)
		phoneNumber, _ := account.PhoneNumber()

		if phoneExist != nil {
			return nil, errors.New("phoneNumber is taken")
		}
		if *input.PhoneNumber != phoneNumber {

			parameter = append(parameter, db.Accounts.PhoneNumber.Set(*input.PhoneNumber))
		}

	}
	if input.UserName != nil {
		userNameExist, _ := GetByUserName(*input.UserName)

		if userNameExist != nil {
			return nil, errors.New("username is taken")
		}
		if *input.UserName != account.UserName {
			parameter = append(parameter, db.Accounts.UserName.Set(*input.UserName))
		}

	}
	if input.Email != nil {
		emailExist, _ := GetByEmail(*input.Email)
		email, _ := account.Email()
		if emailExist != nil {
			return nil, errors.New("email is taken")
		}
		if *input.Email != email {
			parameter = append(parameter, db.Accounts.Email.Set(*input.Email))
		}

	}

	if input.FullName != nil {
		fullName, _ := account.FullName()
		if *input.FullName != fullName {
			parameter = append(parameter, db.Accounts.FullName.Set(*input.FullName))
		}
	}
	if input.BirthDate != nil {
		birthDate, _ := account.BirthDate()
		if *input.BirthDate != birthDate {
			parameter = append(parameter, db.Accounts.BirthDate.Set(*input.BirthDate))
		}
	}
	if input.AccountType != nil {
		if *input.AccountType != service.AccountType(*input.AccountType) {
			parameter = append(parameter, db.Accounts.AccountType.Set(db.AccountType(*input.AccountType)))

		}
	}

	updatedAccount, err := client.Accounts.FindUnique(
		db.Accounts.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updatedAccount, nil
}
