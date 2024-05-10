package languageRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/languageService"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/convertor"
	"context"
)

func GetAll(ctx context.Context, input *model.LanguageFilter) ([]*model.Languages, error) {

	languageData, err := languageService.GetAll(convertor.LanguageFilterConverter(input))

	if err != nil {
		return nil, err
	}

	languages := []*model.Languages{}

	for _, language := range languageData {
		language := utils.LanguageGraphqlConverter(language)
		languages = append(languages, language)
	}

	return languages, nil

}
