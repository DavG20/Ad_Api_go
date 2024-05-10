package accountBankingRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/accountBankingService"
	"context"
)

func Delete(ctx context.Context, id string) (*model.DeletionResult, error) {

	deletionResult, err := accountBankingService.Delete(id)

	return deletionResult, err

}
