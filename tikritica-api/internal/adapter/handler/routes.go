package handler

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", healthCheck)

	// Auth
	mux.HandleFunc("POST /api/auth/register", h.register)
	mux.HandleFunc("POST /api/auth/login", h.login)

	// Movies
	mux.HandleFunc("GET /api/movies", h.listMovies)
	mux.HandleFunc("GET /api/movies/{slug}", h.getMovie)

	// Series
	mux.HandleFunc("GET /api/series", h.listSeries)
	mux.HandleFunc("GET /api/series/{slug}", h.getSeries)

	// Games
	mux.HandleFunc("GET /api/games", h.listGames)
	mux.HandleFunc("GET /api/games/{slug}", h.getGame)

	// Books
	mux.HandleFunc("GET /api/books", h.listBooks)
	mux.HandleFunc("GET /api/books/{slug}", h.getBook)

	// Reviews
	mux.HandleFunc("GET /api/reviews/{id}", h.getReview)
	mux.HandleFunc("POST /api/reviews", h.createReview)

	// Users
	mux.HandleFunc("GET /api/users/{username}", h.getUser)

	// Lists
	mux.HandleFunc("GET /api/lists/{id}", h.getList)

	// Social
	mux.HandleFunc("POST /api/follow/{userId}", h.followUser)
	mux.HandleFunc("DELETE /api/follow/{userId}", h.unfollowUser)
	mux.HandleFunc("GET /api/feed", h.getFeed)

	// Diary
	mux.HandleFunc("GET /api/diary", h.getDiary)
	mux.HandleFunc("POST /api/diary", h.createDiaryEntry)
}

func (h *Handler) RegisterProtectedRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/auth/refresh", h.refresh)
	mux.HandleFunc("POST /api/auth/logout", h.logout)
	mux.HandleFunc("GET /api/auth/me", h.me)
}
