package handler

import (
	"errors"
	"net/http"

	useruc "github.com/celio/tikritica-api/internal/usecase/user"
	"github.com/celio/tikritica-api/pkg/response"
)

// GetUser godoc
// @Summary      Get user profile
// @Description  Get user profile by username
// @Tags         users
// @Produce      json
// @Param        username   path      string  true  "Username"
// @Success      200        {object}  entity.User
// @Failure      404        {object}  response.ErrorResponse
// @Failure      500        {object}  response.ErrorResponse
// @Router       /api/users/{username} [get]
func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	user, err := h.UserUC.GetByUsername(r.Context(), username)
	if err != nil {
		if errors.Is(err, useruc.ErrNotFound) {
			response.Error(w, http.StatusNotFound, "user not found")
			return
		}
		response.Error(w, http.StatusInternalServerError, "failed to get user")
		return
	}

	response.JSON(w, http.StatusOK, user)
}
