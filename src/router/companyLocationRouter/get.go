package companyLocationRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/companyLocationService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, companyId, locationId string) (*model.CompanyLocations, error) {
	companyLocationData, err := companyLocationService.Get(companyId, locationId)

	if err != nil {
		return nil, err
	}

	companyLocation := utils.CompanyLocationGraphqlConverter(companyLocationData)

	return companyLocation, nil

}
