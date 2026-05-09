package handler

import (
	authuc "github.com/celio/tikritica-api/internal/usecase/auth"
	bookuc "github.com/celio/tikritica-api/internal/usecase/book"
	gameuc "github.com/celio/tikritica-api/internal/usecase/game"
	movieuc "github.com/celio/tikritica-api/internal/usecase/movie"
	seriesuc "github.com/celio/tikritica-api/internal/usecase/series"
)

type Handler struct {
	MovieUC  *movieuc.UseCase
	SeriesUC *seriesuc.UseCase
	GameUC   *gameuc.UseCase
	BookUC   *bookuc.UseCase
	AuthUC   *authuc.UseCase
}

func New(
	movieUC *movieuc.UseCase,
	seriesUC *seriesuc.UseCase,
	gameUC *gameuc.UseCase,
	bookUC *bookuc.UseCase,
	authUC *authuc.UseCase,
) *Handler {
	return &Handler{
		MovieUC:  movieUC,
		SeriesUC: seriesUC,
		GameUC:   gameUC,
		BookUC:   bookUC,
		AuthUC:   authUC,
	}
}
