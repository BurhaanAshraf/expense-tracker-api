package handler

import (
	"encoding/json"
	"net/http"

	"github.com/BurhaanAshraf/finance-api/internal/middleware"
	"github.com/BurhaanAshraf/finance-api/internal/service"
)

type ExpenseHandler struct {
	expenseService *service.ExpenseService
}

func NewExpenseHandler(expenseService *service.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{
		expenseService: expenseService,
	}
}

type CreateExpenseRequest struct {
	Title       string  `json:"title"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	ExpenseDate string  `json:"expense_date"`
	Notes       string  `json:"notes"`
}

func (h *ExpenseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateExpenseRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	expense, err := h.expenseService.Create(
		r.Context(),
		userID,
		req.Title,
		req.Amount,
		req.Category,
		req.ExpenseDate,
		req.Notes,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(expense)
}
