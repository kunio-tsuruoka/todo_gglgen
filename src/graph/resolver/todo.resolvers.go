package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"todo/graph/generated"
	"todo/graph/model"
	"todo/dataloader"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Title:  input.Title,
		Desc:   input.Desc,
		UserID: input.UserID,
	}
	return r.todoRepository.Save(todo)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodo) (*model.Todo, error) {
	return r.todoRepository.Destroy(input.ID)
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input model.UpdateTodo) (bool, error) {
	_, err := r.todoRepository.FindOne(input.ID)
	if err != nil {
		return false, errors.New("該当するtodoがありません")
	}
	r.todoRepository.Update(input)
	return true, nil
}

func (r *queryResolver) Todo(ctx context.Context, id int) (*model.Todo, error) {
	todo, err := r.todoRepository.FindOne(id)
	if err != nil {
		return nil, errors.New("該当するtodoがありません")
	}
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todoRepository.GetAll()
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return dataloader.For(ctx).UsersByIDs.Load(obj.UserID)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
