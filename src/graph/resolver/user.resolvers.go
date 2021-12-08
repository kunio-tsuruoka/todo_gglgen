package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"todo/graph/generated"
	"todo/graph/model"
	"todo/dataloader"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		Name: input.Name,
	}
	return r.userRepository.Save(user)
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input model.DeleteUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	user, err := r.userRepository.FindOne(id)
	if err != nil {
		return nil, errors.New("該当するuserがありません")
	}
	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.userRepository.GetAll()
}

func (r *userResolver) Todos(ctx context.Context, obj *model.User) ([]*model.Todo, error) {
	return dataloader.For(ctx).TodosByUserIDs.Load(obj.ID)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
