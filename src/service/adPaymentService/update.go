package adPaymentService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.AdPaymentsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	adPayment, _ := Get(input.ID)

	if adPayment == nil {
		return nil, errors.New("invalid adPayment id or adPayment doesn't exist")
	}

	client, ctx, err := database.GetTransactionClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.AdPaymentsSetParam{}

	if input.Amount != nil {
		if *input.Amount != adPayment.Amount {
			parameter = append(parameter, db.AdPayments.Amount.Set(*input.Amount))
		}
	}
	if input.Currency != nil {
		if *input.Currency != adPayment.Currency {
			parameter = append(parameter, db.AdPayments.Currency.Set(*input.Currency))
		}
	}

	if input.CompletionTime != nil {
		if *input.CompletionTime != adPayment.CompletionTime {
			parameter = append(parameter, db.AdPayments.CompletionTime.Set(*input.CompletionTime))
		}
	}

	updateAdPayment, err := client.AdPayments.FindUnique(
		db.AdPayments.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateAdPayment, nil
}
