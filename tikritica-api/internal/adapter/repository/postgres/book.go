package postgres

import (
	"context"
	"database/sql"

	"github.com/celio/tikritica-api/internal/domain/entity"
	"github.com/celio/tikritica-api/internal/domain/port"
	"github.com/lib/pq"
)

var _ port.BookRepository = (*BookRepo)(nil)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (r *BookRepo) List(ctx context.Context) ([]entity.Book, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, slug, title, description, COALESCE(cover_url, ''), COALESCE(release_date::text, ''),
		       COALESCE(author, ''), COALESCE(pages, 0), COALESCE(isbn, ''), genres,
		       average_rating, ratings_count, created_at, updated_at
		FROM books
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entity.Book
	for rows.Next() {
		var b entity.Book
		if err := rows.Scan(
			&b.ID, &b.Slug, &b.Title, &b.Description, &b.CoverURL, &b.ReleaseDate,
			&b.Author, &b.Pages, &b.ISBN, pq.Array(&b.Genres),
			&b.AverageRating, &b.RatingsCount, &b.CreatedAt, &b.UpdatedAt,
		); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, rows.Err()
}

func (r *BookRepo) GetBySlug(ctx context.Context, slug string) (*entity.Book, error) {
	var b entity.Book
	err := r.db.QueryRowContext(ctx, `
		SELECT id, slug, title, description, COALESCE(cover_url, ''), COALESCE(release_date::text, ''),
		       COALESCE(author, ''), COALESCE(pages, 0), COALESCE(isbn, ''), genres,
		       average_rating, ratings_count, created_at, updated_at
		FROM books
		WHERE slug = $1
	`, slug).Scan(
		&b.ID, &b.Slug, &b.Title, &b.Description, &b.CoverURL, &b.ReleaseDate,
		&b.Author, &b.Pages, &b.ISBN, pq.Array(&b.Genres),
		&b.AverageRating, &b.RatingsCount, &b.CreatedAt, &b.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &b, nil
}
