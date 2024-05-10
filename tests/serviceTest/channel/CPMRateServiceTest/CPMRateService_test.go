package CPMRateServiceTest

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/service/CPMRateService"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/utils/inputGenerator"
	"adtec/backend/tests/serviceTest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	serviceTest.SetUpEnv()
}

var (
	channelId string

	channel *db.ChannelsModel
)

func TestCreateCPMRate(t *testing.T) {

	channel, _ = channelService.Create(inputGenerator.GenerateRandomChannelData())

	if channel == nil {
		channel, _ = channelService.Create(inputGenerator.GenerateRandomChannelData())
	}

	// create an instance of our test object

	input := inputGenerator.GenerateRandomCPMRateData()

	input.ChannelID = channel.ID

	result, err := CPMRateService.Create(input)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	channelId = result.ChannelID
	assert.Equal(t, input.ChannelID, result.ChannelID)

	// add more assertions based on your specific test case
}

func TestCreateCPMRate_InvalidInput(t *testing.T) {
	// create an instance of our test object
	input := inputGenerator.GenerateRandomCPMRateData()

	// make the input invalid

	input.ChannelID = ""

	result, err := CPMRateService.Create(input)

	// assert that the expectations were met
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The channelId is required", "unexpected error message")
}

func TestUpdateCPMRate(t *testing.T) {
	if channelId == "" {
		TestCreateCPMRate(t)
	}

	randomCpm := inputGenerator.GenerateRandomCPMRateData()

	input := CPMRateService.UpdateInput{
		ChannelID: channelId,
		Active:    randomCpm.Active,
		Cpm:       randomCpm.Cpm,
	}

	result, err := CPMRateService.Update(input)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, input.ChannelID, result.ChannelID)

}
