package series

import (
	"context"
	"errors"

	"github.com/celio/tikritica-api/internal/domain/entity"
	"github.com/celio/tikritica-api/internal/domain/port"
)

var ErrNotFound = errors.New("series not found")

type UseCase struct {
	repo port.SeriesRepository
}

func NewUseCase(repo port.SeriesRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) ListSeries(ctx context.Context) ([]entity.Series, error) {
	return uc.repo.List(ctx)
}

func (uc *UseCase) GetSeriesBySlug(ctx context.Context, slug string) (*entity.Series, error) {
	s, err := uc.repo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, ErrNotFound
	}
	return s, nil
}
