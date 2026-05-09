package game

import (
	"context"
	"errors"

	"github.com/celio/tikritica-api/internal/domain/entity"
	"github.com/celio/tikritica-api/internal/domain/port"
)

var ErrNotFound = errors.New("game not found")

type UseCase struct {
	repo port.GameRepository
}

func NewUseCase(repo port.GameRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) ListGames(ctx context.Context) ([]entity.Game, error) {
	return uc.repo.List(ctx)
}

func (uc *UseCase) GetGameBySlug(ctx context.Context, slug string) (*entity.Game, error) {
	g, err := uc.repo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	if g == nil {
		return nil, ErrNotFound
	}
	return g, nil
}
