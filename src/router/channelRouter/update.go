package channelRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/utils"
	"context"
)

func Update(ctx context.Context, input model.UpdateChannelInput) (*model.Channels, error) {
	updatedData, err := channelService.Update(channelService.UpdateInput{
		ID:            input.ID,
		OwnerID:       input.OwnerID,
		UserName:      input.UserName,
		Name:          input.Name,
		Description:   input.Description,
		BotAddAsAdmin: input.BotAddAsAdmin,
		Country:       input.Country,
		Currency:      input.Currency,
	})
	if err != nil {
		return nil, err
	}

	updatedChannel := utils.ChannelGraphqlConverter(updatedData)
	return updatedChannel, nil
}
