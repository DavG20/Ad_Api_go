package companyLocationRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/companyLocationService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateCompanyLocationInput) (*model.CompanyLocations, error) {
	createdCompanyLocation, err := companyLocationService.Create(
		companyLocationService.CreateInput{
			CompanyID:  input.CompanyID,
			LocationID: input.LocationID,
		},
	)
	if err != nil {
		return nil, err
	}

	companyLocation := utils.CompanyLocationGraphqlConverter(createdCompanyLocation)
	return companyLocation, nil
}
