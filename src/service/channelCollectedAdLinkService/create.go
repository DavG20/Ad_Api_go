package channelCollectedAdLinkService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/service/adPaymentService"
	"adtec/backend/src/service/channelBalanceService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input ChannelCollectedAdLinkInput) (*db.ChannelCollectedAdLinksModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	channelBalance, _ := channelBalanceService.GetById(input.ChannelBalanceID)
	if channelBalance == nil {
		return nil, errors.New("channelBalance doesn't exist")
	}

	adPayment, _ := adPaymentService.GetById(input.AdPaymentID)
	if adPayment == nil {
		return nil, errors.New("adPayment doesn't exist")
	}

	channelCollectedAdLinks, _ := GetByChannelIdAndAdPaymentId(input)

	if channelCollectedAdLinks != nil {
		return nil, errors.New("channelBalance id and AdPayment id is already registered")
	}
	channelCollectedAdLink, _ := Get(input.AdPaymentID)

	if channelCollectedAdLink != nil {
		return nil, errors.New("AdPayment id is already registered")
	}
	client, ctx, err := database.GetTransactionClient()
	if err != nil {
		return nil, err
	}
	createdChannelCollectedAdLinks, err := client.ChannelCollectedAdLinks.CreateOne(
		db.ChannelCollectedAdLinks.ChannelBalance.Link(db.ChannelBalances.ID.Equals(input.ChannelBalanceID)),
		db.ChannelCollectedAdLinks.AdPayment.Link(db.AdPayments.ID.Equals(input.AdPaymentID)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdChannelCollectedAdLinks, nil
}
