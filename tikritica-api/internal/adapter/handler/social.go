package handler

import (
	"net/http"
)

// FollowUser godoc
// @Summary      Follow a user
// @Description  Follow another user
// @Tags         social
// @Produce      json
// @Param        userId   path      string  true  "User ID"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /api/follow/{userId} [post]
func (h *Handler) followUser(w http.ResponseWriter, r *http.Request) {
	_ = r.PathValue("userId")
	w.WriteHeader(http.StatusNotImplemented)
}

// UnfollowUser godoc
// @Summary      Unfollow a user
// @Description  Unfollow a user you are currently following
// @Tags         social
// @Produce      json
// @Param        userId   path      string  true  "User ID"
// @Success      200      {object}  map[string]string
// @Failure      400      {object}  response.ErrorResponse
// @Failure      500      {object}  response.ErrorResponse
// @Router       /api/follow/{userId} [delete]
func (h *Handler) unfollowUser(w http.ResponseWriter, r *http.Request) {
	_ = r.PathValue("userId")
	w.WriteHeader(http.StatusNotImplemented)
}

// GetFeed godoc
// @Summary      Get activity feed
// @Description  Get feed of activities from followed users
// @Tags         social
// @Produce      json
// @Success      200  {array}   map[string]interface{}
// @Failure      500  {object}  response.ErrorResponse
// @Router       /api/feed [get]
func (h *Handler) getFeed(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
