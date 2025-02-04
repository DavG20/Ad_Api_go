package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/router/companyLocationRouter"
	"context"
)

// CreateCompanyLocation is the resolver for the createCompanyLocation field.
func (r *mutationResolver) CreateCompanyLocation(ctx context.Context, input model.CreateCompanyLocationInput) (*model.CompanyLocations, error) {
	return companyLocationRouter.Create(ctx, input)
}

// DeleteCompanyLocation is the resolver for the deleteCompanyLocation field.
func (r *mutationResolver) DeleteCompanyLocation(ctx context.Context, companyID string, locationID string) (*model.DeletionResult, error) {
	return companyLocationRouter.Delete(ctx, companyID, locationID)
}

// CompanyLocationByID is the resolver for the companyLocationById field.
func (r *queryResolver) CompanyLocationByID(ctx context.Context, companyID string, locationID string) (*model.CompanyLocations, error) {
	return companyLocationRouter.Get(ctx, companyID, locationID)
}
