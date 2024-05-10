package categoryService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
)

func GetAll(filter FindManyInput) ([]db.CategoriesModel, error) {

	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	fetchCategory := client.Categories.FindMany().Take(int(limit))
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchCategory = client.Categories.FindMany().Skip(1).Take(int(limit)).Cursor(
				db.Categories.Category.Cursor(string(*filter.Filter.After)),
			)
		}
		if filter.Filter.Before != nil {
			fetchCategory = client.Categories.FindMany().Skip(1).Take(int(limit)).Cursor(
				db.Categories.Category.Cursor(string(*filter.Filter.Before)),
			)
		}
	}
	categories, err := fetchCategory.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil

}
