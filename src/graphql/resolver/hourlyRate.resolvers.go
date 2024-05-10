package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/router/hourlyRateRouter"
	"context"
)

// CreateHourlyRate is the resolver for the createHourlyRate field.
func (r *mutationResolver) CreateHourlyRate(ctx context.Context, input model.CreateHourlyRateInput) (*model.HourlyRates, error) {
	return hourlyRateRouter.Create(ctx, input)
}

// UpdateHourlyRate is the resolver for the updateHourlyRate field.
func (r *mutationResolver) UpdateHourlyRate(ctx context.Context, input model.UpdateHourlyRateInput) (*model.HourlyRates, error) {
	return hourlyRateRouter.Update(ctx, input)
}

// HourlyRateByID is the resolver for the HourlyRateById field.
func (r *queryResolver) HourlyRateByID(ctx context.Context, id string) (*model.HourlyRates, error) {
	return hourlyRateRouter.Get(ctx, id)
}
