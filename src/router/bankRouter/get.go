package bankRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/bankService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Banks, error) {
	bankdata, err := bankService.Get(id)

	if err != nil || bankdata == nil {
		return nil, err
	}

	bank := utils.BankGraphqlConverter(bankdata)

	return bank, nil

}
