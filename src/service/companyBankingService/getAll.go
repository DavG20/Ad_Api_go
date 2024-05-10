package companyBankingService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
)

func GetAll(filter FindManyInput) ([]db.CompanyBankingsModel, error) {

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	parameter := []db.CompanyBankingsWhereParam{}
	if filter.CompanyID != nil {
		parameter = append(
			parameter,
			db.CompanyBankings.Company.Where(
				db.Companies.ID.Equals(*filter.CompanyID),
			),
		)
	}
	if filter.BankId != nil {
		parameter = append(
			parameter,
			db.CompanyBankings.Bank.Where(db.Banks.ID.Equals(*filter.BankId)),
		)
	}

	fetchCompanyBankings := client.CompanyBankings.FindMany(
		parameter[:]...,
	).Take(int(limit)).OrderBy(
		db.CompanyBankings.CreatedAt.Order(db.DESC),
	)
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchCompanyBankings = client.CompanyBankings.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.CompanyBankings.ID.Cursor(string(*filter.Filter.After)),
			).OrderBy(
				db.CompanyBankings.CreatedAt.Order(db.DESC),
			)
		}
		if filter.Filter.Before != nil {
			fetchCompanyBankings = client.CompanyBankings.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.CompanyBankings.ID.Cursor(string(*filter.Filter.Before)),
			).OrderBy(
				db.CompanyBankings.CreatedAt.Order(db.DESC),
			)
		}
	}
	companyBankings, err := fetchCompanyBankings.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return companyBankings, nil

}
