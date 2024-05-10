package companyBankingRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/companyBankingService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.CompanyBankings, error) {
	companyBankingData, err := companyBankingService.Get(id)

	if err != nil || companyBankingData == nil {
		return nil, nil
	}

	companyBanking := utils.CompanyBankingGraphqlConverter(companyBankingData)

	return companyBanking, nil

}
