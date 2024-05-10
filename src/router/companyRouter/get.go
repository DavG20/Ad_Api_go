package companyRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/companyService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Companies, error) {
	companyData, err := companyService.Get(id)
	if err != nil {
		return nil, err
	}

	company := utils.CompanyGraphqlConverter(companyData)

	return company, nil
}
