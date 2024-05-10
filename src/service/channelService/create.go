package channelService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Create(input CreateInput) (*db.ChannelsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	channelByUserName, _ := GetChannelByUserName(input.UserName)

	if channelByUserName != nil {
		return nil, errors.New("userName Taken")
	}
	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.ChannelsSetParam{}

	if input.Name != nil {
		parameter = append(parameter, db.Channels.Name.Set(*input.Name))
	}
	if input.Description != nil {
		parameter = append(parameter, db.Channels.Description.Set(*input.Description))
	}
	if input.BotAddAsAdmin != nil {
		parameter = append(parameter, db.Channels.BotAddAsAdmin.Set(*input.BotAddAsAdmin))
	}
	if input.Country != nil {
		parameter = append(parameter, db.Channels.Country.Set(*input.Country))
	}
	if input.ChannelCreatedAt != nil {
		parameter = append(parameter, db.Channels.ChannelCreatedAt.Set(*input.ChannelCreatedAt))
	}
	if input.Currency != nil {
		parameter = append(parameter, db.Channels.Currency.Set(*input.Currency))
	}

	createdChannel, err := client.Channels.CreateOne(
		db.Channels.OwnerID.Set(input.OwnerID),
		db.Channels.UserName.Set(input.UserName),
		parameter[:]...,
	).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return createdChannel, nil
}
