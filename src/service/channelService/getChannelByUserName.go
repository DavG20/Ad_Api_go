package channelService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
)

func GetChannelByUserName(userName string) (*db.ChannelsModel, error) {
	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	channel, err := client.Channels.FindUnique(
		db.Channels.UserName.Equals(userName),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return channel, nil
}
