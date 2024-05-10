package companyRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/companyService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateCompanyInput) (*model.Companies, error) {

	updatedData, err := companyService.Update(companyService.UpdateInput{
		ID:         input.ID,
		Name:       input.Name,
		UniqueName: input.UniqueName,
		TinNumber:  input.TinNumber,
		VatNumber:  input.VatNumber,
	})
	if err != nil {
		return nil, err
	}

	updatedCompany := utils.CompanyGraphqlConverter(updatedData)
	return updatedCompany, nil
}
