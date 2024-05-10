package languageServiceTest

import (
	"adtec/backend/src/service"
	"adtec/backend/src/service/languageService"
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
	language string
)

func TestCreateLanguage(t *testing.T) {
	// create an instance of our test object
	input := inputGenerator.GenerateRandomLanguageData()

	result, err := languageService.Create(input)
	fmt.Println("result", input.Language)
	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	language = result.Language
	assert.Equal(t, input.Language, result.Language)

	// add more assertions based on your specific test case
}

func TestCreateLanguage_InvalidInput(t *testing.T) {
	// create an instance of our test object
	languageName := ""

	input := languageService.CreateInput{
		Language: languageName,
	}

	result, err := languageService.Create(input)

	// assert that the expectations were met
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The language is required", "Unexpected error message")

	// add more assertions based on your specific test case
}

func TestGetLanguage(t *testing.T) {
	if language == "" {
		TestCreateLanguage(t)
	}

	result, err := languageService.Get(language)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, language, result.Language)

	// add more assertions based on your specific test case
}

func TestGetLanguage_InvalidInput(t *testing.T) {
	language := ""
	result, err := languageService.Get(language)

	// assert that the expectations were met
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetAllLanguage(t *testing.T) {
	filter := languageService.FindManyInput{
		Filter: &service.FilterData{},
	}
	result, err := languageService.GetAll(filter)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// add more assertions based on your specific test case
}
