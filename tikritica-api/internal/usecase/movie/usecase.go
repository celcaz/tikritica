package movie

import (
	"context"
	"errors"

	"github.com/celio/tikritica-api/internal/domain/entity"
	"github.com/celio/tikritica-api/internal/domain/port"
)

var ErrNotFound = errors.New("movie not found")

type UseCase struct {
	repo port.MovieRepository
}

func NewUseCase(repo port.MovieRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) ListMovies(ctx context.Context) ([]entity.Movie, error) {
	return uc.repo.List(ctx)
}

func (uc *UseCase) GetMovieBySlug(ctx context.Context, slug string) (*entity.Movie, error) {
	movie, err := uc.repo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	if movie == nil {
		return nil, ErrNotFound
	}
	return movie, nil
}
