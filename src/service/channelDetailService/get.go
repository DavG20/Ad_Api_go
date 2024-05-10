package channelDetailService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
)

func Get(channelId string) (*db.ChannelDetailsModel, error) {
	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	channelDetail, err := client.ChannelDetails.FindUnique(
		db.ChannelDetails.ChannelID.Equals(channelId),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return channelDetail, nil
}
