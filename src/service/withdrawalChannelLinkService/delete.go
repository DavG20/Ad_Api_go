package withdrawalChannelLinkService

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Delete(input WithdrawalChannelLinkInput) (*model.DeletionResult, error) {

	client, ctx, err := database.GetTransactionClient()

	deletionResult := &model.DeletionResult{
		Success: true,
		Message: "deleted Successfully",
	}
	if err != nil {
		deletionResult.Message = "internal error of db"
		deletionResult.Success = false
		return deletionResult, err
	}

	withdrawalChannelLink, _ := GetByChannelBalanceIdAndWithdrawalId(input)
	if withdrawalChannelLink == nil {
		deletionResult.Success = false
		deletionResult.Message = "WithdrawalChannelLink not Found"
		return deletionResult, errors.New("WithdrawalChannelLink not Found")
	}

	if withdrawalChannelLink != nil {
		_, err := client.WithdrawalChannelLinks.FindUnique(
			db.WithdrawalChannelLinks.ID(
				db.WithdrawalChannelLinks.WithdrawalID.Equals(input.WithdrawalID),
				db.WithdrawalChannelLinks.ChannelBalanceID.Equals(input.ChannelBalanceID),
			),
		).Delete().Exec(ctx)

		if err != nil {
			deletionResult.Success = false
			deletionResult.Message = "error deleting"
			return deletionResult, errors.New("error deleting")
		}
	}
	return deletionResult, nil
}
