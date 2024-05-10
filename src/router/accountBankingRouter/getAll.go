package accountBankingRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/accountBankingService"
	"adtec/backend/src/utils"
	"context"
)

func GetAll(ctx context.Context, input model.AccountBankingsFilter) ([]*model.AccountBanking, error) {
	inputs := accountBankingService.FindManyInput{
		AccountId: input.AccountID,
		Filter:    (*service.FilterData)(input.Filter),
		BankId:    input.BankID,
	}
	accountBankingsData, err := accountBankingService.GetAll(inputs)

	if err != nil {
		return nil, err
	}
	accountBankings := []*model.AccountBanking{}

	for _, accountBanking := range accountBankingsData {
		accountBanking := utils.AccountBankingGraphqlConverter(accountBanking)
		accountBankings = append(accountBankings, accountBanking)
	}
	return accountBankings, nil

}
