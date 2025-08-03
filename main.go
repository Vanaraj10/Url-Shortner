package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Vanaraj10/Url-Shortner/config"
	"github.com/Vanaraj10/Url-Shortner/handlers"
	"github.com/Vanaraj10/Url-Shortner/middleware"
	"github.com/Vanaraj10/Url-Shortner/repository"
	"github.com/Vanaraj10/Url-Shortner/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := config.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	userRepo := repository.NewUserPostgres(db)
	urlRepo := repository.NewURLPostgres(db)
	userService := service.NewUserService(userRepo)
	urlService := service.NewURLService(urlRepo)
	userHandler := handlers.NewUserHandler(userService, cfg)
	urlHandler := handlers.NewURLHandler(urlService)

	r := chi.NewRouter()

	r.Post("/register", userHandler.Register)
	r.Post("/login", userHandler.Login)

	r.Group(func(protected chi.Router) {
		protected.Use(middleware.JWTAuth(cfg))
		protected.Post("/shorten", urlHandler.CreateShortURL)
	})

	r.Get("/redirect", urlHandler.Redirect)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the URL Shortener API!"))
	})
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", r)
}
