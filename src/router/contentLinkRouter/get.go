package contentLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/contentLinkService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.ContentLinks, error) {
	contentLinkData, err := contentLinkService.Get(id)
	if err != nil {
		return nil, err
	}

	contentLink := utils.ContentLinkGraphqlConverter(contentLinkData)

	return contentLink, nil
}
