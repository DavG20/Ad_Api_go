package languageRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/languageService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateLanguageInput) (*model.Languages, error) {
	createdLanguage, err := languageService.Create(
		languageService.CreateInput{
			Language: input.Language,
		},
	)
	if err != nil {
		return nil, err
	}

	language := utils.LanguageGraphqlConverter(createdLanguage)
	return language, nil
}