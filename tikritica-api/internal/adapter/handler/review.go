package handler

import (
	"net/http"
)

// GetReview godoc
// @Summary      Get review by ID
// @Description  Get details of a specific review
// @Tags         reviews
// @Produce      json
// @Param        id   path      string  true  "Review ID"
// @Success      200  {object}  model.Review
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/reviews/{id} [get]
func (h *Handler) getReview(w http.ResponseWriter, r *http.Request) {
	_ = r.PathValue("id")
	w.WriteHeader(http.StatusNotImplemented)
}

// CreateReview godoc
// @Summary      Create a new review
// @Description  Create a new review for a media item
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Param        review  body      model.Review  true  "Review data"
// @Success      201     {object}  model.Review
// @Failure      400     {object}  response.ErrorResponse
// @Failure      500     {object}  response.ErrorResponse
// @Router       /api/reviews [post]
func (h *Handler) createReview(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
