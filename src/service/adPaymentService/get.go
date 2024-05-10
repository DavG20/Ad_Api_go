package adPaymentService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Get(advertisementId string) (*db.AdPaymentsModel, error) {
	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	adPayment, err := client.AdPayments.FindUnique(
		db.AdPayments.AdvertisementID.Equals(advertisementId),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("adPayments not found")
	}
	return adPayment, nil
}
