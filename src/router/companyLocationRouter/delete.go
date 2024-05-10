package companyLocationRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/companyLocationService"
	"context"
)

func Delete(ctx context.Context, companyId, locationId string) (*model.DeletionResult, error) {
	return companyLocationService.Delete(companyId, locationId)
}
