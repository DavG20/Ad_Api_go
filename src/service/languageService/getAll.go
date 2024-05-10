package languageService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
)

func GetAll(filter FindManyInput) ([]db.LanguagesModel, error) {

	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}


	

	fetchLanguage := client.Languages.FindMany().Take(int(limit))
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchLanguage = client.Languages.FindMany().Skip(1).Take(int(limit)).Cursor(
				db.Languages.Language.Cursor(string(*filter.Filter.After)),
			)
		}
		if filter.Filter.Before != nil {
			fetchLanguage = client.Languages.FindMany().Skip(1).Take(int(limit)).Cursor(
				db.Languages.Language.Cursor(string(*filter.Filter.Before)),
			)
		}
	}
	languages, err := fetchLanguage.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return languages, nil

}
