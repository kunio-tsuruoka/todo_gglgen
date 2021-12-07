package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"todo/graph/generated"
	"todo/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Title: input.Title,
		Desc:  input.Desc,
	}
	if err := r.DB.Create(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodo) (*model.Todo, error) {
	var todo *model.Todo
	if err := r.DB.Delete(&todo, input.ID).Error; err != nil {
		return nil, err
	}	
	return todo, nil
}

func (r *queryResolver) Todo(ctx context.Context, id string) (*model.Todo, error) {
	var todo *model.Todo
	if err := r.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := r.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
