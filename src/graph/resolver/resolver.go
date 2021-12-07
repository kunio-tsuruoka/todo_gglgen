package resolver
import 	(
	"todo/repository"
)


// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
    todoRepository repository.TodoRepository
}
func NewResolver(todoRepository repository.TodoRepository) *Resolver {
	return &Resolver{
		todoRepository: todoRepository,
	}
}