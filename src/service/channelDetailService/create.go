package channelDetailService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.ChannelDetailsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	channel, err := channelService.Get(input.ChannelID)

	if err != nil || channel == nil {
		return nil, errors.New("channel doesn't exist")
	}

	channelDetail, _ := Get(input.ChannelID)

	if channelDetail != nil {
		return nil, errors.New("channel id  Taken")
	}
	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.ChannelDetailsSetParam{}

	if input.SubCount != nil {
		parameter = append(parameter, db.ChannelDetails.SubCount.Set(*input.SubCount))
	}
	if input.AveragePostView != nil {
		parameter = append(parameter, db.ChannelDetails.AveragePostView.Set(*input.AveragePostView))
	}
	if input.PostToSubRatio != nil {
		parameter = append(parameter, db.ChannelDetails.PostToSubRatio.Set(*input.PostToSubRatio))
	}
	if input.PostFrequency != nil {
		parameter = append(parameter, db.ChannelDetails.PostFrequency.Set(*input.PostFrequency))
	}
	if input.LastPostID != nil {
		parameter = append(parameter, db.ChannelDetails.LastPostID.Set(*input.LastPostID))
	}
	if input.LastPost != nil {
		parameter = append(parameter, db.ChannelDetails.LastPost.Set(*input.LastPost))
	}

	createdChannelDetail, err := client.ChannelDetails.CreateOne(
		db.ChannelDetails.Channel.Link(db.Channels.ID.Equals(input.ChannelID)),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdChannelDetail, nil
}
