package accountBankingService

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Delete(id string) (*model.DeletionResult, error) {

	client, ctx, err := database.GetAccountClient()

	deletionResult := &model.DeletionResult{
		Success: true,
		Message: "deleted Successfully",
	}

	if err != nil {
		deletionResult.Message = "internal error of db"
		deletionResult.Success = false
		return deletionResult, err
	}
	accountBanking, _ := Get(id)

	if accountBanking == nil {
		deletionResult.Success = false
		deletionResult.Message = "AccountBanking not Found"
		return deletionResult, errors.New("AccountBanking not Found")
	}

	_, err = client.AccountBankings.FindUnique(
		db.AccountBankings.ID.Equals(id),
	).Delete().Exec(ctx)

	if err != nil {
		deletionResult.Success = false
		deletionResult.Message = "AccountBanking not Found"
		return deletionResult, err
	}

	return deletionResult, nil
}
