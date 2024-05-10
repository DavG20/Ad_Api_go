package fundingRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/fundingService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.FundingFilter) ([]*model.Funding, error) {

	fundingData, err := fundingService.GetAll(convertor.FundingFilterConverter(input))

	if err != nil {
		return nil, err
	}
	fundings := []*model.Funding{}

	for _, funding := range fundingData {
		funding := utils.FundingGraphqlConverter(funding)
		fundings = append(fundings, funding)
	}
	return fundings, nil

}
