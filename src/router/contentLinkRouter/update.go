package contentLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/contentLinkService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateContentLinkInput) (*model.ContentLinks, error) {

	updatedData, err := contentLinkService.Update(contentLinkService.UpdateInput{
		ContentID: input.ContentID,
		Link:      input.Link,
		Title:     input.Title,
	})
	if err != nil {
		return nil, err
	}

	updatedContentLink := utils.ContentLinkGraphqlConverter(updatedData)
	return updatedContentLink, nil
}
