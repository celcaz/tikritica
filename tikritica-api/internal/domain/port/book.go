package port

import (
	"context"

	"github.com/celio/tikritica-api/internal/domain/entity"
)

type BookRepository interface {
	List(ctx context.Context) ([]entity.Book, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Book, error)
}
