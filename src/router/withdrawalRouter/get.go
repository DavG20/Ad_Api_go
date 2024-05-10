package withdrawalRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/withdrawalService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Withdrawals, error) {
	withdrawalData, err := withdrawalService.Get(id)
	if err != nil {
		return nil, err
	}

	withdrawal := utils.WithdrawalGraphqlConverter(withdrawalData)

	return withdrawal, nil
}
