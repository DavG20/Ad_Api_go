package channelServiceTest

import (
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
	channelID   string
	userNameVal string
)

func TestCreateChannel(t *testing.T) {
	// create an instance of our test object

	input := inputGenerator.GenerateRandomChannelData()

	result, err := channelService.Create(input)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	channelID = result.ID
	userNameVal = result.UserName
	assert.Equal(t, input.OwnerID, result.OwnerID)
	assert.Equal(t, input.UserName, result.UserName)

}

func TestCreateChannel_InvalidInput(t *testing.T) {
	// create an instance of our test object

	input := inputGenerator.GenerateRandomChannelData()

	// make the input invalid

	input.OwnerID = ""

	result, err := channelService.Create(input)

	// assert that the expectations were met
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The ownerId is required", "Unexpected error message")

}

func TestUpdateChannel(t *testing.T) {
	if channelID == "" {
		TestCreateChannel(t)
	}
	id := channelID

	createInput := inputGenerator.GenerateRandomChannelData()

	input := channelService.UpdateInput{
		ID:          id,
		Name:        createInput.Name,
		Description: createInput.Description,
	}
	result, err := channelService.Update(input)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	channelID = result.ID
	userNameVal = result.UserName
	nameVal, ok := result.Name()
	assert.True(t, ok)
	assert.Equal(t, *input.Name, nameVal)
}

func TestUpdateChannel_InvalidId(t *testing.T) {

	id := "invalid_id"

	name := "ChannelName2024"
	description := "This is a sample description"

	input := channelService.UpdateInput{
		ID:          id,
		Name:        &name,
		Description: &description,
	}
	result, err := channelService.Update(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "invalid channel id or channel doesn't exist", "Unexpected error message")
}

func TestGetChannel(t *testing.T) {
	if channelID == "" {
		TestCreateChannel(t)
	}
	id := channelID

	result, err := channelService.Get(id)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)

}

func TestGetChannel_InvalidId(t *testing.T) {
	id := "invalid_id"

	result, err := channelService.Get(id)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "channel not found", "Unexpected error message")
}

func TestGetByUserName(t *testing.T) {
	if userNameVal == "" {
		TestCreateChannel(t)
	}
	userName := userNameVal

	result, err := channelService.GetChannelByUserName(userName)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, userName, result.UserName)
}

func TestGetByUserName_InvalidUserName(t *testing.T) {
	userName := "invalid_userName"

	result, err := channelService.GetChannelByUserName(userName)
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")

}

func TestGetAll(t *testing.T) {

	input := channelService.ChannelFilter{}

	result, err := channelService.GetAll(input)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
