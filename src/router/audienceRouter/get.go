package audienceRouter

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/service/audienceService"
	"adtec/backend/src/utils"
	"context"
)

func Get(ctx context.Context, id string) (*model.Audiences, error) {
	audienceData, err := audienceService.Get(id)

	if err != nil {
		return nil, err
	}

	audience := utils.AudienceGraphqlConverter(audienceData)

	return audience, nil
}
