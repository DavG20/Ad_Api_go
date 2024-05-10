package companyBankingService

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
		deletionResult.Message = "CompanyBanking not Found"
		return deletionResult, errors.New("CompanyBanking not Found")
	}

	_, err = client.CompanyBankings.FindUnique(
		db.CompanyBankings.ID.Equals(id),
	).Delete().Exec(ctx)

	if err != nil {
		return nil, err
	}

	return deletionResult, nil
}
