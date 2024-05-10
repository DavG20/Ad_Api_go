package companyMemberService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/service/companyService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.CompanyMembersModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	client, ctx, err := database.GetAccountClient()
	if err != nil {
		return nil, err
	}

	//check bankname and bankaccount

	// companyMember, _ := GetByCompanyIdAndAccountId(input.CompanyID, input.AccountID)

	// if companyMember != nil {
	// 	return nil, errors.New("companyId and accountId already registered")
	// }
	account, _ := accountService.Get(input.AccountID)

	if account == nil {
		return nil, errors.New("invalid account id")
	}
	company, _ := companyService.Get(input.CompanyID)

	if company == nil {
		return nil, errors.New("invalid company id")
	}

	parameter := []db.CompanyMembersSetParam{}

	if input.Role != nil {
		parameter = append(parameter, db.CompanyMembers.Role.Set(db.CompanyRole(*input.Role)))
	}

	createdCompanyMember, err := client.CompanyMembers.CreateOne(
		db.CompanyMembers.Company.Link(db.Companies.ID.Equals(input.CompanyID)),
		db.CompanyMembers.Account.Link(db.Accounts.ID.Equals(input.AccountID)),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdCompanyMember, nil

}
