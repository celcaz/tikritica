package port

import (
	"context"

	"github.com/celio/tikritica-api/internal/domain/entity"
)

type GameRepository interface {
	List(ctx context.Context) ([]entity.Game, error)
	GetBySlug(ctx context.Context, slug string) (*entity.Game, error)
}
