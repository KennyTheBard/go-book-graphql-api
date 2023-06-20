package dataloaders

import (
	"context"
	"fmt"

	"github.com/KennyTheBard/go-books-graphql-api/db"
	"github.com/graph-gophers/dataloader/v6"
	"gorm.io/gorm"
)

type AuthorReader struct {
	DB *gorm.DB
}

func (r *AuthorReader) GetAuthors(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// get authors
	var authors []db.Author
	if err := r.DB.Find(&authors, keys).Error; err != nil {
		panic(err)
	}

	// sort them by id
	authorMap := map[string]*db.Author{}
	for _, author := range authors {
		authorMap[fmt.Sprintf("%v", author.ID)] = &author
	}

	// convert authors to dataloader result
	result := make([]*dataloader.Result, len(keys))
	for index, authorKey := range keys {
		author, ok := authorMap[authorKey.String()]
		if ok {
			result[index] = &dataloader.Result{
				Data:  author,
				Error: nil,
			}
		} else {
			result[index] = &dataloader.Result{
				Data:  nil,
				Error: fmt.Errorf("User not found %s", authorKey.String()),
			}
		}
	}
	return result
}
