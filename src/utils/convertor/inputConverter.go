package convertor

import (
	"adtec/backend/src/service/channelCategoryService"
	"adtec/backend/src/service/channelCollectedAdLinkService"
	"adtec/backend/src/service/channelContentLanguageService"
	"adtec/backend/src/service/withdrawalChannelLinkService"
	"encoding/json"
)

func ChannelCollectedAdLinkInputConverter(data interface{}) channelCollectedAdLinkService.ChannelCollectedAdLinkInput {
	j, _ := json.Marshal(data)

	var response channelCollectedAdLinkService.ChannelCollectedAdLinkInput

	_ = json.Unmarshal(j, &response)

	return response
}
func ChannelContentLanguageInputConverter(data interface{}) channelContentLanguageService.ChannelContentLanguageInput {
	j, _ := json.Marshal(data)

	var response channelContentLanguageService.ChannelContentLanguageInput

	_ = json.Unmarshal(j, &response)

	return response
}
func ChannelCategoryInputConverter(data interface{}) channelCategoryService.ChannelCategoryInput {
	j, _ := json.Marshal(data)

	var response channelCategoryService.ChannelCategoryInput

	_ = json.Unmarshal(j, &response)

	return response
}
func WithdrawalChannelLinkInputConverter(data interface{}) withdrawalChannelLinkService.WithdrawalChannelLinkInput {
	j, _ := json.Marshal(data)

	var response withdrawalChannelLinkService.WithdrawalChannelLinkInput

	_ = json.Unmarshal(j, &response)

	return response
}
