package channelContentLanguageRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelContentLanguageService"
	"adtec/backend/src/utils/convertor"
	"context"
)

func Delete(ctx context.Context, input *model.ChannelContentLanguageInput) (*model.DeletionResult, error) {

	deletionResult, err := channelContentLanguageService.Delete(convertor.ChannelContentLanguageInputConverter(input))

	return deletionResult, err

}
