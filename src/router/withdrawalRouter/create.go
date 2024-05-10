package withdrawalRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/accountBankingService"
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/service/withdrawalService"
	"adtec/backend/src/utils"
	"context"
	"errors"
	"time"
)

func Create(ctx context.Context, input model.CreateWithdrawalInput) (*model.Withdrawals, error) {
	account, _ := accountService.Get(input.AccountID)

	if account == nil {
		return nil, errors.New("invalid account Id ")
	}
	accountBanking, _ := accountBankingService.Get(input.AccountBankID)

	if accountBanking == nil {
		return nil, errors.New("invalid accountBanking Id ")
	}
	var paymentDate *time.Time

	if input.PaymentDate != nil {
		date, _ := time.Parse("2006-01-02", *input.PaymentDate)

		paymentDate = &date
	}

	WithdrawalData, err := withdrawalService.Create(
		withdrawalService.CreateInput{
			AccountID:     input.AccountID,
			AccountBankID: input.AccountBankID,
			TotalAmount:   input.TotalAmount,
			Currency:      input.Currency,
			FundingKey:    input.FundingKey,
			Method:        input.Method,
			Reference:     input.Reference,
			PaymentDate:   paymentDate,
			Status:        (*service.FundingStatus)(input.Status),
			Extra:         input.Extra,
			Log:           input.Log,
		},
	)
	if err != nil {
		return nil, err
	}

	createdWithdrawal := utils.WithdrawalGraphqlConverter(WithdrawalData)

	return createdWithdrawal, nil
}
