package accountServiceTest

import (
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/utils/inputGenerator"
	"adtec/backend/tests/serviceTest"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func init() {
	serviceTest.SetUpEnv()
}

var (
	id       string
	userName string
	emailVal string
	phoneVal string
)

func TestCreateAccount(t *testing.T) {
	// create an instance of our test object

	input := inputGenerator.GenerateRandomAccountData()

	result, err := accountService.Create(input)

	// assert that the expectations were met
	assert.NoError(t, err)
	assert.NotNil(t, result)
	id = result.ID
	userName = result.UserName
	emailVal, _ = result.Email()
	phoneVal, _ = result.PhoneNumber()
	assert.Equal(t, input.UserID, result.UserID)
	assert.Equal(t, input.UserName, result.UserName)

	// add more assertions based on your specific test case
}

func TestCreateAccount_InvalidInput(t *testing.T) {
	// Create an instance of our test object

	input := inputGenerator.GenerateRandomAccountData()

	input.UserID = ""

	result, err := accountService.Create(input)

	// Add assertions based on the specific error scenario
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The userId is required", "Unexpected error message")
}

func TestCreateAccount_InvalidUserName(t *testing.T) {
	// Create an instance of our test object

	input := inputGenerator.GenerateRandomAccountData()

	input.UserName = "tes"

	result, err := accountService.Create(input)

	// Add assertions based on the specific error scenario
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "The length of userName should be equal to or greater than 4", "Unexpected error message")
}

func TestGetAccount(t *testing.T) {
	// create an instance of our test object

	if id == "" {
		TestCreateAccount(t)
	}
	id := id

	// Call the method we want to test
	result, err := accountService.Get(id)

	// assert that the expectations were met
	// assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, id, result.ID)
	// Add more assertions based on your specific test case
}

// Add more tests as needed for different scenarios

func TestGetAccount_InvalidID(t *testing.T) {
	// Create an instance of our test object

	// Call the method we want to test with an invalid ID
	result, err := accountService.Get("invalidID")

	// Add assertions based on the specific error scenario
	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")
}

func TestUpdateAccount(t *testing.T) {

	if id == "" {
		TestCreateAccount(t)
	}
	id := id
	inputData := inputGenerator.GenerateRandomAccountData()
	email := inputData.Email
	phoneNumber := inputData.PhoneNumber

	birthDateVal := "1999-01-01"
	var birthDate *time.Time

	// Parse the birthDateVal, and if successful, assign it to birthDate
	date, err := time.Parse("2006-01-02", birthDateVal)
	if err != nil {
		t.Errorf("Error parsing date: %v", err)
		return
	}
	fmt.Println("date", date)
	birthDate = &date

	input := accountService.UpdateInput{
		ID:          id,
		Email:       email,
		PhoneNumber: phoneNumber,
		BirthDate:   birthDate,
	}
	result, err := accountService.Update(input)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	userName = result.UserName
	assert.Equal(t, id, result.ID)
	resEmail, _ := result.Email()
	phone, _ := result.PhoneNumber()

	phoneVal = phone
	emailVal = resEmail
	assert.Equal(t, phoneNumber, &phone)
	assert.Equal(t, email, &resEmail)

}

func TestUpdateAccount_InvalidID(t *testing.T) {
	id := "invalid_id"
	userName := "updatedUserName"
	input := accountService.UpdateInput{
		ID:       id,
		UserName: &userName,
	}
	result, err := accountService.Update(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "account not found", "Unexpected error message")
}

func TestGetAccountByEmail(t *testing.T) {
	fmt.Println("emailVal", emailVal)
	if emailVal == "" {
		TestCreateAccount(t)
	}
	email := emailVal

	result, err := accountService.GetByEmail(email)
	fmt.Println("email", email, err)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	resEmail, _ := result.Email()
	assert.Equal(t, email, resEmail)
}

func TestGetAccountByEmail_InvalidEmail(t *testing.T) {

	invalidEmail := "invalid_email"

	result, err := accountService.GetByEmail(invalidEmail)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")

}

func TestGetAccountByPhoneNumber(t *testing.T) {

	if phoneVal == "" {
		TestCreateAccount(t)
	}
	phoneNumber := phoneVal

	result, err := accountService.GetByPhoneNumber(phoneNumber)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	phone, _ := result.PhoneNumber()
	assert.Equal(t, phoneNumber, phone)
}

func TestGetAccountByPhoneNumber_InvalidPhoneNumber(t *testing.T) {

	invalidPhoneNumber := "invalid_phone_number"

	result, err := accountService.GetByPhoneNumber(invalidPhoneNumber)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "Unexpected error message")
}

func TestGetAccountByUsername(t *testing.T) {

	if userName == "" {
		TestCreateAccount(t)
	}
	username := userName
	result, err := accountService.GetByUserName(username)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, username, result.UserName)
}

func TestGetAccountByUsername_InvalidUsername(t *testing.T) {

	invalidUsername := "invalid_username"

	result, err := accountService.GetByPhoneNumber(invalidUsername)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "ErrNotFound", "unexpected error message")
}
