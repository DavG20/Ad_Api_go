package accountRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/utils"
	"context"
	"time"
)

func Update(ctx context.Context, input model.UpdateAccountInput) (*model.Account, error) {

	var birthDate *time.Time
	if input.BirthDate != nil {
		date, _ := time.Parse("2006-01-02", *input.BirthDate)
		birthDate = &date
	}

	updateAccount, err := accountService.Update(
		accountService.UpdateInput{
			ID:          input.ID,
			UserName:    input.UserName,
			Email:       input.Email,
			PhoneNumber: input.PhoneNumber,
			FullName:    input.FullName,
			BirthDate:   birthDate,
			AccountType: (*service.AccountType)(input.AccountType),
		},
	)
	if err != nil {
		return nil, err
	}

	account := utils.AccountGraphqlConverter(updateAccount)

	return account, nil
}
