package dataloader

import (
	"time"
	"todo/dataloader/dataloader_gen"
	"todo/graph/model"
	"gorm.io/gorm"
)

func UsersByIDs(db *gorm.DB) *dataloader_gen.UserLoader {
	return dataloader_gen.NewUserLoader(dataloader_gen.UserLoaderConfig{
		MaxBatch: 100,
		Wait:     1 * time.Millisecond,
		Fetch: func(keys []int) ([]*model.User, []error) {

			var records []*model.User
			if err := db.Find(&records, keys).Error; err != nil {
				return nil, []error{err}
			}

			usersByIDs := map[int]*model.User{}
			for _, user := range records {
				usersByIDs[user.ID] = user
			}

			users := make([]*model.User, len(keys))
			for i, id := range keys {
				users[i] = usersByIDs[id]
			}
			return users, nil
		},
	})
}