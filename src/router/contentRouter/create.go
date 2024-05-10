package contentRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/contentService"
	"adtec/backend/src/utils"
	"context"
)

func Create(ctx context.Context, input model.CreateContentInput) (*model.Contents, error) {

	contentData, err := contentService.Create(
		contentService.CreateInput{
			CampaignID:  input.CampaignID,
			ContentType: (*service.ContentType)(input.ContentType),
			Description: input.Description,
		},
	)
	if err != nil {
		return nil, err
	}

	createdContent := utils.ContentGraphqlConverter(contentData)

	return createdContent, nil
}
