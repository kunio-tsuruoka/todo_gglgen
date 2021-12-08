package dataloader

import (
	"time"
	"todo/dataloader/dataloader_gen"
	"todo/graph/model"
	"gorm.io/gorm"
	"todo/repository"
)

func UsersByIDs(db *gorm.DB) *dataloader_gen.UserLoader {
	return dataloader_gen.NewUserLoader(dataloader_gen.UserLoaderConfig{
		MaxBatch: 100,
		Wait:     1 * time.Millisecond,
		Fetch: func(keys []int) ([]*model.User, []error) {
			repo := repository.NewUserRepository(db)
			users, err := repo.GetAllByUserIDs(keys) 
			if err != nil {
				return nil, []error{err}
			}

			usersByIDs := map[int]*model.User{}
			for _, user := range users {
				usersByIDs[user.ID] = user
			}

			results := make([]*model.User, len(keys))
			for i, id := range keys {
				results[i] = usersByIDs[id]
			}
			return results, nil
		},
	})
}