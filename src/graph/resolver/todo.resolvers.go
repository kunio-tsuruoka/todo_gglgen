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
	return r.todoRepository.Save(todo)
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodo) (*model.Todo, error) {
     return r.todoRepository.Destroy(input.ID)
}

func (r *queryResolver) Todo(ctx context.Context, id int) (*model.Todo, error) {
	return r.todoRepository.FindOne(id)	
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todoRepository.GetAll()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
