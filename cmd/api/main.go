package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurhaanAshraf/finance-api/internal/config"
	"github.com/BurhaanAshraf/finance-api/internal/database"
	"github.com/BurhaanAshraf/finance-api/internal/handler"
	"github.com/BurhaanAshraf/finance-api/internal/middleware"
	"github.com/BurhaanAshraf/finance-api/internal/repository"
	service "github.com/BurhaanAshraf/finance-api/internal/service"
)

func main() {

	cfg := config.Load()
	jwtMiddlware := middleware.JWTMiddleware(cfg.JWTSecret)

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(userService)

	expenseRepository := repository.NewExpenseRepository(db)
	expenseService := service.NewExpenseService(expenseRepository)
	expenseHandler := handler.NewExpenseHandler(expenseService)

	mux := http.NewServeMux()
	loggingMiddleware := middleware.LoggingMiddleware
	mux.HandleFunc("GET /health", handler.HealthHandler)
	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)
	mux.Handle("GET /me", jwtMiddlware(http.HandlerFunc(handler.Me)))
	mux.Handle("POST /expenses", jwtMiddlware(http.HandlerFunc(expenseHandler.Create)))
	mux.Handle("GET /expenses", jwtMiddlware(http.HandlerFunc(expenseHandler.GetAll)))
	mux.Handle("GET /expenses/{id}", jwtMiddlware(http.HandlerFunc(expenseHandler.GetByID)))
	mux.Handle("PUT /expenses/{id}", jwtMiddlware(http.HandlerFunc(expenseHandler.Update)))
	mux.Handle("DELETE /expenses/{id}", jwtMiddlware(http.HandlerFunc(expenseHandler.Delete)))
	mux.Handle("GET /dashboard", jwtMiddlware(http.HandlerFunc(expenseHandler.Dashboard)))
	mux.Handle("GET /dashboard/categories", jwtMiddlware(http.HandlerFunc(expenseHandler.CategorySummary)))

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: loggingMiddleware(mux),
	}

	go func() {
		log.Printf("%s started on http://localhost:%s", cfg.AppName, cfg.AppPort)

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(
		quit,
		os.Interrupt,
		syscall.SIGTERM,
	)

	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}

	if err := db.Close(); err != nil {
		log.Println("Error closing database:", err)
	}

	log.Println("Server stopped gracefully.")

}
