package model

import "time"

type Expense struct {
	ID          int64
	UserID      int64
	Title       string
	Amount      float64
	Category    string
	ExpenseDate time.Time
	Notes       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
