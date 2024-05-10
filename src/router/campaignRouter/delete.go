package campaignRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/campaignService"
	"context"
)

func Delete(ctx context.Context, id string) (*model.DeletionResult, error) {

	deletionResult, err := campaignService.Delete(id)

	return deletionResult, err

}
