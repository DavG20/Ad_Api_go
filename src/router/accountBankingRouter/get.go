package accountBankingRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/accountBankingService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.AccountBanking, error) {
	accountBankingData, err := accountBankingService.Get(id)

	if err != nil || accountBankingData == nil {
		return nil, nil
	}

	accountBanking := utils.AccountBankingGraphqlConverter(accountBankingData)

	return accountBanking, nil

}
