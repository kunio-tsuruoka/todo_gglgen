package dataloader
import (
	"context"
	"net/http"
	"gorm.io/gorm"
	"todo/dataloader/dataloader_gen"

)
const loadersKey = "dataloaders"

type Loaders struct {
	TodosByUserIDs *dataloader_gen.TodoLoader
	UsersByIDs *dataloader_gen.UserLoader
}
func DataLoaderMiddleware(db *gorm.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), loadersKey, &Loaders{
			UsersByIDs: UsersByIDs(db),
		    TodosByUserIDs:TodosByUserIDs(db),
		})
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
