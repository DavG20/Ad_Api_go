package channelService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Get(id string) (*db.ChannelsModel, error) {
	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	channel, err := client.Channels.FindUnique(
		db.Channels.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("channel not found")
	}
	return channel, nil
}
