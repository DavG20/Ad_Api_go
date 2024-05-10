package channelDetailRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelDetailService"
	"adtec/backend/src/utils"
	"context"
	"time"
)

func Update(ctx context.Context, input model.UpdateChannelDetailInput) (*model.ChannelDetails, error) {
	var lastPost *time.Time
	if input.LastPost != nil {

		date, _ := time.Parse("2006-01-02", *input.LastPost)
		lastPost = &date
	}
	var collectedDate *time.Time

	if input.CollectedDate != nil {
		date, _ := time.Parse("2006-01-02", *input.CollectedDate)
		collectedDate = &date
	}

	updateChannelDetailData, err := channelDetailService.Update(
		channelDetailService.UpdateInput{
			ChannelID:       input.ChannelID,
			SubCount:        input.SubCount,
			AveragePostView: input.AveragePostView,
			PostToSubRatio:  input.PostToSubRatio,
			PostFrequency:   input.PostFrequency,
			LastPostID:      input.LastPostID,
			LastPost:        lastPost,
			CollectedDate:   collectedDate,
		},
	)
	if err != nil {
		return nil, err
	}

	updateChannelDetail := utils.ChannelDetailGraphqlConverter(updateChannelDetailData)

	return updateChannelDetail, nil

}
