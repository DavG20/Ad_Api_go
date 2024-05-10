package withdrawalRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/withdrawalService"
	"adtec/backend/src/utils"
	"context"
	"time"
)

func Update(ctx context.Context, input model.UpdateWithdrawalInput) (*model.Withdrawals, error) {
	var paymentDate *time.Time

	if input.PaymentDate != nil {
		date, _ := time.Parse("2006-01-02", *input.PaymentDate)

		paymentDate = &date

	}

	updatedData, err := withdrawalService.Update(withdrawalService.UpdateInput{
		ID:          input.ID,
		TotalAmount: input.TotalAmount,
		Currency:    input.Currency,
		FundingKey:  input.FundingKey,
		Method:      input.Method,
		Reference:   input.Reference,
		PaymentDate: paymentDate,
		Status:      (*service.FundingStatus)(input.Status),
		Extra:       &input.Extra,
		Log:         &input.Log,
	})
	if err != nil {
		return nil, err
	}

	updatedWithdrawal := utils.WithdrawalGraphqlConverter(updatedData)
	return updatedWithdrawal, nil
}
