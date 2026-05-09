package port

import (
	"context"

	"github.com/celio/tikritica-api/internal/domain/entity"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByID(ctx context.Context, id string) (*entity.User, error)
	Create(ctx context.Context, user *entity.User) error
}
