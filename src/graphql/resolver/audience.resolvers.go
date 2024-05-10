package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/router/audienceRouter"
	"context"
)

// CreateAudience is the resolver for the createAudience field.
func (r *mutationResolver) CreateAudience(ctx context.Context, input model.CreateAudienceInput) (*model.Audiences, error) {
	return audienceRouter.Create(ctx, input)
}

// UpdateAudience is the resolver for the updateAudience field.
func (r *mutationResolver) UpdateAudience(ctx context.Context, input model.UpdateAudienceInput) (*model.Audiences, error) {
	return audienceRouter.Update(ctx, input)
}

// AudienceByID is the resolver for the audienceById field.
func (r *queryResolver) AudienceByID(ctx context.Context, id string) (*model.Audiences, error) {
	return audienceRouter.Get(ctx, id)
}
