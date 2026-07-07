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
func (r *ExpenseRepository) GetAllByUserID(ctx context.Context, userID int64) ([]model.Expense, error) {
	query := `
	SELECT
		id,
		user_id,
		title,
		amount,
		category,
		expense_date,
		notes,
		created_at,
		updated_at
	FROM expenses
	WHERE user_id = ?
	ORDER BY expense_date DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []model.Expense

	for rows.Next() {
		var expense model.Expense

		err := rows.Scan(
			&expense.ID,
			&expense.UserID,
			&expense.Title,
			&expense.Amount,
			&expense.Category,
			&expense.ExpenseDate,
			&expense.Notes,
			&expense.CreatedAt,
			&expense.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		expenses = append(expenses, expense)
	}

	return expenses, nil
}
func (r *ExpenseRepository) GetByID(
	ctx context.Context,
	userID int64,
	expenseID int64,
) (*model.Expense, error) {

	query := `
	SELECT
		id,
		user_id,
		title,
		amount,
		category,
		expense_date,
		notes,
		created_at,
		updated_at
	FROM expenses
	WHERE id = ? AND user_id = ?
	`

	expense := &model.Expense{}

	err := r.db.QueryRowContext(
		ctx,
		query,
		expenseID,
		userID,
	).Scan(
		&expense.ID,
		&expense.UserID,
		&expense.Title,
		&expense.Amount,
		&expense.Category,
		&expense.ExpenseDate,
		&expense.Notes,
		&expense.CreatedAt,
		&expense.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return expense, nil
}
func (r *ExpenseRepository) Update(
	ctx context.Context,
	expense *model.Expense,
) error {

	query := `
	UPDATE expenses
	SET
		title = ?,
		amount = ?,
		category = ?,
		expense_date = ?,
		notes = ?
	WHERE id = ? AND user_id = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		expense.Title,
		expense.Amount,
		expense.Category,
		expense.ExpenseDate,
		expense.Notes,
		expense.ID,
		expense.UserID,
	)

	return err
}
func (r *ExpenseRepository) Delete(
	ctx context.Context,
	userID int64,
	expenseID int64,
) error {

	query := `
	DELETE FROM expenses
	WHERE id = ? AND user_id = ?
	`

	_, err := r.db.ExecContext(
		ctx,
		query,
		expenseID,
		userID,
	)

	return err
}
func (r *ExpenseRepository) Dashboard(
	ctx context.Context,
	userID int64,
) (*model.Dashboard, error) {

	query := `
	SELECT
		IFNULL(SUM(amount),0),
		COUNT(*),
		IFNULL(AVG(amount),0),
		IFNULL(MAX(amount),0)
	FROM expenses
	WHERE user_id = ?
	`

	dashboard := &model.Dashboard{}

	err := r.db.QueryRowContext(
		ctx,
		query,
		userID,
	).Scan(
		&dashboard.TotalExpenses,
		&dashboard.TotalTransactions,
		&dashboard.AverageExpense,
		&dashboard.HighestExpense,
	)

	if err != nil {
		return nil, err
	}

	return dashboard, nil
}
func (r *ExpenseRepository) CategorySummary(
	ctx context.Context,
	userID int64,
) ([]model.CategorySummary, error) {

	query := `
	SELECT
		category,
		SUM(amount)
	FROM expenses
	WHERE user_id = ?
	GROUP BY category
	ORDER BY SUM(amount) DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []model.CategorySummary

	for rows.Next() {

		var summary model.CategorySummary

		err := rows.Scan(
			&summary.Category,
			&summary.Amount,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, summary)
	}

	return result, nil
}
