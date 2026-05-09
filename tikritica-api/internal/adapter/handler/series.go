package handler

import (
	"errors"
	"net/http"

	seriesuc "github.com/celio/tikritica-api/internal/usecase/series"
	"github.com/celio/tikritica-api/pkg/response"
)

// ListSeries godoc
// @Summary      List all series
// @Description  Get a list of all series
// @Tags         series
// @Produce      json
// @Success      200  {array}   model.Series
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/series [get]
func (h *Handler) listSeries(w http.ResponseWriter, r *http.Request) {
	series, err := h.SeriesUC.ListSeries(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "failed to list series")
		return
	}

	response.JSON(w, http.StatusOK, series)
}

// GetSeries godoc
// @Summary      Get series by slug
// @Description  Get details of a specific series by its slug
// @Tags         series
// @Produce      json
// @Param        slug   path      string  true  "Series Slug"
// @Success      200    {object}  model.Series
// @Failure      404    {object}  response.ErrorResponse
// @Failure      500    {object}  response.ErrorResponse
// @Router       /api/series/{slug} [get]
func (h *Handler) getSeries(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	s, err := h.SeriesUC.GetSeriesBySlug(r.Context(), slug)
	if err != nil {
		if errors.Is(err, seriesuc.ErrNotFound) {
			response.Error(w, http.StatusNotFound, "series not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to get series")
		return
	}

	response.JSON(w, http.StatusOK, s)
}
