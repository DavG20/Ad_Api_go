package categoryServiceTest

import (
	"adtec/backend/src/service/categoryService"
	"adtec/backend/src/utils/inputGenerator"
	"adtec/backend/tests/serviceTest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	serviceTest.SetUpEnv()
}

var (
	category string
)

func TestCreateCategory(t *testing.T) {
	// create an instance of our test object
	input := inputGenerator.GenerateRandomCategoryData()

	result, err := categoryService.Create(input)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	category = result.Category
	assert.Equal(t, input.Category, result.Category)

	// add more assertions based on your specific test case
}

func TestCreateCategory_InvalidInput(t *testing.T) {
	// create an instance of our test object
	categoryName := ""

	input := categoryService.CreateInput{
		Category: categoryName,
	}

	result, err := categoryService.Create(input)

	// assert that the expectations were met
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The category is required", "Unexpected error message")

	// add more assertions based on your specific test case
}

func TestGetCategory(t *testing.T) {
	if category == "" {
		TestCreateCategory(t)
	}

	result, err := categoryService.Get(category)

	// assert that the expectations were met

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category, result.Category)

}

func TestGetCategory_InvalidInput(t *testing.T) {
	categoryName := ""

	result, err := categoryService.Get(categoryName)

	// assert that the expectations were met

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")

}

func TestGetAll(t *testing.T) {
	input := categoryService.FindManyInput{}

	result, err := categoryService.GetAll(input)

	// assert that the expectations were met

	assert.NoError(t, err)
	assert.NotNil(t, result)

}
