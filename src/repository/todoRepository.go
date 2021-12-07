package repository 
import (
	"todo/graph/model"
	"gorm.io/gorm"
)
type TodoRepository interface {
	Save(todo *model.Todo) (*model.Todo, error)
	GetAll() ([]*model.Todo, error)
	FindOne(id int) (*model.Todo, error)
	Destroy(id int) (*model.Todo, error)
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{
		db: db,
	} 
}

func (repo *todoRepository)Save(todo *model.Todo) (*model.Todo, error) {
	if err := repo.db.Create(&todo).Error; err != nil {
		return nil, err
	}
	return todo, nil
}

func (repo *todoRepository)GetAll() ([]*model.Todo, error) {
	var todos []*model.Todo
	if err := repo.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
func (repo *todoRepository)FindOne(id int) (*model.Todo, error) {
	var todo *model.Todo
	if err := repo.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return todo, nil
}
func (repo *todoRepository)Destroy(id int) (*model.Todo, error) {
	var todo *model.Todo
	if err := repo.db.Delete(&todo, id).Error; err != nil {
		return nil, err
	}	
	return todo, nil
}