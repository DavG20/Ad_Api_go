package channelContentLanguageService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
)

func GetAll(filter FindManyInput) ([]db.ChannelContentLanguagesModel, error) {

	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}
	parameter := []db.ChannelContentLanguagesWhereParam{}

	if filter.ChannelId != nil {
		parameter = append(parameter, db.ChannelContentLanguages.ChannelID.Equals(*filter.ChannelId))
	}
	if filter.Language != nil {
		parameter = append(parameter, db.ChannelContentLanguages.Or(
			db.ChannelContentLanguages.Language.Contains(*filter.Language),
			db.ChannelContentLanguages.Language.Mode(db.QueryModeInsensitive),
		))
	}

	fetchContentLanguage := client.ChannelContentLanguages.FindMany(parameter[:]...).Take(int(limit))
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchContentLanguage = client.ChannelContentLanguages.FindMany(parameter[:]...).Skip(1).Take(int(limit)).Cursor(
				db.ChannelContentLanguages.Language.Cursor(string(*filter.Filter.After)),
			)
		}
		if filter.Filter.Before != nil {
			fetchContentLanguage = client.ChannelContentLanguages.FindMany(parameter[:]...).Skip(1).Take(int(limit)).Cursor(
				db.ChannelContentLanguages.Language.Cursor(string(*filter.Filter.Before)),
			)
		}
	}
	contentLanguages, err := fetchContentLanguage.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return contentLanguages, nil

}
