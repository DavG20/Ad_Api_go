package channelContentLanguageRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelContentLanguageService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.ChannelContentLanguageFilter) ([]*model.ChannelContentLanguages, error) {

	channelLanguageData, err := channelContentLanguageService.GetAll(convertor.ChannelContentLanguageFilterConverter(input))

	if err != nil {
		return nil, err
	}
	channelLanguages := []*model.ChannelContentLanguages{}

	for _, channelCategory := range channelLanguageData {
		channelLanguage := utils.ChannelContentLanguageGraphqlConverter(channelCategory)
		channelLanguages = append(channelLanguages, channelLanguage)
	}
	return channelLanguages, nil

}
