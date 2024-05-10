package withdrawalRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/withdrawalService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.WithdrawalFilter) ([]*model.Withdrawals, error) {

	withdrawalData, err := withdrawalService.GetAll(convertor.WithdrawalFilterConverter(input))

	if err != nil {
		return nil, err
	}
	withdrawals := []*model.Withdrawals{}

	for _, withdrawal := range withdrawalData {
		withdrawal := utils.WithdrawalGraphqlConverter(withdrawal)
		withdrawals = append(withdrawals, withdrawal)
	}
	return withdrawals, nil

}
