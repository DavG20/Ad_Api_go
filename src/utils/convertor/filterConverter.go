package convertor

import (
	"adtec/backend/src/service"
	"adtec/backend/src/service/adPaymentService"
	"adtec/backend/src/service/advertisementService"
	"adtec/backend/src/service/campaignService"
	"adtec/backend/src/service/categoryService"
	"adtec/backend/src/service/channelCategoryService"
	"adtec/backend/src/service/channelCollectedAdLinkService"
	"adtec/backend/src/service/channelContentLanguageService"
	"adtec/backend/src/service/channelDetailService"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/service/companyMemberService"
	"adtec/backend/src/service/fundingService"
	"adtec/backend/src/service/languageService"
	"adtec/backend/src/service/withdrawalChannelLinkService"
	"adtec/backend/src/service/withdrawalService"
	"encoding/json"
)

func ChannelFilterConverter(data interface{}) channelService.ChannelFilter {
	j, _ := json.Marshal(data)

	var response channelService.ChannelFilter
	_ = json.Unmarshal(j, &response)

	return response
}

func FilterConverter(data interface{}) service.FilterData {
	j, _ := json.Marshal(data)

	var response service.FilterData
	_ = json.Unmarshal(j, &response)

	return response
}

func CategoryFilterConverter(data interface{}) categoryService.FindManyInput {
	j, _ := json.Marshal(data)

	var response categoryService.FindManyInput
	_ = json.Unmarshal(j, &response)

	return response
}

func ChannelCategoryFilterConverter(data interface{}) channelCategoryService.FindManyInput {
	j, _ := json.Marshal(data)

	var response channelCategoryService.FindManyInput
	_ = json.Unmarshal(j, &response)

	return response
}

func ChannelContentLanguageFilterConverter(data interface{}) channelContentLanguageService.FindManyInput {
	j, _ := json.Marshal(data)

	var response channelContentLanguageService.FindManyInput
	_ = json.Unmarshal(j, &response)

	return response
}
func ChannelDetailFilterConverter(data interface{}) channelDetailService.FindManyInput {
	j, _ := json.Marshal(data)

	var response channelDetailService.FindManyInput
	_ = json.Unmarshal(j, &response)

	return response
}
func LanguageFilterConverter(data interface{}) languageService.FindManyInput {
	j, _ := json.Marshal(data)

	var response languageService.FindManyInput
	_ = json.Unmarshal(j, &response)

	return response
}
func CampaignFilterConverter(data interface{}) campaignService.CampaignFilter {
	j, _ := json.Marshal(data)

	var response campaignService.CampaignFilter
	_ = json.Unmarshal(j, &response)

	return response
}
func AdvertisementFilterConverter(data interface{}) advertisementService.AdvertisementFilter {
	j, _ := json.Marshal(data)

	var response advertisementService.AdvertisementFilter
	_ = json.Unmarshal(j, &response)

	return response
}
func CompanyMemberFilterConverter(data interface{}) companyMemberService.CompanyMemberFilter {
	j, _ := json.Marshal(data)

	var response companyMemberService.CompanyMemberFilter

	_ = json.Unmarshal(j, &response)

	return response
}
func FundingFilterConverter(data interface{}) fundingService.FundingFilter {
	j, _ := json.Marshal(data)

	var response fundingService.FundingFilter

	_ = json.Unmarshal(j, &response)

	return response
}
func AdPaymentFilterConverter(data interface{}) adPaymentService.AdPaymentFilter {
	j, _ := json.Marshal(data)

	var response adPaymentService.AdPaymentFilter

	_ = json.Unmarshal(j, &response)

	return response
}
func ChannelCollectedAdLinkFilterConverter(data interface{}) channelCollectedAdLinkService.ChannelCollectedAdLinkFilter {
	j, _ := json.Marshal(data)

	var response channelCollectedAdLinkService.ChannelCollectedAdLinkFilter

	_ = json.Unmarshal(j, &response)

	return response
}
func WithdrawalFilterConverter(data interface{}) withdrawalService.WithdrawalFilter {
	j, _ := json.Marshal(data)

	var response withdrawalService.WithdrawalFilter

	_ = json.Unmarshal(j, &response)

	return response
}
func WithdrawalChannelLinkFilterConverter(data interface{}) withdrawalChannelLinkService.WithdrawalChannelLinkFilter {
	j, _ := json.Marshal(data)

	var response withdrawalChannelLinkService.WithdrawalChannelLinkFilter

	_ = json.Unmarshal(j, &response)

	return response
}
