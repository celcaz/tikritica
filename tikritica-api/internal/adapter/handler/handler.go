package handler

import (
	authuc "github.com/celio/tikritica-api/internal/usecase/auth"
	bookuc "github.com/celio/tikritica-api/internal/usecase/book"
	gameuc "github.com/celio/tikritica-api/internal/usecase/game"
	movieuc "github.com/celio/tikritica-api/internal/usecase/movie"
	seriesuc "github.com/celio/tikritica-api/internal/usecase/series"
	useruc "github.com/celio/tikritica-api/internal/usecase/user"
)

type Handler struct {
	MovieUC  *movieuc.UseCase
	SeriesUC *seriesuc.UseCase
	GameUC   *gameuc.UseCase
	BookUC   *bookuc.UseCase
	AuthUC   *authuc.UseCase
	UserUC   *useruc.UseCase
}

func New(
	movieUC *movieuc.UseCase,
	seriesUC *seriesuc.UseCase,
	gameUC *gameuc.UseCase,
	bookUC *bookuc.UseCase,
	authUC *authuc.UseCase,
	userUC *useruc.UseCase,
) *Handler {
	return &Handler{
		MovieUC:  movieUC,
		SeriesUC: seriesUC,
		GameUC:   gameUC,
		BookUC:   bookUC,
		AuthUC:   authUC,
		UserUC:   userUC,
	}
}
