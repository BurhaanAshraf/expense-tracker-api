package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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
func (h *ExpenseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(int64)

	expenses, err := h.expenseService.GetAll(
		r.Context(),
		userID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}
func (h *ExpenseHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	expenseID, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(w, "Invalid expense id", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	expense, err := h.expenseService.GetByID(
		r.Context(),
		userID,
		expenseID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expense)
}
func (h *ExpenseHandler) Update(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	expenseID, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(w, "Invalid expense id", http.StatusBadRequest)
		return
	}

	var req CreateExpenseRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	err = h.expenseService.Update(
		r.Context(),
		userID,
		expenseID,
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

	w.WriteHeader(http.StatusOK)
}
func (h *ExpenseHandler) Delete(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	expenseID, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(w, "Invalid expense id", http.StatusBadRequest)
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	err = h.expenseService.Delete(
		r.Context(),
		userID,
		expenseID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
