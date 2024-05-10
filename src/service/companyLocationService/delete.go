package companyLocationService

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/utils/database"
	"errors"
)

func Delete(companyId, locationId string) (*model.DeletionResult, error) {
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

	companyLocation, _ := Get(companyId, locationId)
	if companyLocation == nil {
		deletionResult.Success = false
		deletionResult.Message = "companyLocation not Found"
		return deletionResult, errors.New("companyLocation not Found")
	}

	_, err = client.CompanyLocations.FindUnique(
		db.CompanyLocations.UniqueCompanyLocationID(
			db.CompanyLocations.CompanyID.Equals(companyId),
			db.CompanyLocations.LocationID.Equals(locationId),
		),
	).Delete().Exec(ctx)

	if err != nil {
		deletionResult.Success = false
		deletionResult.Message = "error deleting"
		return deletionResult, err
	}

	return deletionResult, nil
}
