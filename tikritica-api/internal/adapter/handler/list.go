package handler

import (
	"net/http"
)

// GetList godoc
// @Summary      Get list by ID
// @Description  Get details of a specific list
// @Tags         lists
// @Produce      json
// @Param        id   path      string  true  "List ID"
// @Success      200  {object}  entity.List
// @Failure      404  {object}  response.ErrorResponse
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/lists/{id} [get]
func (h *Handler) getList(w http.ResponseWriter, r *http.Request) {
	_ = r.PathValue("id")
	w.WriteHeader(http.StatusNotImplemented)
}
