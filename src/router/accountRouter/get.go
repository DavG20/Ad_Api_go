package accountRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/accountService"

	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Account, error) {
	accountData, err := accountService.Get(id)

	if err != nil || accountData == nil {
		return nil, nil
	}

	account := utils.AccountGraphqlConverter(accountData)

	return account, nil

}
