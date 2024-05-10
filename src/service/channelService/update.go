package channelService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/validator"
	"errors"
)

func Update(input UpdateInput) (*db.ChannelsModel, error) {

	if ok, validationErrors := validator.ValidateInputs(input); !ok {
		return nil, errors.New(validationErrors[0].Message)
	}

	channel, _ := Get(input.ID)

	if channel == nil {
		return nil, errors.New("invalid channel id or channel doesn't exist")
	}

	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	parameter := []db.ChannelsSetParam{}

	if input.UserName != nil {
		userNameExist, _ := GetChannelByUserName(*input.UserName)

		if userNameExist != nil {
			return nil, errors.New("username is taken")
		}
		if *input.UserName != channel.UserName {
			parameter = append(parameter, db.Channels.UserName.Set(*input.UserName))
		}

	}
	// have to check the existence of ownerID
	if input.OwnerID != nil {
		if *input.OwnerID != channel.OwnerID {
			parameter = append(parameter, db.Channels.OwnerID.Set(*input.OwnerID))

		}
	}
	if input.Name != nil {
		name, _ := channel.Name()
		if *input.Name != name {
			parameter = append(parameter, db.Channels.Name.Set(*input.Name))

		}
	}
	if input.Description != nil {
		description, _ := channel.Description()
		if *input.Description != description {
			parameter = append(parameter, db.Channels.Description.Set(*input.Description))

		}
	}
	if input.BotAddAsAdmin != nil {
		if *input.BotAddAsAdmin != channel.BotAddAsAdmin {
			parameter = append(parameter, db.Channels.BotAddAsAdmin.Set(*input.BotAddAsAdmin))
		}
	}
	if input.Country != nil {
		country, _ := channel.Country()
		if *input.Country != country {
			parameter = append(parameter, db.Channels.Country.Set(*input.Country))

		}
	}
	if input.Currency != nil {

		currency, _ := channel.Currency()
		if *input.Currency != currency {
			parameter = append(parameter, db.Channels.Currency.Set(*input.Currency))

		}
	}

	updateChannel, err := client.Channels.FindUnique(
		db.Channels.ID.Equals(input.ID),
	).Update(parameter[:]...).Exec(ctx)

	if err != nil {
		return nil, err
	}

	return updateChannel, nil
}
