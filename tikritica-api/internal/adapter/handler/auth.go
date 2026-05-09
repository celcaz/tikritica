package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/celio/tikritica-api/internal/domain/entity"
	"github.com/celio/tikritica-api/internal/infra/middleware"
	"github.com/celio/tikritica-api/internal/usecase/auth"
	"github.com/celio/tikritica-api/pkg/response"
)

const authCookieName = "authToken"

var authCookieMaxAge = int((72 * time.Hour).Seconds())

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

type authUserResponse struct {
	User *entity.User `json:"user"`
}

type logoutResponse struct {
	Message string `json:"message"`
}

func setAuthCookie(w http.ResponseWriter, token string) {
	// Step 1: Set HttpOnly auth cookie to prevent JS access (XSS mitigation).
	// Step 2: SameSite=Lax reduces CSRF for same-site requests.
	// Step 3: Secure=false for local dev; enable Secure in production.
	http.SetCookie(w, &http.Cookie{
		Name:     authCookieName,
		Value:    token,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   authCookieMaxAge,
	})
}

func clearAuthCookie(w http.ResponseWriter) {
	// Step 1: Expire the cookie immediately on logout.
	// Step 2: Keep flags consistent with login cookie settings.
	http.SetCookie(w, &http.Cookie{
		Name:     authCookieName,
		Value:    "",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Now().Add(-1 * time.Hour),
	})
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	// Step 1: Decode input.
	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Step 2: Validate required fields.
	if req.Username == "" || req.Email == "" || req.Password == "" {
		response.Error(w, http.StatusBadRequest, "username, email, and password are required")
		return
	}

	// Step 3: Validate password strength.
	if len(req.Password) < 8 {
		response.Error(w, http.StatusBadRequest, "password must be at least 8 characters")
		return
	}

	// Step 4: Normalize display name.
	displayName := req.DisplayName
	if displayName == "" {
		displayName = req.Username
	}

	// Step 5: Call use case to create user and token.
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

	// Step 6: Store token in HttpOnly cookie.
	setAuthCookie(w, result.Token)

	// Step 7: Return user payload (cookie holds the token).
	response.JSON(w, http.StatusCreated, authUserResponse{User: result.User})
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	// Step 1: Decode input.
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Step 2: Validate required fields.
	if req.Email == "" || req.Password == "" {
		response.Error(w, http.StatusBadRequest, "email and password are required")
		return
	}

	// Step 3: Authenticate user and generate token.
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

	// Step 4: Store token in HttpOnly cookie.
	setAuthCookie(w, result.Token)

	// Step 5: Return user payload only.
	response.JSON(w, http.StatusOK, authUserResponse{User: result.User})
}

func (h *Handler) refresh(w http.ResponseWriter, r *http.Request) {
	// Step 1: Read user ID from auth middleware context.
	userID := middleware.GetUserID(r.Context())
	if userID == "" {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	// Step 2: Generate a fresh token for the user.
	result, err := h.AuthUC.Refresh(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}

	// Step 3: Rotate the auth cookie with the new token.
	setAuthCookie(w, result.Token)

	// Step 4: Return user payload only.
	response.JSON(w, http.StatusOK, authUserResponse{User: result.User})
}

func (h *Handler) logout(w http.ResponseWriter, r *http.Request) {
	// Step 1: Clear the auth cookie on logout.
	clearAuthCookie(w)

	// Step 2: Return a simple confirmation.
	response.JSON(w, http.StatusOK, logoutResponse{Message: "logged out"})
}

func (h *Handler) me(w http.ResponseWriter, r *http.Request) {
	// Step 1: Read user ID from auth middleware context.
	userID := middleware.GetUserID(r.Context())
	if userID == "" {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	// Step 2: Load user data.
	user, err := h.AuthUC.GetUser(r.Context(), userID)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if user == nil {
		response.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	// Step 3: Return user payload.
	response.JSON(w, http.StatusOK, authUserResponse{User: user})
}
