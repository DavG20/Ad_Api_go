package accountBankingRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/accountBankingService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateAccountBankingInput) (*model.AccountBanking, error) {
	accountData, err := accountBankingService.Create(
		accountBankingService.CreateInput{
			AccountId:      input.AccountID,
			BankId:         input.BankID,
			FullNameOnBank: input.FullNameOnBank,
			BankAccount:    input.BankAccount,
			Currency:       input.Currency,
		})
	if err != nil {
		return nil, err
	}

	accountBanking := utils.AccountBankingGraphqlConverter(accountData)

	return accountBanking, nil

}
