package companyBankingRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/companyBankingService"
	"adtec/backend/src/utils"
	"context"
)

func GetAll(ctx context.Context, input model.CompanyBankingFilter) ([]*model.CompanyBankings, error) {
	inputs := companyBankingService.FindManyInput{
		CompanyID: input.CompanyID,
		Filter:    (*service.FilterData)(input.Filter),
		BankId:    input.BankID,
	}
	companyBankingsData, err := companyBankingService.GetAll(inputs)

	if err != nil {
		return nil, err
	}
	companyBankings := []*model.CompanyBankings{}

	for _, companyBanking := range companyBankingsData {
		companyBanking := utils.CompanyBankingGraphqlConverter(companyBanking)
		companyBankings = append(companyBankings, companyBanking)
	}
	return companyBankings, nil

}
