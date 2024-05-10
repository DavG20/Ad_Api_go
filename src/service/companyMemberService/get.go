package companyMemberService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
)

func Get(id string) (*db.CompanyMembersModel, error) {

	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	companyMember, err := client.CompanyMembers.FindUnique(
		db.CompanyMembers.ID.Equals(id),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}
	return companyMember, nil

}
