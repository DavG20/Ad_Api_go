package channelContentLanguageService

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Delete(input ChannelContentLanguageInput) (*model.DeletionResult, error) {
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

	channelLanguage, _ := Get(input)
	if channelLanguage == nil {
		deletionResult.Success = false
		deletionResult.Message = "channelLanguage not found"
		return deletionResult, errors.New("channelLanguage not found")
	}

	if channelLanguage != nil {
		_, err := client.ChannelContentLanguages.FindUnique(
			db.ChannelContentLanguages.ChannelContentLanguageID(
				db.ChannelContentLanguages.ChannelID.Equals(input.ChannelID),
				db.ChannelContentLanguages.Language.Equals(input.Language),
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
