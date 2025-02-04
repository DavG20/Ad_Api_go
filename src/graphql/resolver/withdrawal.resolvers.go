package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"adtec/backend/src/graphql/generated"
	"adtec/backend/src/graphql/model"
	"adtec/backend/src/router/withdrawalRouter"
	"adtec/backend/src/utils"
	"context"
	"encoding/json"
)

// CreateWithdrawal is the resolver for the createWithdrawal field.
func (r *mutationResolver) CreateWithdrawal(ctx context.Context, input model.CreateWithdrawalInput) (*model.Withdrawals, error) {
	return withdrawalRouter.Create(ctx, input)
}

// UpdateWithdrawal is the resolver for the updateWithdrawal field.
func (r *mutationResolver) UpdateWithdrawal(ctx context.Context, input model.UpdateWithdrawalInput) (*model.Withdrawals, error) {
	return withdrawalRouter.Update(ctx, input)
}

// WithdrawalByID is the resolver for the withdrawalById field.
func (r *queryResolver) WithdrawalByID(ctx context.Context, id string) (*model.Withdrawals, error) {
	return withdrawalRouter.Get(ctx, id)
}

// GetAllWithdrawals is the resolver for the getAllWithdrawals field.
func (r *queryResolver) GetAllWithdrawals(ctx context.Context, withdrawalFilter *model.WithdrawalFilter) ([]*model.Withdrawals, error) {
	return withdrawalRouter.GetAll(ctx, withdrawalFilter)
}

// Extra is the resolver for the extra field.
func (r *withdrawalsResolver) Extra(ctx context.Context, obj *model.Withdrawals) (interface{}, error) {
	if obj != nil && obj.Extra != nil {
		log := utils.ByteConvertor(obj.Extra)
		var response interface{}
		_ = json.Unmarshal(log, &response)
		return response, nil
	}
	return nil, nil
}

// Log is the resolver for the log field.
func (r *withdrawalsResolver) Log(ctx context.Context, obj *model.Withdrawals) (interface{}, error) {
	if obj != nil && obj.Log != nil {
		log := utils.ByteConvertor(obj.Log)
		var response interface{}
		_ = json.Unmarshal(log, &response)
		return response, nil
	}
	return nil, nil
}

// Withdrawals returns generated.WithdrawalsResolver implementation.
func (r *Resolver) Withdrawals() generated.WithdrawalsResolver { return &withdrawalsResolver{r} }

type withdrawalsResolver struct{ *Resolver }
