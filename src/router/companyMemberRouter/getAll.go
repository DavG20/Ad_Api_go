package companyMemberRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/companyMemberService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.CompanyMemberFilter) ([]*model.CompanyMembers, error) {

	companyMemberData, err := companyMemberService.GetAll(convertor.CompanyMemberFilterConverter(input))

	if err != nil {
		return nil, err
	}
	companyMembers := []*model.CompanyMembers{}

	for _, companyMember := range companyMemberData {
		companyMember := utils.CompanyMemberGraphqlConverter(companyMember)
		companyMembers = append(companyMembers, companyMember)
	}
	return companyMembers, nil

}
