package bankRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/bankService"
	"adtec/backend/src/utils"
	"context"
)

func GetAll(ctx context.Context, input model.BankFilter) ([]*model.Banks, error) {
	inputs := bankService.FindManyInput{
		Filter: (*service.FilterData)(input.Filter),
	}
	bankingData, err := bankService.GetAll(inputs)

	if err != nil {
		return nil, err
	}
	banks := []*model.Banks{}

	for _, bank := range bankingData {
		bank := utils.BankGraphqlConverter(bank)
		banks = append(banks, bank)
	}
	return banks, nil

}
