package channelCategoryService

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Delete(input ChannelCategoryInput) (*model.DeletionResult, error) {
	//TODO: Should I add validation here?

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

	channelCategory, _ := Get(input)
	if channelCategory == nil {
		deletionResult.Success = false
		deletionResult.Message = "ChannelCategory not Found"
		return deletionResult, errors.New("ChannelCategory not Found")
	}

	if channelCategory != nil {
		_, err := client.ChannelCategories.FindUnique(
			db.ChannelCategories.ChannelCategoryID(
				db.ChannelCategories.ChannelID.Equals(input.ChannelID),
				db.ChannelCategories.Category.Equals(input.Category),
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
