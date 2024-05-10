package bankRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/bankService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateBankInput) (*model.Banks, error) {
	updatedBankdata, err := bankService.Update(
		bankService.UpdateInput{
			ID:       input.ID,
			BankName: input.BankName,
			BankCode: input.BankCode,
		},
	)

	if err != nil {
		return nil, err
	}

	bank := utils.BankGraphqlConverter(updatedBankdata)
	return bank, nil
}
