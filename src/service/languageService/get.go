package languageService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
)

func Get(languageId string) (*db.LanguagesModel, error) {
	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	language, err := client.Languages.FindUnique(
		db.Languages.Language.Equals(languageId),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return language, nil
}
