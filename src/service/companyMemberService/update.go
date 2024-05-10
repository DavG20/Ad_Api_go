package companyMemberService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/service"
	"adtec/backend/src/utils/database"
	"errors"
)

func Update(input UpdateInput) (*db.CompanyMembersModel, error) {

	companyMember, _ := Get(input.ID)

	if companyMember == nil {
		return nil, errors.New("companyMember not found")
	}

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.CompanyMembersSetParam{}

	if input.Role != nil {
		if *input.Role != service.CompanyRole(companyMember.Role) {
			parameter = append(parameter, db.CompanyMembers.Role.Set(db.CompanyRole(*input.Role)))

		}
	}

	updatedCompanyMember, err := client.CompanyMembers.FindUnique(
		db.CompanyMembers.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updatedCompanyMember, nil
}
