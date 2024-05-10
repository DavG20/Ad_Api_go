package contentLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/contentLinkService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateContentLinkInput) (*model.ContentLinks, error) {

	contentData, err := contentLinkService.Create(
		contentLinkService.CreateInput{
			ContentID: input.ContentID,
			Link:      input.Link,
			Title:     input.Title,
		},
	)
	if err != nil {
		return nil, err
	}

	createdContentLink := utils.ContentLinkGraphqlConverter(contentData)

	return createdContentLink, nil
}
