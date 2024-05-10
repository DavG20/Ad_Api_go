package companyMemberRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/companyMemberService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateCompanyMemberInput) (*model.CompanyMembers, error) {

	updatedData, err := companyMemberService.Update(companyMemberService.UpdateInput{
		ID:   input.ID,
		Role: (*service.CompanyRole)(input.Role),
	})
	if err != nil {
		return nil, err
	}

	updatedCompanyMember := utils.CompanyMemberGraphqlConverter(updatedData)
	return updatedCompanyMember, nil
}
