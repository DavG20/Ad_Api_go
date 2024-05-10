package channelRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/CPMRateService"
	"adtec/backend/src/service/channelCategoryService"
	"adtec/backend/src/service/channelContentLanguageService"
	"adtec/backend/src/service/channelDetailService"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/utils"
	"context"
	"errors"
	"time"
)

func Create(ctx context.Context, input model.CreateChannelInput) (*model.Channels, error) {
	var channelCreatedAt *time.Time
	if input.ChannelCreatedAt != nil {
		date, _ := time.Parse("2006-01-02", *input.ChannelCreatedAt)
		channelCreatedAt = &date
	}
	channelData, err := channelService.Create(
		channelService.CreateInput{
			OwnerID:          input.OwnerID,
			UserName:         input.UserName,
			Name:             input.Name,
			Description:      input.Description,
			BotAddAsAdmin:    input.BotAddAsAdmin,
			ChannelCreatedAt: channelCreatedAt,
			Country:          input.Country,
			Currency:         input.Currency,
		},
	)
	if err != nil {
		return nil, err
	}

	if input.Languages != nil && len(input.Languages) > 0 {
		languageValues := make([]channelContentLanguageService.ChannelContentLanguages, 0, len(input.Languages))
		for _, language := range input.Languages {
			channelLanguage := channelContentLanguageService.ChannelContentLanguages{
				ChannelID: channelData.ID,
				Language:  language,
			}

			languageValues = append(languageValues, channelLanguage)
		}
		_, err := channelContentLanguageService.CreateMany(languageValues)
		if err != nil {
			channelService.Delete(channelData.ID)
			return nil, err
		}
	}

	if input.Categories != nil && len(input.Categories) > 0 {

		categoryValues := make([]channelCategoryService.ChannelCategories, 0, len(input.Categories))
		for _, category := range input.Categories {
			channelCategory := channelCategoryService.ChannelCategories{
				ChannelId: channelData.ID,
				Category:  category,
			}
			categoryValues = append(categoryValues, channelCategory)
		}
		_, err := channelCategoryService.CreateMany(categoryValues)
		if err != nil {
			channelService.Delete(channelData.ID)
			return nil, err
		}

	}
	if input.CpmRate != nil {
		input.CpmRate.ChannelID = channelData.ID

		cpmRate, err := CPMRateService.Create(CPMRateService.CreateInput(*input.CpmRate))
		if err != nil || cpmRate == nil {
			channelService.Delete(channelData.ID)
			return nil, errors.New("error creating cpmRate")
		}
	}
	// TODO: integrate the service that fetch the needed data

	channelDetail, err := channelDetailService.Create(channelDetailService.CreateInput{
		ChannelID:     channelData.ID,
		SubCount:      new(int),
		PostFrequency: new(float64),
	})
	if err != nil || channelDetail == nil {
		channelService.Delete(channelData.ID)
		return nil, errors.New("error creating ChannelDetail")
	}

	createdChannel := utils.ChannelGraphqlConverter(channelData)

	return createdChannel, nil
}
