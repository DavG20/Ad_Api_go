package channelRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/channelService"
	"context"
)

func Delete(ctx context.Context, id string) (*model.DeletionResult, error) {

	deletionResult, err := channelService.Delete(id)

	return deletionResult, err

}
