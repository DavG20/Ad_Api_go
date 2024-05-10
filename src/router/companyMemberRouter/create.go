package companyMemberRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/companyMemberService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateCompanyMemberInput) (*model.CompanyMembers, error) {

	companyMemberData, err := companyMemberService.Create(
		companyMemberService.CreateInput{
			AccountID: input.AccountID,
			CompanyID: input.CompanyID,
			Role:      (*service.CompanyRole)(input.Role),
		},
	)
	if err != nil {
		return nil, err
	}

	createdCompanyMember := utils.CompanyMemberGraphqlConverter(companyMemberData)

	return createdCompanyMember, nil
}
