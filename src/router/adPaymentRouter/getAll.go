package adPaymentRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/adPaymentService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.AdPaymentFilter) ([]*model.AdPayments, error) {

	adPaymentData, err := adPaymentService.GetAll(convertor.AdPaymentFilterConverter(input))

	if err != nil {
		return nil, err
	}
	adPayments := []*model.AdPayments{}

	for _, adPayment := range adPaymentData {
		adPayment := utils.AdPaymentGraphqlConverter(adPayment)
		adPayments = append(adPayments, adPayment)
	}
	return adPayments, nil

}
