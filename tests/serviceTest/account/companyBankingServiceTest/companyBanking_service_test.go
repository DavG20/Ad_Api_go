package companyBankingServiceTest

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/service/bankService"
	"adtec/backend/src/service/companyBankingService"
	"adtec/backend/src/service/companyService"
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
	id          string
	bankId      string
	bankAccount string
	companyId   string

	bank *db.BanksModel

	company *db.CompaniesModel
)

func TestCreateCompanyBanking(t *testing.T) {

	bank, _ = bankService.Create(inputGenerator.GenerateRandomBankData())

	if bank == nil {
		bank, _ = bankService.Create(inputGenerator.GenerateRandomBankData())
	}
	company, _ = companyService.Create(inputGenerator.GenerateRandomCompanyData())

	if company == nil {
		company, _ = companyService.Create(inputGenerator.GenerateRandomCompanyData())
	}

	input := inputGenerator.GenerateRandomCompanyBankingData()

	input.CompanyID = company.ID

	input.BankId = bank.ID

	result, err := companyBankingService.Create(input)
	fmt.Println(err)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	id = result.ID
	companyId = result.CompanyID
	bankId = result.BankID
	bankAccount = result.BankAccount
	assert.Equal(t, input.CompanyID, result.CompanyID)
	assert.Equal(t, input.BankId, result.BankID)
}

func TestCreateCompanyBanking_InvalidBankID(t *testing.T) {

	// here the bankID is just a random string , so it is  invalid

	input := inputGenerator.GenerateRandomCompanyBankingData()

	// make the companyID valid , incase
	input.CompanyID = companyId

	result, err := companyBankingService.Create(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid bank id", "Unexpected error message")
}

func TestCreateCompanyBanking_InvalidCreatorID(t *testing.T) {

	// here the companyID is just a random string , so it will be invalid
	input := inputGenerator.GenerateRandomCompanyBankingData()

	input.BankId = bankId

	result, err := companyBankingService.Create(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid company id", "Unexpected error message")

}

func TestGetCompanyBanking(t *testing.T) {
	if id == "" {
		TestCreateCompanyBanking(t)
	}

	id := id

	result, err := companyBankingService.Get(id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
}

func TestGetCompanyBanking_InvalidID(t *testing.T) {
	id := "invalid_ID"
	result, err := companyBankingService.Get(id)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")
}

func TestGetByBankAccountAndId(t *testing.T) {

	if bankId == "" || bankAccount == "" {
		TestCreateCompanyBanking(t)
	}
	bankAct := bankAccount
	bankId := bankId

	result, err := companyBankingService.GetByBankAccountAndId(bankAct, bankId)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, bankAct, result.BankAccount)
	assert.Equal(t, bankId, result.BankID)
}

func TestGetByBankAccountAndId_InvalidBankID(t *testing.T) {
	bankId := "invalid_bankID"
	bancAct := bankAccount

	result, err := companyBankingService.GetByBankAccountAndId(bancAct, bankId)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")
}

func TestGetByBankAccountAndId_InvalidBankAccount(t *testing.T) {
	
	bankId := bankId
	bancAccount := "invalid_bankAccount"

	result, err := companyBankingService.GetByBankAccountAndId(bancAccount, bankId)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")
}
func TestGetAllCompanyBanking(t *testing.T) {
	companyID := companyId
	input := companyBankingService.FindManyInput{
		CompanyID: &companyID,
		BankId:    nil,
	}
	result, err := companyBankingService.GetAll(input)
	fmt.Println("result", result)
	assert.NoError(t, err)
	assert.NotNil(t, result)
}
