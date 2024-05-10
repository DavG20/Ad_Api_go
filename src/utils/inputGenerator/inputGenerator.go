package inputGenerator

import (
	"adtec/backend/src/service"
	"adtec/backend/src/service/CPMRateService"
	"adtec/backend/src/service/accountBankingService"
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/service/bankService"
	"adtec/backend/src/service/categoryService"
	"adtec/backend/src/service/channelCategoryService"
	"adtec/backend/src/service/channelContentLanguageService"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/service/companyBankingService"
	"adtec/backend/src/service/companyLocationService"
	"adtec/backend/src/service/companyMemberService"
	"adtec/backend/src/service/companyService"
	"adtec/backend/src/service/hourlyRateService"
	"adtec/backend/src/service/languageService"
	"adtec/backend/src/service/locationService"
	"fmt"
	"math/rand"
)

func GenerateRandomAccountData() accountService.CreateInput {

	userID := fmt.Sprintf("userID%d", rand.Intn(1000000)) // use the generated userId for now, will change later when the user service is ready
	userName := fmt.Sprintf("userName%d", rand.Intn(1000000))
	email := fmt.Sprintf("email%d@example.com", rand.Intn(100000))
	phoneNumber := fmt.Sprintf("+251456789%d", rand.Intn(100000))

	// You can add more fields as needed

	return accountService.CreateInput{
		UserID:      userID,
		UserName:    userName,
		Email:       &email,
		PhoneNumber: &phoneNumber,
	}
}

func GenerateRandomAccountBankingData() accountBankingService.CreateInput {
	accountId := fmt.Sprintf("accountId%d", rand.Intn(1000000))
	bankId := fmt.Sprintf("bankId%d", rand.Intn(1000000))
	fullNameOnBank := fmt.Sprintf("fullNameOnBank%d", rand.Intn(10000))
	BankAccount := fmt.Sprintf("BankAccount%d", rand.Intn(100000))
	currency := "USD"

	return accountBankingService.CreateInput{
		AccountId:      accountId,
		BankId:         bankId,
		FullNameOnBank: fullNameOnBank,
		BankAccount:    BankAccount,
		Currency:       currency,
	}
}
func GenerateRandomBankData() bankService.CreateInput {
	bankName := fmt.Sprintf("bankName%d", rand.Intn(1000000))
	bankCode := fmt.Sprintf("bankCode%d", rand.Intn(1000000))

	return bankService.CreateInput{
		BankName: bankName,
		BankCode: bankCode,
	}
}

func GenerateRandomCompanyBankingData() companyBankingService.CreateInput {

	companyID := fmt.Sprintf("companyID%d", rand.Intn(1000000))
	bankID := fmt.Sprintf("bankID%d", rand.Intn(1000000))
	fullNameOnBank := fmt.Sprintf("fullNameOnBank%d", rand.Intn(10000))
	bankAccount := fmt.Sprintf("bankAccount%d", rand.Intn(100000))
	currency := "USD"

	return companyBankingService.CreateInput{
		CompanyID:      companyID,
		BankId:         bankID,
		FullNameOnBank: fullNameOnBank,
		BankAccount:    bankAccount,
		Currency:       currency,
	}
}

func GenerateRandomCompanyData() companyService.CreateInput {

	creatorId := fmt.Sprintf("creatorId%d", rand.Intn(1000000))
	companyName := fmt.Sprintf("companyName%d", rand.Intn(1000000))
	uniqueName := fmt.Sprintf("uniqueName%d", rand.Intn(1000000))
	tinNumber := fmt.Sprintf("tinNumber%d", rand.Intn(1000000))
	vatNumber := fmt.Sprintf("vatNumber%d", rand.Intn(1000000))

	return companyService.CreateInput{
		CreatorID:  creatorId,
		Name:       companyName,
		UniqueName: uniqueName,
		TinNumber:  &tinNumber,
		VatNumber:  &vatNumber,
	}

}

func GenerateRandomCompanyMemberData() companyMemberService.CreateInput {
	companyID := fmt.Sprintf("companyID%d", rand.Intn(1000000))
	accountID := fmt.Sprintf("account%d", rand.Intn(1000000))

	role := service.CompanyRoleMember
	return companyMemberService.CreateInput{
		CompanyID: companyID,
		AccountID: accountID,
		Role:      &role,
	}

}

func GenerateRandomCompanyLocationData() companyLocationService.CreateInput {
	companyID := fmt.Sprintf("companyID%d", rand.Intn(1000000))
	locationID := fmt.Sprintf("locationID%d", rand.Intn(1000000))

	return companyLocationService.CreateInput{
		CompanyID:  companyID,
		LocationID: locationID,
	}
}

func GenerateRandomLocationData() locationService.CreateInput {
	country := "ETH" // since the validation is iso_3166-1_alpha-3 , it means three letter country code
	city := fmt.Sprintf("city%d", rand.Intn(1000000))
	state := fmt.Sprintf("state%d", rand.Intn(1000000))
	address := fmt.Sprintf("address%d", rand.Intn(1000000))
	postalCode := fmt.Sprintf("postalCode%d", rand.Intn(1000000))

	return locationService.CreateInput{
		Country:    country,
		State:      state,
		City:       city,
		Address:    &address,
		PostalCode: &postalCode,
	}
}

// channel part

func GenerateRandomCategoryData() categoryService.CreateInput {
	category := fmt.Sprintf("category%d", rand.Intn(1000000))

	return categoryService.CreateInput{
		Category: category,
	}
}

func GenerateRandomChannelCategoryData() channelCategoryService.ChannelCategoryInput {
	channelID := fmt.Sprintf("channelID%d", rand.Intn(1000000))
	category := fmt.Sprintf("category%d", rand.Intn(1000000))

	return channelCategoryService.ChannelCategoryInput{
		ChannelID: channelID,
		Category:  category,
	}
}

func GenerateRandomChannelData() channelService.CreateInput {
	ownerID := fmt.Sprintf("ownerID%d", rand.Intn(1000000))
	userName := fmt.Sprintf("userName%d", rand.Intn(1000000))
	name := fmt.Sprintf("name%d", rand.Intn(1000000))
	description := fmt.Sprintf("this is sample description %d", rand.Intn(1000000))
	country := "ETH"

	return channelService.CreateInput{
		OwnerID:     ownerID,
		UserName:    userName,
		Name:        &name,
		Description: &description,
		Country:     &country,
	}
}

func GenerateRandomCPMRateData() CPMRateService.CreateInput {
	channelId := fmt.Sprintf("channelId%d", rand.Intn(1000000))
	cpmRate := rand.Float64()
	active := false
	minCPMVolume := rand.Float64()

	return CPMRateService.CreateInput{
		ChannelID:    channelId,
		Active:       &active,
		Cpm:          &cpmRate,
		MinCPMVolume: &minCPMVolume,
	}
}

func GenerateRandomHourlyRateData() hourlyRateService.CreateInput {
	channelID := fmt.Sprintf("channelID%d", rand.Intn(1000000))
	active := false
	hourlyRate := rand.Float64()
	minHourlyVolume := rand.Float64()
	maxHourlyVolume := rand.Float64()

	return hourlyRateService.CreateInput{
		ChannelID:       channelID,
		Active:          &active,
		HourlyRate:      &hourlyRate,
		MinHourlyVolume: &minHourlyVolume,
		MaxHourlyVolume: &maxHourlyVolume,
	}
}

func GenerateRandomLanguageData() languageService.CreateInput {
	language := fmt.Sprintf("language%d", rand.Intn(100))

	return languageService.CreateInput{
		Language: language,
	}

}

func GenerateRandomChannelLanguageData() channelContentLanguageService.ChannelContentLanguageInput {
	channelID := fmt.Sprintf("channelID%d", rand.Intn(1000000))
	language := fmt.Sprintf("language%d", rand.Intn(100))

	return channelContentLanguageService.ChannelContentLanguageInput{
		ChannelID: channelID,
		Language:  language,
	}
}
