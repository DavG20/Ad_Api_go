package channelContentLanguageService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/service/languageService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input ChannelContentLanguageInput) (*db.ChannelContentLanguagesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	language, _ := languageService.Get(input.Language)
	if language == nil {
		return nil, errors.New("language doesn't exist")
	}

	channel, _ := channelService.Get(input.ChannelID)
	if channel == nil {
		return nil, errors.New("channel doesn't exist")
	}

	contentLanguage, _ := Get(input)

	if contentLanguage != nil {
		return nil, errors.New("language and channel id is already registered")
	}

	client, ctx, err := database.GetChannelClient()
	if err != nil {
		return nil, err
	}
	createdContentLanguage, err := client.ChannelContentLanguages.CreateOne(
		db.ChannelContentLanguages.LanguageData.Link(db.Languages.Language.Equals(input.Language)),
		db.ChannelContentLanguages.Channel.Link(db.Channels.ID.Equals(input.ChannelID)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdContentLanguage, nil
}
