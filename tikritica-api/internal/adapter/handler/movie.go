package handler

import (
	"errors"
	"net/http"

	movieuc "github.com/celio/tikritica-api/internal/usecase/movie"
	"github.com/celio/tikritica-api/pkg/response"
)

// ListMovies godoc
// @Summary      List all movies
// @Description  Get a list of all movies
// @Tags         movies
// @Produce      json
// @Success      200  {array}   model.Movie
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/movies [get]
func (h *Handler) listMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.MovieUC.ListMovies(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to list movies")
		return
	}

	response.JSON(w, http.StatusOK, movies)
}

// GetMovie godoc
// @Summary      Get movie by slug
// @Description  Get details of a specific movie by its slug
// @Tags         movies
// @Produce      json
// @Param        slug   path      string  true  "Movie Slug"
// @Success      200    {object}  model.Movie
// @Failure      404    {object}  response.ErrorResponse
// @Failure      500    {object}  response.ErrorResponse
// @Router       /api/movies/{slug} [get]
func (h *Handler) getMovie(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	movie, err := h.MovieUC.GetMovieBySlug(r.Context(), slug)
	if err != nil {
		if errors.Is(err, movieuc.ErrNotFound) {
			response.Error(w, http.StatusNotFound, "movie not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to get movie")
		return
	}

	response.JSON(w, http.StatusOK, movie)
}
