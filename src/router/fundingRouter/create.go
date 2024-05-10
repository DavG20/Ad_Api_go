package fundingRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/companyService"
	"adtec/backend/src/service/fundingService"
	"adtec/backend/src/utils"
	"context"
	"errors"
	"time"
)

func Create(ctx context.Context, input model.CreateFundingInput) (*model.Funding, error) {
	company, _ := companyService.Get(input.CompanyID)

	if company == nil {
		return nil, errors.New("invalid company Id ")
	}
	var paymentDate *time.Time

	if input.PaymentDate != nil {
		date, _ := time.Parse("2006-01-02", *input.PaymentDate)

		paymentDate = &date
	}

	fundingData, err := fundingService.Create(
		fundingService.CreateInput{
			CompanyID:    input.CompanyID,
			Amount:       input.Amount,
			Currency:     *input.Currency,
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
		},
	)
	if err != nil {
		return nil, err
	}

	createdFunding := utils.FundingGraphqlConverter(fundingData)

	return createdFunding, nil
}
