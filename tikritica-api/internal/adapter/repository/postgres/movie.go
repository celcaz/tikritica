package postgres

import (
	"context"
	"database/sql"

	"github.com/celio/tikritica-api/internal/domain/entity"
	"github.com/celio/tikritica-api/internal/domain/port"
	"github.com/lib/pq"
)

var _ port.MovieRepository = (*MovieRepo)(nil)

type MovieRepo struct {
	db *sql.DB
}

func NewMovieRepo(db *sql.DB) *MovieRepo {
	return &MovieRepo{db: db}
}

func (r *MovieRepo) List(ctx context.Context) ([]entity.Movie, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, slug, title, description, COALESCE(cover_url, ''), COALESCE(release_date::text, ''),
		       COALESCE(director, ''), COALESCE(runtime, 0), genres, average_rating, ratings_count,
		       created_at, updated_at
		FROM movies
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []entity.Movie
	for rows.Next() {
		var m entity.Movie
		if err := rows.Scan(
			&m.ID, &m.Slug, &m.Title, &m.Description, &m.CoverURL, &m.ReleaseDate,
			&m.Director, &m.Runtime, pq.Array(&m.Genres), &m.AverageRating, &m.RatingsCount,
			&m.CreatedAt, &m.UpdatedAt,
		); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, rows.Err()
}

func (r *MovieRepo) GetBySlug(ctx context.Context, slug string) (*entity.Movie, error) {
	var m entity.Movie
	err := r.db.QueryRowContext(ctx, `
		SELECT id, slug, title, description, COALESCE(cover_url, ''), COALESCE(release_date::text, ''),
		       COALESCE(director, ''), COALESCE(runtime, 0), genres, average_rating, ratings_count,
		       created_at, updated_at
		FROM movies
		WHERE slug = $1
	`, slug).Scan(
		&m.ID, &m.Slug, &m.Title, &m.Description, &m.CoverURL, &m.ReleaseDate,
		&m.Director, &m.Runtime, pq.Array(&m.Genres), &m.AverageRating, &m.RatingsCount,
		&m.CreatedAt, &m.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &m, nil
}
