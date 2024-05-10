package accountBankingServiceTest

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/service/accountBankingService"
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/service/bankService"
	"adtec/backend/src/utils/inputGenerator"
	"adtec/backend/tests/serviceTest"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	serviceTest.SetUpEnv()
}

var (
	id            string
	bankID        string
	bankAccountNo string

	account *db.AccountsModel

	bank *db.BanksModel
)

func TestCreateAccountBanking(t *testing.T) {

	bank, _ = bankService.Create(inputGenerator.GenerateRandomBankData())

	if bank == nil {
		bank, _ = bankService.Create(inputGenerator.GenerateRandomBankData())
	}
	account, _ = accountService.Create(inputGenerator.GenerateRandomAccountData())

	if account == nil {
		account, _ = accountService.Create(inputGenerator.GenerateRandomAccountData())
	}

	input := inputGenerator.GenerateRandomAccountBankingData()

	input.AccountId = account.ID
	input.BankId = bank.ID

	result, err := accountBankingService.Create(input)
	fmt.Println(err)
	assert.NoError(t, err)
	assert.NotNil(t, result)

	id = result.ID
	bankID = result.BankID
	bankAccountNo = result.BankAccount

	assert.Equal(t, input.AccountId, result.AccountID)
	assert.Equal(t, input.BankId, result.BankID)
}

func TestCreateAccountBanking_InvalidBankID(t *testing.T) {
	accountId := account.ID
	bankId := "BankId"
	fullNameOnBank := "FullNameOnBank"
	bankAccount := "BankAccount"
	currency := "USD"

	input := accountBankingService.CreateInput{
		AccountId:      accountId,
		BankId:         bankId,
		FullNameOnBank: fullNameOnBank,
		BankAccount:    bankAccount,
		Currency:       currency,
	}

	result, err := accountBankingService.Create(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid bank id", "Unexpected error message")
}
func TestCreateAccountBanking_InvalidAccountID(t *testing.T) {
	accountId := "invalid_accountID"
	bankId := bank.ID
	fullNameOnBank := "FullNameOnBank"
	bankAccount := "BankAccount"
	currency := "USD"

	input := accountBankingService.CreateInput{
		AccountId:      accountId,
		BankId:         bankId,
		FullNameOnBank: fullNameOnBank,
		BankAccount:    bankAccount,
		Currency:       currency,
	}

	result, err := accountBankingService.Create(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid account id", "Unexpected error message")
}
func TestCreateAccountBanking_InvalidCurrency(t *testing.T) {
	if account == nil || bank == nil {
		TestCreateAccountBanking(t)
	}
	accountId := account.ID
	bankId := bank.ID
	fullNameOnBank := "FullNameOnBank"
	bankAccount := "BankAccount"
	currency := "invalid_currency"

	input := accountBankingService.CreateInput{
		AccountId:      accountId,
		BankId:         bankId,
		FullNameOnBank: fullNameOnBank,
		BankAccount:    bankAccount,
		Currency:       currency,
	}

	result, err := accountBankingService.Create(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The currency should be a valid ISO 4217 currency code", "Unexpected error message")
}

func TestGetAccountBanking(t *testing.T) {

	if id == "" {
		TestCreateAccountBanking(t)
	}
	id := id

	result, err := accountBankingService.Get(id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)

}

func TestGetAccountBanking_InvalidID(t *testing.T) {
	id := "invalid_id"
	result, err := accountBankingService.Get(id)

	// Add assertions based on the specific error scenario
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")
}

func TestGetAllAccountBanking(t *testing.T) {

	input := accountBankingService.FindManyInput{
		AccountId: nil,
		BankId:    nil,
	}

	result, err := accountBankingService.GetAll(input)
	fmt.Println("result", result)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	// Add more assertions based on your specific test case
}

func TestGetByBankAccountId(t *testing.T) {

	if bankID == "" || bankAccountNo == "" {
		TestCreateAccountBanking(t)
	}
	bankId := bankID
	bankAccount := bankAccountNo
	result, err := accountBankingService.GetByBankAccountAndId(bankAccount, bankId)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, bankAccount, result.BankAccount)
	assert.Equal(t, bankId, result.BankID)

}

func TestGetByBankAccountId_InvalidInput(t *testing.T) {

	bankId := "invalid_bankId"
	bancAccount := "invalid_BankAccount"

	result, err := accountBankingService.GetByBankAccountAndId(bancAccount, bankId)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")

}

func TestDeleteAccountBanking(t *testing.T) {

	if id == "" {
		TestCreateAccountBanking(t)
	}
	id := id

	result, err := accountBankingService.Delete(id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.Equal(t, "deleted Successfully", result.Message)
}
func TestDeleteAccountBanking_InvalidID(t *testing.T) {

	id := "invalid_id"

	result, err := accountBankingService.Delete(id)

	assert.Error(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.EqualError(t, err, "AccountBanking not Found", result.Message)
}
