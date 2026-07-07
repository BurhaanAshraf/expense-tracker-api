package repository

import (
	"context"
	"database/sql"

	"github.com/BurhaanAshraf/finance-api/internal/model"
)

type ExpenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) *ExpenseRepository {
	return &ExpenseRepository{
		db: db,
	}
}

func (r *ExpenseRepository) Create(ctx context.Context, expense *model.Expense) error {
	query := `
	INSERT INTO expenses (
		user_id,
		title,
		amount,
		category,
		expense_date,
		notes
	)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	result, err := r.db.ExecContext(
		ctx,
		query,
		expense.UserID,
		expense.Title,
		expense.Amount,
		expense.Category,
		expense.ExpenseDate,
		expense.Notes,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	expense.ID = id

	return nil
}
