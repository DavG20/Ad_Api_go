package companyMemberRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/companyMemberService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.CompanyMembers, error) {
	companyMemberData, err := companyMemberService.Get(id)
	if err != nil {
		return nil, err
	}

	companyMember := utils.CompanyMemberGraphqlConverter(companyMemberData)

	return companyMember, nil
}
