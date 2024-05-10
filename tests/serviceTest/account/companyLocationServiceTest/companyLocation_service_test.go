package companyLocationServiceTest

import (
	"adtec/backend/src/prisma/account/db"
	"adtec/backend/src/service/companyLocationService"
	"adtec/backend/src/service/companyService"
	"adtec/backend/src/service/locationService"
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
	companyId  string
	locationId string

	location *db.LocationsModel

	company *db.CompaniesModel
)

func TestCreateCompanyLocation(t *testing.T) {

	company, _ = companyService.Create(inputGenerator.GenerateRandomCompanyData())

	if company == nil {

		company, _ = companyService.Create(inputGenerator.GenerateRandomCompanyData())
	}

	location, _ = locationService.Create(inputGenerator.GenerateRandomLocationData())

	if location == nil {

		location, _ = locationService.Create(inputGenerator.GenerateRandomLocationData())
	}

	input := inputGenerator.GenerateRandomCompanyLocationData()

	input.CompanyID = company.ID
	input.LocationID = location.ID

	result, err := companyLocationService.Create(input)
	fmt.Println(err)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	companyId = result.CompanyID
	locationId = result.LocationID
	assert.Equal(t, input.CompanyID, result.CompanyID)
	assert.Equal(t, input.LocationID, result.LocationID)
}

func TestCreateCompanyLocation_InvalidCompanyID(t *testing.T) {

	input := inputGenerator.GenerateRandomCompanyLocationData()

	// change random locationId and left the companyId
	input.LocationID = location.ID

	result, err := companyLocationService.Create(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "company not found", "unexpected error message")

}
func TestCreateCompanyLocation_InvalidLocationID(t *testing.T) {

	input := inputGenerator.GenerateRandomCompanyLocationData()

	// change random companyId and left the locationId
	input.CompanyID = company.ID

	result, err := companyLocationService.Create(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "location not found", "unexpected error message")

}

func TestGetCompanyLocation(t *testing.T) {

	if locationId == "" || companyId == "" {
		TestCreateCompanyLocation(t)
	}

	result, err := companyLocationService.Get(companyId, locationId)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, companyId, result.CompanyID)
}
func TestGetCompanyLocation_InvalidCompanyId(t *testing.T) {

	result, err := companyLocationService.Get("invalid", locationId)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "unexpected error message")
}
func TestGetCompanyLocation_InvalidLocationId(t *testing.T) {

	result, err := companyLocationService.Get(companyId, "invalid")

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "unexpected error message")
}
func TestDeleteCompanyLocation(t *testing.T) {

	result, err := companyLocationService.Delete(companyId, locationId)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.Equal(t, "deleted Successfully", result.Message)
}
func TestDeleteCompanyLocation_InvalidCompanyID(t *testing.T) {

	result, err := companyLocationService.Delete("invalid", locationId)

	assert.Error(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.EqualError(t, err, "companyLocation not Found", "unexpected error message")
}
func TestDeleteCompanyLocation_InvalidLocationID(t *testing.T) {

	result, err := companyLocationService.Delete(companyId, "invalid")

	assert.Error(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.EqualError(t, err, "companyLocation not Found", "unexpected error message")
}
