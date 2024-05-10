package accountRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service"
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/utils"
	"context"
	"time"
)

func Create(ctx context.Context, input model.CreateAccountInput) (*model.Account, error) {
	var birthDate *time.Time
	if input.BirthDate != nil {
		date, _ := time.Parse("2006-01-02", *input.BirthDate)
		birthDate = &date
	}
	accountData, err := accountService.Create(
		accountService.CreateInput{
			UserID:      input.UserID,
			UserName:    input.UserName,
			Email:       input.Email,
			PhoneNumber: input.PhoneNumber,
			FullName:    input.FullName,
			BirthDate:   birthDate,
			AccountType: (*service.AccountType)(input.AccountType),
		})
	if err != nil {
		return nil, err
	}

	account := utils.AccountGraphqlConverter(accountData)

	return account, nil

}
