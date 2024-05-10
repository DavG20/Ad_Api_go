package channelCategoryService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/service/categoryService"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input ChannelCategoryInput) (*db.ChannelCategoriesModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	category, _ := categoryService.Get(input.Category)
	if category == nil {
		return nil, errors.New("category doesn't exist")
	}

	channel, _ := channelService.Get(input.ChannelID)
	if channel == nil {
		return nil, errors.New("channel doesn't exist")
	}

	categoryCategory, _ := Get(input)

	if categoryCategory != nil {
		return nil, errors.New("category and channel id is already registered")
	}

	client, ctx, err := database.GetChannelClient()
	if err != nil {
		return nil, err
	}
	createdCategory, err := client.ChannelCategories.CreateOne(
		db.ChannelCategories.CategoryData.Link(db.Categories.Category.Equals(input.Category)),
		db.ChannelCategories.Channel.Link(db.Channels.ID.Equals(input.ChannelID)),
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdCategory, nil
}
