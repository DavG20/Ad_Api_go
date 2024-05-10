package channelCategoryServiceTest

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/service"
	"adtec/backend/src/service/categoryService"
	"adtec/backend/src/service/channelCategoryService"
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
	category  string

	channel *db.ChannelsModel

	categoryM *db.CategoriesModel
)

func TestCreateChannelCategory(t *testing.T) {
	// create an instance of our test object

	channel, _ = channelService.Create(inputGenerator.GenerateRandomChannelData())

	if channel == nil {
		channel, _ = channelService.Create(inputGenerator.GenerateRandomChannelData())

	}

	categoryM, _ = categoryService.Create(inputGenerator.GenerateRandomCategoryData())

	if categoryM == nil {
		categoryM, _ = categoryService.Create(inputGenerator.GenerateRandomCategoryData())

	}

	input := inputGenerator.GenerateRandomChannelCategoryData()

	input.ChannelID = channel.ID
	input.Category = categoryM.Category

	result, err := channelCategoryService.Create(input)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	channelId = result.ChannelID
	category = result.Category
	assert.Equal(t, input.ChannelID, result.ChannelID)
	assert.Equal(t, input.Category, result.Category)

	// add more assertions based on your specific test case
}

func TestCreateChannelCategory_InvalidInput(t *testing.T) {
	// create an instance of our test object

	input := inputGenerator.GenerateRandomChannelCategoryData()

	// make the input invalid

	input.ChannelID = ""

	result, err := channelCategoryService.Create(input)

	// assert that the expectations were met
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The channelId is required", "Unexpected error message")

}

func TestGetChannelCategory(t *testing.T) {
	if category == "" || channelId == "" {
		TestCreateChannelCategory(t)
	}
	input := channelCategoryService.ChannelCategoryInput{
		ChannelID: channelId,
		Category:  category,
	}
	result, err := channelCategoryService.Get(input)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, input.ChannelID, result.ChannelID)
	assert.Equal(t, input.Category, result.Category)
}

func TestGetChannelCategory_InvalidInput(t *testing.T) {
	channelID := ""
	categoryName := ""
	input := channelCategoryService.ChannelCategoryInput{
		ChannelID: channelID,
		Category:  categoryName,
	}
	result, err := channelCategoryService.Get(input)

	// assert that the expectations were met

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")
}

func TestDeleteChannelCategory(t *testing.T) {
	if category == "" || channelId == "" {
		TestCreateChannelCategory(t)
	}
	input := channelCategoryService.ChannelCategoryInput{
		ChannelID: channelId,
		Category:  category,
	}
	result, err := channelCategoryService.Delete(input)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.True(t, result.Success)
	assert.Equal(t, "deleted Successfully", result.Message)

}

func TestGetAllChannelCategory(t *testing.T) {
	input := channelCategoryService.FindManyInput{
		Filter: &service.FilterData{},
	}

	result, err := channelCategoryService.GetAll(input)

	assert.NoError(t, err)
	assert.NotNil(t, result)
}
func TestDeleteChannelCategory_InvalidInput(t *testing.T) {
	channelID := ""
	categoryName := ""
	input := channelCategoryService.ChannelCategoryInput{
		ChannelID: channelID,
		Category:  categoryName,
	}
	result, err := channelCategoryService.Delete(input)

	// assert that the expectations were met

	assert.Error(t, err)
	assert.NotNil(t, result)
	assert.EqualError(t, err, "ChannelCategory not Found", "Unexpected error message")
}
