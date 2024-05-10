package companyBankingRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/companyBankingService"
	"context"
)

func Delete(ctx context.Context, id string) (*model.DeletionResult, error) {

	deletionResult, err := companyBankingService.Delete(id)

	return deletionResult, err

}
