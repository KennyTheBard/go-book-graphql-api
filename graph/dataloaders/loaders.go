package dataloaders

import (
	"context"
	"net/http"

	"github.com/graph-gophers/dataloader/v6"
	"gorm.io/gorm"
)

type Loaders struct {
	AuthorLoader *dataloader.Loader
}

func NewLoaders(db *gorm.DB) *Loaders {
	authorReader := &AuthorReader{
		DB: db,
	}
	loaders := &Loaders{
		AuthorLoader: dataloader.NewBatchedLoader(authorReader.GetAuthors),
	}
	return loaders
}

const LoadersKey = "loaders"

func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), LoadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

func GetLoaders(ctx context.Context) *Loaders {
	return ctx.Value(LoadersKey).(*Loaders)
}
