package bankRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/bankService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateBankInput) (*model.Banks, error) {
	bankData, err := bankService.Create(
		bankService.CreateInput{
			BankName: input.BankName,
			BankCode: input.BankCode,
		},
	)
	if err != nil {
		return nil, err
	}

	bank := utils.BankGraphqlConverter(bankData)
	return bank, nil

}
