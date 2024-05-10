package adPaymentRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/accountService"
	"adtec/backend/src/service/adPaymentService"
	"adtec/backend/src/service/advertisementService"
	"adtec/backend/src/service/campaignService"
	"adtec/backend/src/service/channelService"
	"adtec/backend/src/utils"
	"context"
	"errors"
	"time"
)

func Create(ctx context.Context, input model.CreateAdPaymentInput) (*model.AdPayments, error) {
	campaign, _ := campaignService.Get(input.CampaignID)

	if campaign == nil {
		return nil, errors.New("invalid campaign Id ")
	}
	advertisement, _ := advertisementService.Get(input.AdvertisementID)

	if advertisement == nil {
		return nil, errors.New("invalid advertisement Id ")
	}
	channel, _ := channelService.Get(input.ChannelID)

	if channel == nil {
		return nil, errors.New("invalid channel Id ")
	}
	account, _ := accountService.Get(input.AccountID)

	if account == nil {
		return nil, errors.New("invalid account Id ")
	}

	var completionTime *time.Time

	if input.CompletionTime != nil {
		date, _ := time.Parse("2006-01-02", *input.CompletionTime)

		completionTime = &date
	}

	adPaymentData, err := adPaymentService.Create(
		adPaymentService.CreateInput{
			AdvertisementID: input.AdvertisementID,
			AccountID:       input.AccountID,
			CampaignID:      input.CampaignID,
			ChannelID:       input.ChannelID,
			Amount:          input.Amount,
			Currency:        input.Currency,
			CompletionTime:  completionTime,
		},
	)
	if err != nil {
		return nil, err
	}

	createdAdPayment := utils.AdPaymentGraphqlConverter(adPaymentData)

	return createdAdPayment, nil
}
