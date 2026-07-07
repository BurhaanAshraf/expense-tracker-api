package dto

type CreateExpenseRequest struct {
	Title       string  `json:"title"`
	Amount      float64 `json:"amount"`
	Category    string  `json:"category"`
	ExpenseDate string  `json:"expense_date"`
	Notes       string  `json:"notes"`
}
