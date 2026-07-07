package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BurhaanAshraf/finance-api/internal/dto"
	"github.com/BurhaanAshraf/finance-api/internal/middleware"
	"github.com/BurhaanAshraf/finance-api/internal/response"
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

func (h *ExpenseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateExpenseRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.BadRequest(w, "Invalid request body")
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
		response.BadRequest(w, err.Error())
		return
	}
	response.Created(w, expense)
}
func (h *ExpenseHandler) GetAll(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	category := r.URL.Query().Get("category")

	page := 1
	limit := 10

	if value := r.URL.Query().Get("page"); value != "" {
		p, err := strconv.Atoi(value)
		if err == nil {
			page = p
		}
	}

	if value := r.URL.Query().Get("limit"); value != "" {
		l, err := strconv.Atoi(value)
		if err == nil {
			limit = l
		}
	}
	sort := r.URL.Query().Get("sort")
	expenses, err := h.expenseService.GetAll(
		r.Context(),
		userID,
		category,
		page,
		limit,
		sort,
	)

	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	response.OK(w, expenses)
}
func (h *ExpenseHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	expenseID, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		response.BadRequest(w, "Invalid expense id")
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	expense, err := h.expenseService.GetByID(
		r.Context(),
		userID,
		expenseID,
	)

	if err != nil {
		response.NotFound(w, err.Error())
		return
	}

	response.OK(w, expense)
}
func (h *ExpenseHandler) Update(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	expenseID, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		response.BadRequest(w, "Invalid expense id")
		return
	}

	var req dto.CreateExpenseRequest

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.BadRequest(w, "Invalid request body")
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
		response.BadRequest(w, err.Error())
		return
	}

	response.OK(w, map[string]string{
		"message": "Expense updated successfully",
	})
}
func (h *ExpenseHandler) Delete(w http.ResponseWriter, r *http.Request) {

	idString := r.PathValue("id")

	expenseID, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		response.BadRequest(w, "Invalid expense id")
		return
	}

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	err = h.expenseService.Delete(
		r.Context(),
		userID,
		expenseID,
	)

	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	response.NoContent(w)
}
func (h *ExpenseHandler) Dashboard(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	dashboard, err := h.expenseService.Dashboard(
		r.Context(),
		userID,
	)

	if err != nil {
		response.InternalServerError(w, err.Error())
		return
	}

	response.OK(w, dashboard)
}
func (h *ExpenseHandler) CategorySummary(w http.ResponseWriter, r *http.Request) {

	userID := r.Context().Value(middleware.UserIDKey).(int64)

	summary, err := h.expenseService.CategorySummary(
		r.Context(),
		userID,
	)

	if err != nil {
		response.InternalServerError(w, "Internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	response.OK(w, summary)
}
