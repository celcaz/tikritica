package book

import (
	"context"
	"errors"

	"github.com/celio/tikritica-api/internal/domain/entity"
	"github.com/celio/tikritica-api/internal/domain/port"
)

var ErrNotFound = errors.New("book not found")

type UseCase struct {
	repo port.BookRepository
}

func NewUseCase(repo port.BookRepository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) ListBooks(ctx context.Context) ([]entity.Book, error) {
	return uc.repo.List(ctx)
}

func (uc *UseCase) GetBookBySlug(ctx context.Context, slug string) (*entity.Book, error) {
	b, err := uc.repo.GetBySlug(ctx, slug)
	if err != nil {
		return nil, err
	}
	if b == nil {
		return nil, ErrNotFound
	}
	return b, nil
}
