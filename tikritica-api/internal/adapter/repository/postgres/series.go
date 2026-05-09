package postgres

import (
	"context"
	"database/sql"

	"github.com/celio/tikritica-api/internal/domain/entity"
	"github.com/celio/tikritica-api/internal/domain/port"
	"github.com/lib/pq"
)

var _ port.SeriesRepository = (*SeriesRepo)(nil)

type SeriesRepo struct {
	db *sql.DB
}

func NewSeriesRepo(db *sql.DB) *SeriesRepo {
	return &SeriesRepo{db: db}
}

func (r *SeriesRepo) List(ctx context.Context) ([]entity.Series, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, slug, title, description, COALESCE(cover_url, ''), COALESCE(release_date::text, ''),
		       COALESCE(creator, ''), seasons, episodes, genres, average_rating, ratings_count,
		       created_at, updated_at
		FROM series
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var series []entity.Series
	for rows.Next() {
		var s entity.Series
		if err := rows.Scan(
			&s.ID, &s.Slug, &s.Title, &s.Description, &s.CoverURL, &s.ReleaseDate,
			&s.Creator, &s.Seasons, &s.Episodes, pq.Array(&s.Genres), &s.AverageRating, &s.RatingsCount,
			&s.CreatedAt, &s.UpdatedAt,
		); err != nil {
			return nil, err
		}
		series = append(series, s)
	}

	return series, rows.Err()
}

func (r *SeriesRepo) GetBySlug(ctx context.Context, slug string) (*entity.Series, error) {
	var s entity.Series
	err := r.db.QueryRowContext(ctx, `
		SELECT id, slug, title, description, COALESCE(cover_url, ''), COALESCE(release_date::text, ''),
		       COALESCE(creator, ''), seasons, episodes, genres, average_rating, ratings_count,
		       created_at, updated_at
		FROM series
		WHERE slug = $1
	`, slug).Scan(
		&s.ID, &s.Slug, &s.Title, &s.Description, &s.CoverURL, &s.ReleaseDate,
		&s.Creator, &s.Seasons, &s.Episodes, pq.Array(&s.Genres), &s.AverageRating, &s.RatingsCount,
		&s.CreatedAt, &s.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}
