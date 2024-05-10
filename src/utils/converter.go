package utils

import (
	"adtec/backend/src/graphql/model"
	"encoding/json"
)

func LimitSetter(limit int32) int32 {
	if limit == 0 || limit >= 50 {
		limit = 50
	}
	return limit
}

func AccountGraphqlConverter(data interface{}) *model.Account {
	j, _ := json.Marshal(data)

	var account *model.Account
	_ = json.Unmarshal(j, &account)

	return account
}
func CompanyBankingGraphqlConverter(data interface{}) *model.CompanyBankings {
	j, _ := json.Marshal(data)

	var companyBanking *model.CompanyBankings
	_ = json.Unmarshal(j, &companyBanking)

	return companyBanking
}
func BankGraphqlConverter(data interface{}) *model.Banks {
	j, _ := json.Marshal(data)

	var bank *model.Banks
	_ = json.Unmarshal(j, &bank)

	return bank
}
func ChannelGraphqlConverter(data interface{}) *model.Channels {
	j, _ := json.Marshal(data)

	var channel *model.Channels
	_ = json.Unmarshal(j, &channel)

	return channel
}

func CategoryGraphqlConverter(data interface{}) *model.Categories {
	j, _ := json.Marshal(data)

	var category *model.Categories
	_ = json.Unmarshal(j, &category)

	return category
}
func ChannelCategoryGraphqlConverter(data interface{}) *model.ChannelCategories {
	j, _ := json.Marshal(data)

	var channelCategory *model.ChannelCategories
	_ = json.Unmarshal(j, &channelCategory)

	return channelCategory
}
func LanguageGraphqlConverter(data interface{}) *model.Languages {
	j, _ := json.Marshal(data)

	var language *model.Languages
	_ = json.Unmarshal(j, &language)

	return language
}

func ChannelContentLanguageGraphqlConverter(data interface{}) *model.ChannelContentLanguages {
	j, _ := json.Marshal(data)

	var channelContentLanguage *model.ChannelContentLanguages
	_ = json.Unmarshal(j, &channelContentLanguage)

	return channelContentLanguage
}

func HourlyRateGraphqlConverter(data interface{}) *model.HourlyRates {
	j, _ := json.Marshal(data)

	var hourlyRates *model.HourlyRates
	_ = json.Unmarshal(j, &hourlyRates)

	return hourlyRates
}

func CPMRateGraphqlConverter(data interface{}) *model.CpmRates {
	j, _ := json.Marshal(data)

	var cmpRates *model.CpmRates
	_ = json.Unmarshal(j, &cmpRates)

	return cmpRates
}

func ChannelDetailGraphqlConverter(data interface{}) *model.ChannelDetails {
	j, _ := json.Marshal(data)

	var channelDetail *model.ChannelDetails
	_ = json.Unmarshal(j, &channelDetail)

	return channelDetail
}

func CampaignGraphqlConverter(data interface{}) *model.Campaigns {
	j, _ := json.Marshal(data)

	var campaign *model.Campaigns
	_ = json.Unmarshal(j, &campaign)

	return campaign
}

func AudienceGraphqlConverter(data interface{}) *model.Audiences {
	j, _ := json.Marshal(data)

	var audience *model.Audiences
	_ = json.Unmarshal(j, &audience)

	return audience
}
func BudgetGraphqlConverter(data interface{}) *model.Budgets {
	j, _ := json.Marshal(data)

	var budget *model.Budgets
	_ = json.Unmarshal(j, &budget)

	return budget
}
func ContentGraphqlConverter(data interface{}) *model.Contents {
	j, _ := json.Marshal(data)

	var content *model.Contents
	_ = json.Unmarshal(j, &content)

	return content
}

func ContentLinkGraphqlConverter(data interface{}) *model.ContentLinks {
	j, _ := json.Marshal(data)

	var contentLink *model.ContentLinks
	_ = json.Unmarshal(j, &contentLink)

	return contentLink
}
func AdvertisementGraphqlConverter(data interface{}) *model.Advertisements {
	j, _ := json.Marshal(data)

	var advertisement *model.Advertisements
	_ = json.Unmarshal(j, &advertisement)

	return advertisement
}
func AdvertisementCPMGraphqlConverter(data interface{}) *model.AdvertisementCPMs {
	j, _ := json.Marshal(data)

	var advertisementCPM *model.AdvertisementCPMs
	_ = json.Unmarshal(j, &advertisementCPM)

	return advertisementCPM
}
func AdvertisementResultGraphqlConverter(data interface{}) *model.AdvertisementResults {
	j, _ := json.Marshal(data)

	var advertisementResult *model.AdvertisementResults
	_ = json.Unmarshal(j, &advertisementResult)

	return advertisementResult
}
func LocationGraphqlConverter(data interface{}) *model.Locations {
	j, _ := json.Marshal(data)

	var location *model.Locations
	_ = json.Unmarshal(j, &location)

	return location
}
func CompanyGraphqlConverter(data interface{}) *model.Companies {
	j, _ := json.Marshal(data)

	var company *model.Companies
	_ = json.Unmarshal(j, &company)

	return company
}
func AccountBankingGraphqlConverter(data interface{}) *model.AccountBanking {
	j, _ := json.Marshal(data)

	var accountBanking *model.AccountBanking
	_ = json.Unmarshal(j, &accountBanking)

	return accountBanking
}
func CompanyLocationGraphqlConverter(data interface{}) *model.CompanyLocations {
	j, _ := json.Marshal(data)

	var companyLocation *model.CompanyLocations
	_ = json.Unmarshal(j, &companyLocation)

	return companyLocation
}

func CompanyMemberGraphqlConverter(data interface{}) *model.CompanyMembers {
	j, _ := json.Marshal(data)

	var companyMember *model.CompanyMembers
	_ = json.Unmarshal(j, &companyMember)

	return companyMember
}
func FundingGraphqlConverter(data interface{}) *model.Funding {
	j, _ := json.Marshal(data)

	var funding *model.Funding
	_ = json.Unmarshal(j, &funding)

	return funding
}

func ByteConvertor(data interface{}) []byte {
	j, _ := json.Marshal(data)

	var response []byte
	_ = json.Unmarshal(j, &response)

	return response
}
func ChannelLifeTimeBalanceGraphqlConverter(data interface{}) *model.ChannelLifeTimeBalances {
	j, _ := json.Marshal(data)

	var channelLifeTimeBalance *model.ChannelLifeTimeBalances
	_ = json.Unmarshal(j, &channelLifeTimeBalance)

	return channelLifeTimeBalance
}
func ChannelBalanceGraphqlConverter(data interface{}) *model.ChannelBalances {
	j, _ := json.Marshal(data)

	var channelBalance *model.ChannelBalances
	_ = json.Unmarshal(j, &channelBalance)

	return channelBalance
}
func AdPaymentGraphqlConverter(data interface{}) *model.AdPayments {
	j, _ := json.Marshal(data)

	var adPayments *model.AdPayments
	_ = json.Unmarshal(j, &adPayments)

	return adPayments
}
func ChannelCollectedAdLinkGraphqlConverter(data interface{}) *model.ChannelCollectedAdLinks {
	j, _ := json.Marshal(data)

	var channelCollectedAdLink *model.ChannelCollectedAdLinks
	_ = json.Unmarshal(j, &channelCollectedAdLink)

	return channelCollectedAdLink
}
func WithdrawalGraphqlConverter(data interface{}) *model.Withdrawals {
	j, _ := json.Marshal(data)

	var withdrawal *model.Withdrawals
	_ = json.Unmarshal(j, &withdrawal)

	return withdrawal
}
func WithdrawalChannelLinkGraphqlConverter(data interface{}) *model.WithdrawalChannelLinks {
	j, _ := json.Marshal(data)

	var withdrawalChannelLink *model.WithdrawalChannelLinks
	_ = json.Unmarshal(j, &withdrawalChannelLink)

	return withdrawalChannelLink
}
