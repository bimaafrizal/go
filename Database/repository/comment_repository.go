package repository

import (
	"belajar-database/entity"
	"context"
)

type CommentsRepository interface {
	Insert(ctx context.Context, comment entity.Comments) (entity.Comments, error)
	FindById(ctx context.Context, id int32) (entity.Comments, error)
	FindAll(ctx context.Context) ([]entity.Comments, error)
	Update(ctx context.Context, comment entity.Comments, id int32) (entity.Comments, error)
}
