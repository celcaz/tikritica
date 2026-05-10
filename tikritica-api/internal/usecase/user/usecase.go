package user

import (
	"context"
	"errors"

	"github.com/celio/tikritica-api/internal/domain/entity"
	"github.com/celio/tikritica-api/internal/domain/port"
)

var ErrNotFound = errors.New("user not found")

type UseCase struct {
	repo port.UserRepository
}

func NewUseCase(repo port.UserRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := uc.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrNotFound
	}
	return user, nil
}
