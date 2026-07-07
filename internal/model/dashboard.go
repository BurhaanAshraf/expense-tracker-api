package model

type Dashboard struct {
	TotalExpenses     float64 `json:"total_expenses"`
	TotalTransactions int64   `json:"total_transactions"`
	AverageExpense    float64 `json:"average_expense"`
	HighestExpense    float64 `json:"highest_expense"`
}
