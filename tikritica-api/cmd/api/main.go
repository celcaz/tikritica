package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/celio/tikritica-api/internal/adapter/handler"
	"github.com/celio/tikritica-api/internal/adapter/repository/postgres"
	"github.com/celio/tikritica-api/internal/infra/config"
	"github.com/celio/tikritica-api/internal/infra/database"
	"github.com/celio/tikritica-api/internal/infra/middleware"
	authuc "github.com/celio/tikritica-api/internal/usecase/auth"
	bookuc "github.com/celio/tikritica-api/internal/usecase/book"
	gameuc "github.com/celio/tikritica-api/internal/usecase/game"
	movieuc "github.com/celio/tikritica-api/internal/usecase/movie"
	seriesuc "github.com/celio/tikritica-api/internal/usecase/series"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Tikritica API
// @version         1.0
// @description     API REST for Tikritica platform.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cfg := config.Load()

	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Connected to database")

	// Repositories (adapters)
	movieRepo := postgres.NewMovieRepo(db)
	seriesRepo := postgres.NewSeriesRepo(db)
	gameRepo := postgres.NewGameRepo(db)
	bookRepo := postgres.NewBookRepo(db)
	userRepo := postgres.NewUserRepo(db)

	// Use cases
	movieUC := movieuc.NewUseCase(movieRepo)
	seriesUC := seriesuc.NewUseCase(seriesRepo)
	gameUC := gameuc.NewUseCase(gameRepo)
	bookUC := bookuc.NewUseCase(bookRepo)
	authUC := authuc.NewUseCase(userRepo, cfg.JWTSecret)

	// Handler
	h := handler.New(movieUC, seriesUC, gameUC, bookUC, authUC)

	mux := http.NewServeMux()
	h.RegisterRoutes(mux)

	// Protected routes (require auth)
	protectedMux := http.NewServeMux()
	h.RegisterProtectedRoutes(protectedMux)
	mux.Handle("/api/auth/refresh", middleware.Auth(cfg.JWTSecret)(protectedMux))

	// Swagger UI
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	stack := middleware.CORS(middleware.Logging(mux))

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on %s", addr)
	log.Fatal(http.ListenAndServe(addr, stack))
}
