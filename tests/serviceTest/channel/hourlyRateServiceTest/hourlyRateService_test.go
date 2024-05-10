package hourlyRateServiceTest

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/service/hourlyRateService"
	"adtec/backend/src/utils/inputGenerator"
	"adtec/backend/tests/serviceTest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	serviceTest.SetUpEnv()
}

var (
	hourlyRateID string

	channel *db.ChannelsModel
)

func TestCreateHourlyRate(t *testing.T) {

	channel, _ = channelService.Create(inputGenerator.GenerateRandomChannelData())

	if channel == nil {
		channel, _ = channelService.Create(inputGenerator.GenerateRandomChannelData())
	}
	// create an instance of our test object

	input := inputGenerator.GenerateRandomHourlyRateData()

	// make the chennelId a valid one

	input.ChannelID = channel.ID

	result, err := hourlyRateService.Create(input)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	hourlyRateID = result.ID
	assert.Equal(t, input.ChannelID, result.ChannelID)

	// add more assertions based on your specific test case
}

func TestCreateHourlyRate_InvalidInput(t *testing.T) {
	// create an instance of our test object
	input := inputGenerator.GenerateRandomHourlyRateData()

	// make the chennelId a valid one

	input.ChannelID = ""

	result, err := hourlyRateService.Create(input)

	// assert that the expectations were met
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The channelId is required", "Unexpected error message")

	// add more assertions based on your specific test case
}
func TestGetHourlyRate(t *testing.T) {
	// create an instance of our test object
	if hourlyRateID == "" {
		TestCreateHourlyRate(t)
	}
	id := hourlyRateID

	// Call the method we want to test
	result, err := hourlyRateService.Get(id)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
	// Add more assertions based on your specific test case
}

func TestGetHourlyRate_InvalidInput(t *testing.T) {
	// create an instance of our test object
	id := ""

	// Call the method we want to test
	result, err := hourlyRateService.Get(id)

	// assert that the expectations were met
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")
	// Add more assertions based on your specific test case
}

func TestUpdateHourlyRate(t *testing.T) {
	// create an instance of our test object
	if hourlyRateID == "" {
		TestCreateHourlyRate(t)
	}
	id := hourlyRateID

	hourlyRateData := inputGenerator.GenerateRandomHourlyRateData()

	input := hourlyRateService.UpdateInput{
		ID:              id,
		Active:          hourlyRateData.Active,
		HourlyRate:      hourlyRateData.HourlyRate,
		MinHourlyVolume: hourlyRateData.MinHourlyVolume,
		MaxHourlyVolume: hourlyRateData.MaxHourlyVolume,
	}

	result, err := hourlyRateService.Update(input)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
	// Add more assertions based on your specific test case
}

func TestUpdateHourlyRate_InvalidInput(t *testing.T) {
	// create an instance of our test object
	if hourlyRateID == "" {
		TestCreateHourlyRate(t)
	}
	id := hourlyRateID
	active := true
	hourlyRate := -10.0
	minHourlyVolume := 15.0

	input := hourlyRateService.UpdateInput{
		ID:              id,
		Active:          &active,
		HourlyRate:      &hourlyRate,
		MinHourlyVolume: &minHourlyVolume,
	}

	result, err := hourlyRateService.Update(input)

	// assert that the expectations were met

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The hourlyRate must be greater than 0", "Unexpected error message")

}
