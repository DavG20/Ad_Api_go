package contentLinkService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/service/contentService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.ContentLinksModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	content, err := contentService.Get(input.ContentID)
	if err != nil || content == nil {
		return nil, errors.New("content  not found or invalid content id")
	}

	contentLink, _ := Get(input.ContentID)
	if contentLink != nil {
		return nil, errors.New("contentId is Taken")
	}
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.ContentLinksSetParam{}

	if input.Link != nil {
		parameter = append(parameter, db.ContentLinks.Link.Set(*input.Link))
	}
	if input.Title != nil {
		parameter = append(parameter, db.ContentLinks.Title.Set(*input.Title))
	}

	createdContentLink, err := client.ContentLinks.CreateOne(
		db.ContentLinks.Content.Link(db.Contents.ID.Equals(input.ContentID)),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return createdContentLink, nil
}
