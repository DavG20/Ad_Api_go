package companyRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/companyService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateCompanyInput) (*model.Companies, error) {

	contentData, err := companyService.Create(
		companyService.CreateInput{
			CreatorID:  input.CreatorID,
			Name:       input.Name,
			UniqueName: input.UniqueName,
			TinNumber:  input.TinNumber,
			VatNumber:  input.VatNumber,
		},
	)
	if err != nil {
		return nil, err
	}
	createdCompany := utils.CompanyGraphqlConverter(contentData)

	return createdCompany, nil
}
