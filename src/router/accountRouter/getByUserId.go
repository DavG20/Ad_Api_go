package accountRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/utils"
	"context"
	"errors"
	"log"
)

func GetByUserId(ctx context.Context) (*model.Account, error) {
	userId := ""

	accountData, _ := accountService.GetByUserID(userId)

	if accountData == nil {
		log.Println("invalid userId ")
		return nil, errors.New("account not found")
	}

	account := utils.AccountGraphqlConverter(accountData)

	return account, nil

}
