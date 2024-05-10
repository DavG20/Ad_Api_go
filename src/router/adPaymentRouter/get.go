package adPaymentRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/adPaymentService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, advertisementId string) (*model.AdPayments, error) {
	adPaymentData, err := adPaymentService.Get(advertisementId)
	if err != nil {
		return nil, err
	}

	adPayment := utils.AdPaymentGraphqlConverter(adPaymentData)

	return adPayment, nil
}
