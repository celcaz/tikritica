package port

import (
	"context"

	"github.com/celio/tikritica-api/internal/domain/entity"
)

type ReviewRepository interface {
	GetByID(ctx context.Context, id string) (*entity.Review, error)
	Create(ctx context.Context, review *entity.Review) error
	ListByMediaID(ctx context.Context, mediaID string, mediaType string) ([]entity.Review, error)
	ListByUserID(ctx context.Context, userID string) ([]entity.Review, error)
}
