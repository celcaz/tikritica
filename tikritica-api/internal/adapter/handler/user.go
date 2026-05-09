package handler

import (
	"net/http"
)

// GetUser godoc
// @Summary      Get user profile
// @Description  Get user profile by username
// @Tags         users
// @Produce      json
// @Param        username   path      string  true  "Username"
// @Success      200        {object}  model.User
// @Failure      404        {object}  response.ErrorResponse
// @Failure      500        {object}  response.ErrorResponse
// @Router       /api/users/{username} [get]
func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	_ = r.PathValue("username")
	w.WriteHeader(http.StatusNotImplemented)
}
