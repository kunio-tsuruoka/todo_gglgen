package resolver

import (
	"todo/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todoRepository repository.TodoRepository
	userRepository repository.UserRepository
}
func NewResolver(todoRepository repository.TodoRepository, userRepository repository.UserRepository) *Resolver {
	return &Resolver{todoRepository: todoRepository, userRepository: userRepository}
}