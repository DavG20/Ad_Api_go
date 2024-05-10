package companyMemberServiceTest

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/service/companyMemberService"
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
	id      string
	account *db.AccountsModel
	company *db.CompaniesModel
)

func TestCreateCompanyMember(t *testing.T) {

	account, _ = accountService.Create(inputGenerator.GenerateRandomAccountData())

	if account == nil {
		account, _ = accountService.Create(inputGenerator.GenerateRandomAccountData())
	}

	company, _ = companyService.Create(inputGenerator.GenerateRandomCompanyData())

	if company == nil {
		company, _ = companyService.Create(inputGenerator.GenerateRandomCompanyData())
	}

	input := inputGenerator.GenerateRandomCompanyMemberData()

	input.AccountID = account.ID
	input.CompanyID = company.ID

	result, err := companyMemberService.Create(input)
	fmt.Println(err)
	assert.NoError(t, err)
	id = result.ID
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, input.AccountID, result.AccountID)
	assert.Equal(t, input.CompanyID, result.CompanyID)
}

func TestCreateCompanyMember_InvalidAccountID(t *testing.T) {

	// use random account id to test validation

	input := inputGenerator.GenerateRandomCompanyMemberData()

	// change random company id

	input.CompanyID = company.ID

	result, err := companyMemberService.Create(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid account id", "unexpected error message")
}

func TestCreateCompanyMember_InvalidCompanyID(t *testing.T) {

	// use random company id to test validation
	input := inputGenerator.GenerateRandomCompanyMemberData()

	input.AccountID = account.ID

	result, err := companyMemberService.Create(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid company id", "unexpected error message")
}

func TestUpdateCompanyMember(t *testing.T) {
	if id == "" {
		TestCreateCompanyMember(t)
	}
	iD := id

	inputGen := inputGenerator.GenerateRandomCompanyMemberData()

	input := companyMemberService.UpdateInput{
		ID:   iD,
		Role: inputGen.Role,
	}

	result, err := companyMemberService.Update(input)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, iD, result.ID)
	assert.Equal(t, db.CompanyRole(*input.Role), result.Role)
}

func TestGetCompanyMember(t *testing.T) {
	if id == "" {
		TestCreateCompanyMember(t)
	}
	iD := id

	result, err := companyMemberService.Get(iD)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, iD, result.ID)
}

func TestGetCompanyMember_InvalidID(t *testing.T) {
	id := "invalid"

	result, err := companyMemberService.Get(id)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "unexpected error message")
}
