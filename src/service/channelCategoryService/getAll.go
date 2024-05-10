package channelCategoryService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
)

func GetAll(filter FindManyInput) ([]db.ChannelCategoriesModel, error) {

	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	parameter := []db.ChannelCategoriesWhereParam{}

	if filter.ChannelId != nil {
		parameter = append(parameter, db.ChannelCategories.ChannelID.Equals(*filter.ChannelId))
	}
	if filter.Category != nil {
		parameter = append(parameter, db.ChannelCategories.Or(
			db.ChannelCategories.Category.Contains(*filter.Category),
			db.ChannelCategories.Category.Mode(db.QueryModeInsensitive),
		))
	}

	fetchChannelCategory := client.ChannelCategories.FindMany(parameter[:]...).Take(int(limit))
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchChannelCategory = client.ChannelCategories.FindMany(parameter[:]...).Skip(1).Take(int(limit)).Cursor(
				db.ChannelCategories.Category.Cursor(string(*filter.Filter.After)),
			)
		}
		if filter.Filter.Before != nil {
			fetchChannelCategory = client.ChannelCategories.FindMany(parameter[:]...).Skip(1).Take(int(limit)).Cursor(
				db.ChannelCategories.Category.Cursor(string(*filter.Filter.Before)),
			)
		}
	}
	channelCategorys, err := fetchChannelCategory.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return channelCategorys, nil

}
