package categoryService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
)

func Get(categoryId string) (*db.CategoriesModel, error) {
	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	category, err := client.Categories.FindUnique(
		db.Categories.Category.Equals(categoryId),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}
	return category, nil
}
