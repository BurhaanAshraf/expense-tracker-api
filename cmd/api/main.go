package main

import (
	"log"
	"net/http"

	"github.com/BurhaanAshraf/finance-api/internal/config"
	"github.com/BurhaanAshraf/finance-api/internal/database"
	"github.com/BurhaanAshraf/finance-api/internal/handlers"
)

func main() {

	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handlers.HealthHandler)

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
