package port

import (
	"context"

	"github.com/celio/tikritica-api/internal/domain/entity"
)

type ListRepository interface {
	GetByID(ctx context.Context, id string) (*entity.List, error)
	Create(ctx context.Context, list *entity.List) error
	ListByUserID(ctx context.Context, userID string) ([]entity.List, error)
}
