package locationServiceTest

import (
	"adtec/backend/src/service/locationService"
	"adtec/backend/src/utils/inputGenerator"
	"adtec/backend/tests/serviceTest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	serviceTest.SetUpEnv()
}

var id string

func TestCreateLocation(t *testing.T) {

	input := inputGenerator.GenerateRandomLocationData()

	result, err := locationService.Create(input)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	id = result.ID
	assert.NotEmpty(t, result.ID)
	assert.Equal(t, input.Country, result.Country)
	assert.Equal(t, input.State, result.State)
}
func TestGetLocation(t *testing.T) {

	if id == "" {
		TestCreateLocation(t)
	}
	iD := id

	result, err := locationService.Get(iD)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
}

func TestGetLocation_InvalidID(t *testing.T) {
	id := "invalid"

	result, err := locationService.Get(id)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "location not found", "unexpected error message")
}

func TestUpdateLocation(t *testing.T) {
	if id == "" {
		TestCreateLocation(t)
	}
	id := id
	createInput := inputGenerator.GenerateRandomLocationData()

	input := locationService.UpdateInput{
		ID:         id,
		State:      &createInput.State,
		PostalCode: createInput.PostalCode,
	}

	result, err := locationService.Update(input)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
}

func TestUpdateLocation_InvalidID(t *testing.T) {
	id := "invalid"
	newState := "California"
	newPostalCode := "90001"

	input := locationService.UpdateInput{
		ID:         id,
		State:      &newState,
		PostalCode: &newPostalCode,
	}

	result, err := locationService.Update(input)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "location not found", "unexpected error message")
}
