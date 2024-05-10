package languageService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.LanguagesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	language, _ := Get(input.Language)

	if language != nil {
		return nil, errors.New("language is already registerd")
	}

	client, ctx, err := database.GetChannelClient()
	if err != nil {
		return nil, err
	}
	createdLanguage, err := client.Languages.CreateOne(
		db.Languages.Language.Set(input.Language),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdLanguage, nil
}
