package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/router/advertisementRouter"
	"context"
)

// CreateAdvertisement is the resolver for the createAdvertisement field.
func (r *mutationResolver) CreateAdvertisement(ctx context.Context, input model.CreateAdvertisementInput) (*model.Advertisements, error) {
	return advertisementRouter.Create(ctx, input)
}

// UpdateAdvertisement is the resolver for the updateAdvertisement field.
func (r *mutationResolver) UpdateAdvertisement(ctx context.Context, input model.UpdateAdvertisementInput) (*model.Advertisements, error) {
	return advertisementRouter.Update(ctx, input)
}

// AdvertisementByID is the resolver for the advertisementById field.
func (r *queryResolver) AdvertisementByID(ctx context.Context, id string) (*model.Advertisements, error) {
	return advertisementRouter.Get(ctx, id)
}

// GetAllAdvertisements is the resolver for the getAllAdvertisements field.
func (r *queryResolver) GetAllAdvertisements(ctx context.Context, advertisementFilter *model.AdvertisementFilter) ([]*model.Advertisements, error) {
	return advertisementRouter.GetAll(ctx, advertisementFilter)
}
