package categoryService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.CategoriesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	category, _ := Get(input.Category)

	if category != nil {
		return nil, errors.New("category is taken")
	}

	client, ctx, err := database.GetChannelClient()
	if err != nil {
		return nil, err
	}
	createdCategory, err := client.Categories.CreateOne(
		db.Categories.Category.Set(input.Category),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}
