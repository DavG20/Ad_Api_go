package languageRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/languageService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Languages, error) {
	languageData, err := languageService.Get(id)
	if err != nil {
		return nil, err
	}

	language := utils.LanguageGraphqlConverter(languageData)

	return language, err
}
