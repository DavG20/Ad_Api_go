package channelContentLanguageRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelContentLanguageService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.ChannelContentLanguageInput) (*model.ChannelContentLanguages, error) {
	createdChannelContentLanguage, err := channelContentLanguageService.Create(
		channelContentLanguageService.ChannelContentLanguageInput{
			ChannelID: input.ChannelID,
			Language:  input.Language,
		},
	)
	if err != nil {
		return nil, err
	}

	channelLanguage := utils.ChannelContentLanguageGraphqlConverter(createdChannelContentLanguage)
	return channelLanguage, nil
}
