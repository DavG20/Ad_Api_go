package channelService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
)

func GetAll(filter ChannelFilter) ([]db.ChannelsModel, error) {

	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	parameter := []db.ChannelsWhereParam{}

	if filter.Name != nil {
		parameter = append(parameter, db.Channels.Or(
			db.Channels.Name.Contains(*filter.Name),
			db.Channels.Name.Mode(db.QueryModeInsensitive),
		))
	}
	if filter.Country != nil {
		parameter = append(parameter, db.Channels.Or(
			db.Channels.Country.Equals(*filter.Country),
			db.Channels.Country.Mode(db.QueryModeInsensitive),
		))
	}
	if filter.Languages != nil {
		parameter = append(parameter, db.Channels.ChannelContentLanguages.Some(
			db.ChannelContentLanguages.Language.In(filter.Languages),
		))
	}
	if filter.Categories != nil {
		parameter = append(parameter, db.Channels.ChannelCategories.Some(
			db.ChannelCategories.Category.In(filter.Categories),
		))
	}

	if filter.CpmFilter != nil {
		if filter.CpmFilter.Max != nil {
			parameter = append(parameter, db.Channels.CpmRates.Where(db.CpmRates.Cpm.Lte(*filter.CpmFilter.Max)))
		}
		parameter = append(parameter, db.Channels.CpmRates.Where(db.CpmRates.Cpm.Gte(filter.CpmFilter.Min)))

	}
	if filter.SubCount != nil {
		if filter.SubCount.Max != nil {
			parameter = append(parameter, db.Channels.ChannelDetails.Where(db.ChannelDetails.SubCount.Lte(int(*filter.SubCount.Max))))
		}
		parameter = append(parameter, db.Channels.ChannelDetails.Where(db.ChannelDetails.SubCount.Gte(int(filter.SubCount.Min))))

	}
	if filter.PostViewFilter != nil {
		if filter.PostViewFilter.Max != nil {
			parameter = append(parameter, db.Channels.ChannelDetails.Where(db.ChannelDetails.AveragePostView.Lte(int(*filter.PostViewFilter.Max))))
		}
		parameter = append(parameter, db.Channels.ChannelDetails.Where(db.ChannelDetails.AveragePostView.Gte(int(filter.PostViewFilter.Min))))

	}

	fetchChannels := client.Channels.FindMany(
		parameter[:]...,
	).Take(int(limit)).OrderBy(
		db.Channels.CreatedAt.Order(db.DESC),
	)
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchChannels = client.Channels.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.Channels.ID.Cursor(string(*filter.Filter.After)),
			).OrderBy(
				db.Channels.CreatedAt.Order(db.DESC),
			)
		}
		if filter.Filter.Before != nil {
			fetchChannels = client.Channels.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.Channels.ID.Cursor(string(*filter.Filter.Before)),
			).OrderBy(
				db.Channels.CreatedAt.Order(db.DESC),
			)
		}
	}
	channel, err := fetchChannels.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return channel, nil

}
