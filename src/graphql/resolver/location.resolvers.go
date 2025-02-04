package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/router/locationRouter"
	"context"
)

// CreateLocation is the resolver for the createLocation field.
func (r *mutationResolver) CreateLocation(ctx context.Context, input model.CreateLocationInput) (*model.Locations, error) {
	return locationRouter.Create(ctx, input)
}

// UpdateLocation is the resolver for the updateLocation field.
func (r *mutationResolver) UpdateLocation(ctx context.Context, input model.UpdateLocationInput) (*model.Locations, error) {
	return locationRouter.Update(ctx, input)
}

// LocationByID is the resolver for the locationById field.
func (r *queryResolver) LocationByID(ctx context.Context, id string) (*model.Locations, error) {
	return locationRouter.Get(ctx, id)
}
