package channelContentLanguageService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
)

func Get(input ChannelContentLanguageInput) (*db.ChannelContentLanguagesModel, error) {
	//TODO: Should I add validation here?

	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	contentLanguage, err := client.ChannelContentLanguages.FindUnique(
		db.ChannelContentLanguages.ChannelContentLanguageID(
			db.ChannelContentLanguages.ChannelID.Equals(input.ChannelID),
			db.ChannelContentLanguages.Language.Equals(input.Language),
		)).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return contentLanguage, nil
}
