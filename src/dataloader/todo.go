package dataloader

import (
	"time"
	"todo/dataloader/dataloader_gen"
	"todo/graph/model"
	"gorm.io/gorm"
)

func TodosByUserIDs(db *gorm.DB) *dataloader_gen.TodoLoader {
	return dataloader_gen.NewTodoLoader(dataloader_gen.TodoLoaderConfig{
		MaxBatch: 100,
		Wait:     1 * time.Millisecond,
		Fetch: func(keys []int) ([][]*model.Todo, []error) {

			var records []*model.Todo
			if err := db.Where("user_id IN ?", keys).Find(&records).Error; err != nil {
				return nil, []error{err}
			}

			todosByIDs := map[int][]*model.Todo{}
			for _, todo := range records {
			todosByIDs[todo.UserID] = append(todosByIDs[todo.UserID], todo)
			}

			todos := make([][]*model.Todo, len(keys))
			for i, id := range keys {
				todos[i] =todosByIDs[id]
			}
			return todos, nil
		},
	})
}