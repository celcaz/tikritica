package postgres

import (
	"context"
	"database/sql"

	"github.com/celio/tikritica-api/internal/domain/entity"
	"github.com/celio/tikritica-api/internal/domain/port"
)

var _ port.UserRepository = (*UserRepo)(nil)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	var u entity.User
	err := r.db.QueryRowContext(ctx, `
		SELECT id, username, display_name, email, password_hash,
		       avatar_url, bio, followers_count, following_count,
		       reviews_count, lists_count, created_at
		FROM users
		WHERE username = $1
	`, username).Scan(
		&u.ID, &u.Username, &u.DisplayName, &u.Email, &u.PasswordHash,
		&u.AvatarURL, &u.Bio, &u.FollowersCount, &u.FollowingCount,
		&u.ReviewsCount, &u.ListsCount, &u.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	var u entity.User
	err := r.db.QueryRowContext(ctx, `
		SELECT id, username, display_name, email, password_hash,
		       avatar_url, bio, followers_count, following_count,
		       reviews_count, lists_count, created_at
		FROM users
		WHERE email = $1
	`, email).Scan(
		&u.ID, &u.Username, &u.DisplayName, &u.Email, &u.PasswordHash,
		&u.AvatarURL, &u.Bio, &u.FollowersCount, &u.FollowingCount,
		&u.ReviewsCount, &u.ListsCount, &u.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) GetByID(ctx context.Context, id string) (*entity.User, error) {
	var u entity.User
	err := r.db.QueryRowContext(ctx, `
		SELECT id, username, display_name, email, password_hash,
		       avatar_url, bio, followers_count, following_count,
		       reviews_count, lists_count, created_at
		FROM users
		WHERE id = $1
	`, id).Scan(
		&u.ID, &u.Username, &u.DisplayName, &u.Email, &u.PasswordHash,
		&u.AvatarURL, &u.Bio, &u.FollowersCount, &u.FollowingCount,
		&u.ReviewsCount, &u.ListsCount, &u.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepo) Create(ctx context.Context, user *entity.User) error {
	return r.db.QueryRowContext(ctx, `
		INSERT INTO users (username, display_name, email, password_hash)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at
	`, user.Username, user.DisplayName, user.Email, user.PasswordHash,
	).Scan(&user.ID, &user.CreatedAt)
}
