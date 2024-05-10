package adPaymentRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/adPaymentService"
	"adtec/backend/src/utils"
	"context"
	"time"
)

func Update(ctx context.Context, input model.UpdateAdPaymentInput) (*model.AdPayments, error) {
	var completionTime *time.Time

	if input.CompletionTime != nil {
		date, _ := time.Parse("2006-01-02", *input.CompletionTime)

		completionTime = &date

	}

	updatedData, err := adPaymentService.Update(adPaymentService.UpdateInput{
		ID:             input.ID,
		Amount:         input.Amount,
		Currency:       input.Currency,
		CompletionTime: completionTime,
	})
	if err != nil {
		return nil, err
	}

	updatedAdPayment := utils.AdPaymentGraphqlConverter(updatedData)

	return updatedAdPayment, nil
}
