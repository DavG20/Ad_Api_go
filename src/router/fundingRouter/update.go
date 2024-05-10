package fundingRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/fundingService"
	"adtec/backend/src/utils"
	"context"
	"time"
)

func Update(ctx context.Context, input model.UpdateFundingInput) (*model.Funding, error) {
	var paymentDate *time.Time

	if input.PaymentDate != nil {
		date, _ := time.Parse("2006-01-02", *input.PaymentDate)

		paymentDate = &date

	}

	updatedData, err := fundingService.Update(fundingService.UpdateInput{
		ID:           input.ID,
		Amount:       input.Amount,
		Currency:     input.Currency,
		Key:          input.Key,
		Method:       input.Method,
		FundingTxRef: input.FundingTxRef,
		Reference:    input.Reference,
		PaymentDate:  paymentDate,
		RedirectURL:  input.RedirectURL,
		Status:       (*service.FundingStatus)(input.Status),
		Extra:        &input.Extra,
		Log:          &input.Log,
		Tax:          input.Tax,
	})
	if err != nil {
		return nil, err
	}

	updatedFunding := utils.FundingGraphqlConverter(updatedData)
	return updatedFunding, nil
}
