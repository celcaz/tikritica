package postgres

import (
	"context"
	"database/sql"

	"github.com/celio/tikritica-api/internal/domain/entity"
	"github.com/celio/tikritica-api/internal/domain/port"
	"github.com/lib/pq"
)

var _ port.GameRepository = (*GameRepo)(nil)

type GameRepo struct {
	db *sql.DB
}

func NewGameRepo(db *sql.DB) *GameRepo {
	return &GameRepo{db: db}
}

func (r *GameRepo) List(ctx context.Context) ([]entity.Game, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, slug, title, description, COALESCE(cover_url, ''), COALESCE(release_date::text, ''),
		       COALESCE(developer, ''), COALESCE(publisher, ''), platforms, genres,
		       average_rating, ratings_count, created_at, updated_at
		FROM games
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var games []entity.Game
	for rows.Next() {
		var g entity.Game
		if err := rows.Scan(
			&g.ID, &g.Slug, &g.Title, &g.Description, &g.CoverURL, &g.ReleaseDate,
			&g.Developer, &g.Publisher, pq.Array(&g.Platforms), pq.Array(&g.Genres),
			&g.AverageRating, &g.RatingsCount, &g.CreatedAt, &g.UpdatedAt,
		); err != nil {
			return nil, err
		}
		games = append(games, g)
	}

	return games, rows.Err()
}

func (r *GameRepo) GetBySlug(ctx context.Context, slug string) (*entity.Game, error) {
	var g entity.Game
	err := r.db.QueryRowContext(ctx, `
		SELECT id, slug, title, description, COALESCE(cover_url, ''), COALESCE(release_date::text, ''),
		       COALESCE(developer, ''), COALESCE(publisher, ''), platforms, genres,
		       average_rating, ratings_count, created_at, updated_at
		FROM games
		WHERE slug = $1
	`, slug).Scan(
		&g.ID, &g.Slug, &g.Title, &g.Description, &g.CoverURL, &g.ReleaseDate,
		&g.Developer, &g.Publisher, pq.Array(&g.Platforms), pq.Array(&g.Genres),
		&g.AverageRating, &g.RatingsCount, &g.CreatedAt, &g.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &g, nil
}
