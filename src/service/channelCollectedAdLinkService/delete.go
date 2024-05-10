package channelCollectedAdLinkService

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/prisma/transaction/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Delete(input ChannelCollectedAdLinkInput) (*model.DeletionResult, error) {

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

	channelCollectedAdLink, _ := GetByChannelIdAndAdPaymentId(input)
	if channelCollectedAdLink == nil {
		deletionResult.Success = false
		deletionResult.Message = "channelCollectedAdLink not Found"
		return deletionResult, errors.New("channelCollectedAdLink not Found")
	}

	if channelCollectedAdLink != nil {
		_, err := client.ChannelCollectedAdLinks.FindUnique(
			db.ChannelCollectedAdLinks.ID(
				db.ChannelCollectedAdLinks.ChannelBalanceID.Equals(input.ChannelBalanceID),
				db.ChannelCollectedAdLinks.AdPaymentID.Equals(input.AdPaymentID),
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
