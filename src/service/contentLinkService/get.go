package contentLinkService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Get(contentId string) (*db.ContentLinksModel, error) {
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	contentLink, err := client.ContentLinks.FindUnique(
		db.ContentLinks.ContentID.Equals(contentId),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("contentLink not found")
	}
	return contentLink, nil
}
