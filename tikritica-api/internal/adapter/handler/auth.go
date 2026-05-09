package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/celio/tikritica-api/internal/infra/middleware"
	"github.com/celio/tikritica-api/internal/usecase/auth"
	"github.com/celio/tikritica-api/pkg/response"
)

type registerRequest struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		response.Error(w, http.StatusBadRequest, "username, email, and password are required")
		return
	}

	if len(req.Password) < 8 {
		response.Error(w, http.StatusBadRequest, "password must be at least 8 characters")
		return
	}

	displayName := req.DisplayName
	if displayName == "" {
		displayName = req.Username
	}

	result, err := h.AuthUC.Register(r.Context(), auth.RegisterInput{
		Username:    req.Username,
		DisplayName: displayName,
		Email:       req.Email,
		Password:    req.Password,
	})
	if err != nil {
		switch {
		case errors.Is(err, auth.ErrEmailTaken):
			response.Error(w, http.StatusConflict, "email already in use")
		case errors.Is(err, auth.ErrUsernameTaken):
			response.Error(w, http.StatusConflict, "username already in use")
		default:
			response.Error(w, http.StatusInternalServerError, "internal server error")
		}
		return
	}

	response.JSON(w, http.StatusCreated, result)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if req.Email == "" || req.Password == "" {
		response.Error(w, http.StatusBadRequest, "email and password are required")
		return
	}

	result, err := h.AuthUC.Login(r.Context(), auth.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		if errors.Is(err, auth.ErrInvalidCreds) {
			response.Error(w, http.StatusUnauthorized, "invalid email or password")
			return
		}
		response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	response.JSON(w, http.StatusOK, result)
}

func (h *Handler) refresh(w http.ResponseWriter, r *http.Request) {
	userID := middleware.GetUserID(r.Context())
	if userID == "" {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	result, err := h.AuthUC.Refresh(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	response.JSON(w, http.StatusOK, result)
}
