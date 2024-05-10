package contentService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/service"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.ContentsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	content, _ := Get(input.ID)

	if content == nil {
		return nil, errors.New("invalid Content id or Content doesn't exist")
	}

	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.ContentsSetParam{}

	if input.ContentType != nil {
		contentType, _ := content.ContentType()
		if *input.ContentType != service.ContentType(contentType) {
			parameter = append(parameter, db.Contents.ContentType.Set(db.ContentType(*input.ContentType)))

		}
	}
	if input.Description != nil {
		description, _ := content.Description()
		if *input.Description != description {
			parameter = append(parameter, db.Contents.Description.Set(*input.Description))

		}
	}

	updateContent, err := client.Contents.FindUnique(
		db.Contents.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateContent, nil
}
