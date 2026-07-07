package main

import (
	"log"
	"net/http"

	"github.com/BurhaanAshraf/finance-api/internal/config"
	"github.com/BurhaanAshraf/finance-api/internal/database"
	"github.com/BurhaanAshraf/finance-api/internal/handler"
	"github.com/BurhaanAshraf/finance-api/internal/repository"
	service "github.com/BurhaanAshraf/finance-api/internal/service"
)

func main() {

	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	authHandler := handler.NewAuthHandler(userService)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.HealthHandler)
	mux.HandleFunc("POST /register", authHandler.Register)
	mux.HandleFunc("POST /login", authHandler.Login)

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: mux,
	}

	log.Printf("%s started on http://localhost:%s", cfg.AppName, cfg.AppPort)

	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
