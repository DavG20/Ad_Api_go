package channelDetailService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.ChannelDetailsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	channelDetail, _ := Get(input.ChannelID)

	if channelDetail == nil {
		return nil, errors.New("invalid channelDetail id or channelDetail doesn't exist")
	}

	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.ChannelDetailsSetParam{}

	// have to check the existence of ownerID
	if input.SubCount != nil {
		if *input.SubCount != channelDetail.SubCount {
			parameter = append(parameter, db.ChannelDetails.SubCount.Set(*input.SubCount))

		}
	}
	if input.AveragePostView != nil {
		if *input.AveragePostView != channelDetail.AveragePostView {
			parameter = append(parameter, db.ChannelDetails.AveragePostView.Set(*input.AveragePostView))

		}
	}
	if input.PostFrequency != nil {
		if *input.PostFrequency != channelDetail.PostFrequency {
			parameter = append(parameter, db.ChannelDetails.PostFrequency.Set(*input.PostFrequency))

		}
	}
	if input.LastPost != nil {
		lastPost, _ := channelDetail.LastPost()
		if *input.LastPost != lastPost {
			parameter = append(parameter, db.ChannelDetails.LastPost.Set(*input.LastPost))
		}
	}
	if input.LastPostID != nil {

		lastPostId, _ := channelDetail.LastPostID()
		if *input.LastPostID != lastPostId {
			parameter = append(parameter, db.ChannelDetails.LastPostID.Set(*input.LastPostID))

		}
	}
	if input.CollectedDate != nil {

		collectedDate, _ := channelDetail.CollectedDate()
		if *input.CollectedDate != collectedDate {
			parameter = append(parameter, db.ChannelDetails.CollectedDate.Set(*input.CollectedDate))

		}
	}

	updateChannelDetail, err := client.ChannelDetails.FindUnique(
		db.ChannelDetails.ChannelID.Equals(input.ChannelID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateChannelDetail, nil
}
