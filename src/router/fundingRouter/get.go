package fundingRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/fundingService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Funding, error) {
	fundingData, err := fundingService.Get(id)
	if err != nil {
		return nil, err
	}

	funding := utils.FundingGraphqlConverter(fundingData)

	return funding, nil
}
