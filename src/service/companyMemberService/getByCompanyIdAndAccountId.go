package companyMemberService

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
)

func GetByCompanyIdAndAccountId(companyId, accountId string) (*db.CompanyMembersModel, error) {
	client, ctx, err := database.GetAccountClient()

	if err != nil {
		return nil, err
	}

	companyMember, err := client.CompanyMembers.FindUnique(
		db.CompanyMembers.UniqueAccountIdAndCompanyId(
			db.CompanyMembers.AccountID.Equals(accountId),
			db.CompanyMembers.CompanyID.Equals(companyId),
		),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return companyMember, nil
}
