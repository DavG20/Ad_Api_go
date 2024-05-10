package withdrawalChannelLinkService

import (
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/service/channelBalanceService"
	"adtec/backend/src/service/withdrawalService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input WithdrawalChannelLinkInput) (*db.WithdrawalChannelLinksModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	channelBalance, _ := channelBalanceService.GetById(input.ChannelBalanceID)
	if channelBalance == nil {
		return nil, errors.New("channelBalance doesn't exist")
	}

	withdrawal, _ := withdrawalService.Get(input.WithdrawalID)
	if withdrawal == nil {
		return nil, errors.New("withdrawal doesn't exist")
	}

	WithdrawalChannelLinks, _ := GetByChannelBalanceIdAndWithdrawalId(input)

	if WithdrawalChannelLinks != nil {
		return nil, errors.New("channelBalance id and withdrawal id is already registered")
	}
	withdrawalChannelLink, _ := Get(input.WithdrawalID)

	if withdrawalChannelLink != nil {
		return nil, errors.New("withdrawal id is already registered")
	}
	client, ctx, err := database.GetTransactionClient()
	if err != nil {
		return nil, err
	}
	createdWithdrawalChannelLink, err := client.WithdrawalChannelLinks.CreateOne(
		db.WithdrawalChannelLinks.Withdrawal.Link(db.Withdrawals.ID.Equals(input.WithdrawalID)),
		db.WithdrawalChannelLinks.ChannelBalance.Link(db.ChannelBalances.ID.Equals(input.ChannelBalanceID)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdWithdrawalChannelLink, nil
}
