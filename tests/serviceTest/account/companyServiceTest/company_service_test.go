package companyServiceTest

import (
	"adtec/backend/src/service/companyService"
	"adtec/backend/src/utils/inputGenerator"
	"adtec/backend/tests/serviceTest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	serviceTest.SetUpEnv()
}

var (
	id         string
	uniqueName string
)

func TestCreateCompany(t *testing.T) {

	input := inputGenerator.GenerateRandomCompanyData()

	result, err := companyService.Create(input)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	id = result.ID
	uniqueName = result.UniqueName
	assert.Equal(t, input.Name, result.Name)
	assert.Equal(t, input.UniqueName, result.UniqueName)

}

func TestCreateCompany_InvalidInput(t *testing.T) {
	input := inputGenerator.GenerateRandomCompanyData()

	// set creatorId to empty string to test validation

	input.CreatorID = ""

	result, err := companyService.Create(input)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The creatorId is required", "unexpected error message")

}

func TestUpdateCompany(t *testing.T) {
	if id == "" {

		TestCreateCompany(t)
	}
	iD := id

	createInputGen := inputGenerator.GenerateRandomCompanyData()

	input := companyService.UpdateInput{
		ID:         iD,
		Name:       &createInputGen.Name,
		UniqueName: &createInputGen.UniqueName,
		TinNumber:  createInputGen.TinNumber,
	}

	result, err := companyService.Update(input)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	id = result.ID
	uniqueName = result.UniqueName
	assert.Equal(t, *input.Name, result.Name)
	assert.Equal(t, *input.UniqueName, result.UniqueName)
}

func TestUpdateCompany_InvalidInput(t *testing.T) {

	id := "invalid_id"

	companyName := "ABAY"

	input := companyService.UpdateInput{
		ID:   id,
		Name: &companyName,
	}
	result, err := companyService.Update(input)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid company id or company doesn't exist", "unexpected error message")
}

func TestGetCompany(t *testing.T) {
	if id == "" {
		TestCreateCompany(t)
	}
	iD := id
	result, err := companyService.Get(iD)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, iD, result.ID)
}

func TestGetCompany_InvalidInput(t *testing.T) {

	id := "invalid_id"

	result, err := companyService.Get(id)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "company not found", "unexpected error message")
}

func TestGetByUniqueName(t *testing.T) {
	if uniqueName == "" {
		TestCreateCompany(t)
	}

	result, err := companyService.GetByUniqueName(uniqueName)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, uniqueName, result.UniqueName)
}
