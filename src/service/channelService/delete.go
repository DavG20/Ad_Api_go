package channelService

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Delete(id string) (*model.DeletionResult, error) {

	client, ctx, err := database.GetChannelClient()

	deletionResult := &model.DeletionResult{
		Success: true,
		Message: "deleted Successfully",
	}

	if err != nil {
		deletionResult.Message = "internal error of db"
		deletionResult.Success = false
		return deletionResult, err
	}
	channel, _ := Get(id)

	if channel == nil {
		deletionResult.Success = false
		deletionResult.Message = "channel not Found"
		return deletionResult, errors.New("channel not Found")
	}

	_, err = client.Channels.FindUnique(
		db.Channels.ID.Equals(id),
	).Delete().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return deletionResult, nil
}
