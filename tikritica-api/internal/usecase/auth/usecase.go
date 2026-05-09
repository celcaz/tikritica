package auth

import (
	"context"
	"errors"
	"time"

	"github.com/celio/tikritica-api/internal/domain/entity"
	"github.com/celio/tikritica-api/internal/domain/port"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailTaken    = errors.New("email already in use")
	ErrUsernameTaken = errors.New("username already in use")
	ErrInvalidCreds  = errors.New("invalid email or password")
)

type UseCase struct {
	repo      port.UserRepository
	jwtSecret []byte
}

func NewUseCase(repo port.UserRepository, jwtSecret string) *UseCase {
	return &UseCase{
		repo:      repo,
		jwtSecret: []byte(jwtSecret),
	}
}

type RegisterInput struct {
	Username    string
	DisplayName string
	Email       string
	Password    string
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  *entity.User `json:"user"`
}

func (uc *UseCase) Register(ctx context.Context, input RegisterInput) (*AuthResponse, error) {
	existing, err := uc.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrEmailTaken
	}

	existing, err = uc.repo.GetByUsername(ctx, input.Username)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrUsernameTaken
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Username:     input.Username,
		DisplayName:  input.DisplayName,
		Email:        input.Email,
		PasswordHash: string(hash),
	}

	if err := uc.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	token, err := uc.generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{Token: token, User: user}, nil
}

type LoginInput struct {
	Email    string
	Password string
}

func (uc *UseCase) Login(ctx context.Context, input LoginInput) (*AuthResponse, error) {
	user, err := uc.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCreds
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return nil, ErrInvalidCreds
	}

	token, err := uc.generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{Token: token, User: user}, nil
}

func (uc *UseCase) Refresh(ctx context.Context, userID string) (*AuthResponse, error) {
	user, err := uc.repo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, ErrInvalidCreds
	}

	token, err := uc.generateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{Token: token, User: user}, nil
}

func (uc *UseCase) generateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(72 * time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(uc.jwtSecret)
}
