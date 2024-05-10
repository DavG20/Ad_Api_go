package bankServiceTest

import (
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

var id string

func TestCreateBank(t *testing.T) {

	input := inputGenerator.GenerateRandomBankData()

	result, err := bankService.Create(input)
	fmt.Println(err)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	id = result.ID
	assert.Equal(t, input.BankName, result.BankName)
	assert.Equal(t, input.BankCode, result.BankCode)
}

func TestCreateBank_InvalidBankCode(t *testing.T) {

	bankCode := ""
	bankName := "Zemen"

	input := bankService.CreateInput{
		BankName: bankName,
		BankCode: bankCode,
	}

	result, err := bankService.Create(input)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The bankCode is required", "Unexpected error message")
}

func TestCreateBank_InvalidBankName(t *testing.T) {

	bankCode := "10"
	bankName := ""

	input := bankService.CreateInput{
		BankName: bankName,
		BankCode: bankCode,
	}

	result, err := bankService.Create(input)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The bankName is required", "Unexpected error message")
}

func TestUpdateBank(t *testing.T) {
	if id == "" {
		TestCreateBank(t)
	}
	bankId := id

	input := inputGenerator.GenerateRandomBankData()

	updateInput := bankService.UpdateInput{
		ID:       bankId,
		BankName: &input.BankName,
	}

	result, err := bankService.Update(updateInput)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, input.BankName, result.BankName)
	assert.Equal(t, bankId, result.ID)
}
func TestUpdateBank_InvalidID(t *testing.T) {
	bankId := "invalidBankID"
	bankCode := "03"
	bankName := "ABAY"

	input := bankService.UpdateInput{
		ID:       bankId,
		BankName: &bankName,
		BankCode: &bankCode,
	}

	result, err := bankService.Update(input)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid bank ID", "Unexpected error message")
}

func TestGetBank(t *testing.T) {
	if id == "" {
		TestCreateBank(t)
	}
	bankId := id

	result, err := bankService.Get(bankId)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, bankId, result.ID)
}

func TestGetBank_InvalidID(t *testing.T) {
	bankId := "invalidBankID"

	result, err := bankService.Get(bankId)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")
}

func TestGetAllBank(t *testing.T) {
	input := bankService.FindManyInput{}
	result, err := bankService.GetAll(input)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
