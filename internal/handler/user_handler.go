package handler

import (
	"encoding/json"
	"net/http"

	"github.com/BurhaanAshraf/finance-api/internal/middleware"
)

func Me(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int64)

	response := map[string]any{
		"user_id": userID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
