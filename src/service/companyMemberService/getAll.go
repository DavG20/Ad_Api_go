package companyMemberService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/enumConverter"
)

func GetAll(filter CompanyMemberFilter) ([]db.CompanyMembersModel, error) {

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	parameter := []db.CompanyMembersWhereParam{}
	if filter.CompanyID != nil {
		parameter = append(
			parameter,
			db.CompanyMembers.Company.Where(db.Companies.ID.Equals(*filter.CompanyID)))
	}
	if filter.AccountID != nil {
		parameter = append(
			parameter, db.CompanyMembers.Account.Where(db.Accounts.ID.Equals(*filter.AccountID)))
	}

	if filter.Role != nil {
		parameter = append(parameter, db.CompanyMembers.Role.In(enumConverter.CompanyRoleConverter(filter.Role)))

	}

	fetchCompanyMembers := client.CompanyMembers.FindMany(
		parameter[:]...,
	).Take(int(limit)).OrderBy(
		db.CompanyMembers.CreatedAt.Order(db.DESC),
	)
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchCompanyMembers = client.CompanyMembers.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.CompanyMembers.ID.Cursor(string(*filter.Filter.After)),
			).OrderBy(
				db.CompanyMembers.CreatedAt.Order(db.DESC),
			)
		}
		if filter.Filter.Before != nil {
			fetchCompanyMembers = client.CompanyMembers.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.CompanyMembers.ID.Cursor(string(*filter.Filter.Before)),
			).OrderBy(
				db.CompanyMembers.CreatedAt.Order(db.DESC),
			)
		}
	}
	companyMembers, err := fetchCompanyMembers.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return companyMembers, nil

}
