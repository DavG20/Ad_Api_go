package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/router/categoryRouter"
	"context"
)

// CreateCategory is the resolver for the createCategory field.
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.CreateCategoryInput) (*model.Categories, error) {
	return categoryRouter.Create(ctx, input)
}

// CategoryByID is the resolver for the categoryById field.
func (r *queryResolver) CategoryByID(ctx context.Context, id string) (*model.Categories, error) {
	return categoryRouter.Get(ctx, id)
}

// GetAllCategories is the resolver for the getAllCategories field.
func (r *queryResolver) GetAllCategories(ctx context.Context, categoryFilter *model.CategoryFilter) ([]*model.Categories, error) {
	return categoryRouter.GetAll(ctx, categoryFilter)
}
