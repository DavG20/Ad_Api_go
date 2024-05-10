package channelContentLanguageServiceTest

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/service"
	"adtec/backend/src/service/channelContentLanguageService"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/service/languageService"
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
	language  string

	channel   *db.ChannelsModel
	languageM *db.LanguagesModel
)

func TestCreateChannelContentLanguage(t *testing.T) {

	channel, _ = channelService.Create(inputGenerator.GenerateRandomChannelData())

	if channel == nil {
		channel, _ = channelService.Create(inputGenerator.GenerateRandomChannelData())

	}

	languageM, _ = languageService.Create(inputGenerator.GenerateRandomLanguageData())

	if languageM == nil {
		languageM, _ = languageService.Create(inputGenerator.GenerateRandomLanguageData())

	}

	input := inputGenerator.GenerateRandomChannelLanguageData()

	input.Language = languageM.Language

	input.ChannelID = channel.ID

	// create an instance of our test object

	result, err := channelContentLanguageService.Create(input)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	channelId = result.ChannelID
	language = result.Language
	assert.Equal(t, channelId, result.ChannelID)
	assert.Equal(t, language, result.Language)

}

func TestCreateChannelContentLanguage_InvalidInput(t *testing.T) {

	// create an instance of our test object

	input := inputGenerator.GenerateRandomChannelLanguageData()

	input.Language = language

	result, err := channelContentLanguageService.Create(input)

	// assert that the expectations were met
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetChannelContentLanguage(t *testing.T) {
	if channelId == "" || language == "" {
		TestCreateChannelContentLanguage(t)
	}

	input := channelContentLanguageService.ChannelContentLanguageInput{
		ChannelID: channelId,
		Language:  language,
	}

	result, err := channelContentLanguageService.Get(input)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, input.ChannelID, result.ChannelID)
	assert.Equal(t, input.Language, result.Language)
}

func TestGetChannelContentLanguage_InvalidInput(t *testing.T) {
	// create an instance of our test object

	input := channelContentLanguageService.ChannelContentLanguageInput{
		ChannelID: "",
		Language:  language,
	}

	result, err := channelContentLanguageService.Get(input)

	// assert that the expectations were met
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetAllChannelContentLanguage(t *testing.T) {
	filter := channelContentLanguageService.FindManyInput{
		Filter: &service.FilterData{},
	}

	result, err := channelContentLanguageService.GetAll(filter)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDeleteLanguage(t *testing.T) {
	if channelId == "" || language == "" {
		TestCreateChannelContentLanguage(t)
	}

	input := channelContentLanguageService.ChannelContentLanguageInput{
		ChannelID: channelId,
		Language:  language,
	}

	result, err := channelContentLanguageService.Delete(input)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.Equal(t, "deleted Successfully", result.Message)
}

func TestDeleteLanguage_InvalidInput(t *testing.T) {
	// create an instance of our test object
	channelID := ""
	languageVal := "English"

	input := channelContentLanguageService.ChannelContentLanguageInput{
		ChannelID: channelID,
		Language:  languageVal,
	}

	result, err := channelContentLanguageService.Delete(input)

	// assert that the expectations were met
	assert.Error(t, err)
	assert.NotNil(t, result)
	assert.False(t, result.Success)
	assert.Equal(t, "channelLanguage not found", result.Message)
}
