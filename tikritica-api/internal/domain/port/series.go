package port

import (
	"context"

	"github.com/celio/tikritica-api/internal/domain/entity"
)

type SeriesRepository interface {
	List(ctx context.Context) ([]entity.Series, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Series, error)
}
