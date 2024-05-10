package contentRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/contentService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Contents, error) {
	contentData, err := contentService.Get(id)
	if err != nil {
		return nil, err
	}

	content := utils.ContentGraphqlConverter(contentData)

	return content, nil
}
