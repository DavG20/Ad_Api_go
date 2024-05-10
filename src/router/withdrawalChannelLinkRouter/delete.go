package withdrawalChannelLinkRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/withdrawalChannelLinkService"
	"adtec/backend/src/utils/convertor"
	"context"
)

func Delete(ctx context.Context, input *model.WithdrawalChannelLinkInput) (*model.DeletionResult, error) {

	deletionResult, err := withdrawalChannelLinkService.Delete(convertor.WithdrawalChannelLinkInputConverter(input))

	return deletionResult, err

}
