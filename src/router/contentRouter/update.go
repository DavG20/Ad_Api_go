package contentRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/contentService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateContentInput) (*model.Contents, error) {

	updatedData, err := contentService.Update(contentService.UpdateInput{
		ID:          input.ID,
		ContentType: (*service.ContentType)(input.ContentType),
		Description: input.Description,
	})
	if err != nil {
		return nil, err
	}

	updatedContent := utils.ContentGraphqlConverter(updatedData)
	return updatedContent, nil
}
