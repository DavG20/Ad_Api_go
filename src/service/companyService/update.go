package companyService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.CompaniesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	company, _ := Get(input.ID)

	if company == nil {
		return nil, errors.New("invalid company id or company doesn't exist")
	}

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.CompaniesSetParam{}

	if input.Name != nil {
		if *input.Name != company.Name {
			parameter = append(parameter, db.Companies.Name.Set(*input.Name))

		}
	}
	if input.UniqueName != nil {
		if *input.UniqueName != company.UniqueName {
			parameter = append(parameter, db.Companies.UniqueName.Set(*input.UniqueName))

		}
	}

	if input.TinNumber != nil {
		tinNumber, _ := company.TinNumber()
		if *input.TinNumber != tinNumber {
			parameter = append(parameter, db.Companies.TinNumber.Set(*input.TinNumber))

		}
	}

	if input.VatNumber != nil {
		vatNumber, _ := company.VatNumber()
		if *input.VatNumber != vatNumber {
			parameter = append(parameter, db.Companies.VatNumber.Set(*input.VatNumber))

		}
	}

	updateCompany, err := client.Companies.FindUnique(
		db.Companies.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateCompany, nil
}
