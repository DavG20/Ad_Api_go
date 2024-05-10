package contentLinkService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.ContentLinksModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	contentLink, _ := Get(input.ContentID)

	if contentLink == nil {
		return nil, errors.New("invalid ContentLink id or ContentLink doesn't exist")
	}

	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.ContentLinksSetParam{}

	if input.Link != nil {
		link, _ := contentLink.Link()
		if *input.Link != link {
			parameter = append(parameter, db.ContentLinks.Link.Set(*input.Link))

		}
	}
	if input.Title != nil {
		title, _ := contentLink.Title()
		if *input.Title != title {
			parameter = append(parameter, db.ContentLinks.Title.Set(*input.Title))

		}
	}

	updateContentLink, err := client.ContentLinks.FindUnique(
		db.ContentLinks.ContentID.Equals(input.ContentID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateContentLink, nil
}
