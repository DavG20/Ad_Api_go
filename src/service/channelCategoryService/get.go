package channelCategoryService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
)

func Get(input ChannelCategoryInput) (*db.ChannelCategoriesModel, error) {
	//TODO: Should I add validation here?
	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	channelCategory, err := client.ChannelCategories.FindUnique(
		db.ChannelCategories.ChannelCategoryID(
			db.ChannelCategories.ChannelID.Equals(input.ChannelID),
			db.ChannelCategories.Category.Equals(input.Category),
		)).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return channelCategory, nil
}
