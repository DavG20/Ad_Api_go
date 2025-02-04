package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/router/advertisementResultRouter"
	"context"
)

// CreateAdvertisementResult is the resolver for the createAdvertisementResult field.
func (r *mutationResolver) CreateAdvertisementResult(ctx context.Context, input model.CreateAdvertisementResultInput) (*model.AdvertisementResults, error) {
	return advertisementResultRouter.Create(ctx, input)
}

// UpdateAdvertisementResult is the resolver for the updateAdvertisementResult field.
func (r *mutationResolver) UpdateAdvertisementResult(ctx context.Context, input model.UpdateAdvertisementResultInput) (*model.AdvertisementResults, error) {
	return advertisementResultRouter.Update(ctx, input)
}

// AdvertisementResultByID is the resolver for the advertisementResultById field.
func (r *queryResolver) AdvertisementResultByID(ctx context.Context, id string) (*model.AdvertisementResults, error) {
	return advertisementResultRouter.Get(ctx, id)
}
