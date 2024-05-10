package contentService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Get(id string) (*db.ContentsModel, error) {
	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	content, err := client.Contents.FindUnique(
		db.Contents.ID.Equals(id),
	).Exec(ctx)

	if err != nil {
		return nil, errors.New("content not found")
	}
	return content, nil
}
