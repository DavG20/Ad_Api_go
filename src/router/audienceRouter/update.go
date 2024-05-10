package audienceRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/audienceService"
	"adtec/backend/src/service/categoryService"
	"adtec/backend/src/service/languageService"
	"adtec/backend/src/utils"
	"context"
	"errors"
)

func Update(ctx context.Context, input model.UpdateAudienceInput) (*model.Audiences, error) {
	if input.Language != nil {
		language, _ := languageService.Get(*input.Language)
		if language == nil {
			return nil, errors.New("language not found")
		}
	}
	if input.Category != nil {
		category, _ := categoryService.Get(*input.Category)
		if category == nil {
			return nil, errors.New("category not found")
		}
	}
	updatedData, err := audienceService.Update(audienceService.UpdateInput{
		CampaignID: input.CampaignID,
		Category:   input.Category,
		Language:   input.Language,
	})
	if err != nil {
		return nil, err
	}

	updatedAudience := utils.AudienceGraphqlConverter(updatedData)
	return updatedAudience, nil
}
