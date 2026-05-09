package port

import (
	"context"

	"github.com/celio/tikritica-api/internal/domain/entity"
)

type MovieRepository interface {
	List(ctx context.Context) ([]entity.Movie, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Movie, error)
}
