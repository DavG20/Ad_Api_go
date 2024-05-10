package companyService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.CompaniesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	// TODO check creatorId
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}
	companyByUniqueName, _ := GetByUniqueName(input.UniqueName)
	if companyByUniqueName != nil {
		return nil, errors.New("uniqueName is taken , use another name")
	}
	parameter := []db.CompaniesSetParam{}

	if input.TinNumber != nil {
		parameter = append(parameter, db.Companies.TinNumber.Set(*input.TinNumber))
	}
	if input.VatNumber != nil {
		parameter = append(parameter, db.Companies.VatNumber.Set(*input.VatNumber))
	}

	createdCompany, err := client.Companies.CreateOne(
		db.Companies.CreatorID.Set(input.CreatorID),
		db.Companies.Name.Set(input.Name),
		db.Companies.UniqueName.Set(input.UniqueName),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdCompany, nil
}
