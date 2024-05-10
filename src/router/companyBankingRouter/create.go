package companyBankingRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/companyBankingService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateCompanyBankingInput) (*model.CompanyBankings, error) {
	companyBankingData, err := companyBankingService.Create(
		companyBankingService.CreateInput{
			CompanyID:      input.CompanyID,
			BankId:         input.BankID,
			FullNameOnBank: input.FullNameOnBank,
			BankAccount:    input.BankAccount,
			Currency:       input.Currency,
		})
	if err != nil {
		return nil, err
	}

	companyBanking := utils.CompanyBankingGraphqlConverter(companyBankingData)

	return companyBanking, nil

}
