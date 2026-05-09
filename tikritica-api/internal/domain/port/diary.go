package port

import (
	"context"

	"github.com/celio/tikritica-api/internal/domain/entity"
)

type DiaryRepository interface {
	ListByUserID(ctx context.Context, userID string) ([]entity.DiaryEntry, error)
	Create(ctx context.Context, entry *entity.DiaryEntry) error
}
