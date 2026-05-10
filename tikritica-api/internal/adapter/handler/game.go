package handler

import (
	"errors"
	"net/http"

	gameuc "github.com/celio/tikritica-api/internal/usecase/game"
	"github.com/celio/tikritica-api/pkg/response"
)

// ListGames godoc
// @Summary      List all games
// @Description  Get a list of all games
// @Tags         games
// @Produce      json
// @Success      200  {array}   entity.Game
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/games [get]
func (h *Handler) listGames(w http.ResponseWriter, r *http.Request) {
	games, err := h.GameUC.ListGames(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to list games")
		return
	}

	response.JSON(w, http.StatusOK, games)
}

// GetGame godoc
// @Summary      Get game by slug
// @Description  Get details of a specific game by its slug
// @Tags         games
// @Produce      json
// @Param        slug   path      string  true  "Game Slug"
// @Success      200    {object}  entity.Game
// @Failure      404    {object}  response.ErrorResponse
// @Failure      500    {object}  response.ErrorResponse
// @Router       /api/games/{slug} [get]
func (h *Handler) getGame(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	game, err := h.GameUC.GetGameBySlug(r.Context(), slug)
	if err != nil {
		if errors.Is(err, gameuc.ErrNotFound) {
			response.Error(w, http.StatusNotFound, "game not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to get game")
		return
	}

	response.JSON(w, http.StatusOK, game)
}
