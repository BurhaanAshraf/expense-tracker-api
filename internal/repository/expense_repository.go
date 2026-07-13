package repository

import (
	"context"
	"database/sql"
	"fmt"

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

func (r *ExpenseRepository) GetAll(ctx context.Context, userID int64, category string, page int, limit int, sort string) ([]model.Expense, error) {

	offset := (page - 1) * limit
	sortClause := getSortClause(sort)

	var (
		rows *sql.Rows
		err  error
	)

	if category != "" {

		query := fmt.Sprintf(`
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
		WHERE user_id = ? AND category = ?
		ORDER BY %s
		LIMIT ? OFFSET ?
		`, sortClause)

		rows, err = r.db.QueryContext(
			ctx,
			query,
			userID,
			category,
			limit,
			offset,
		)

	} else {

		query := fmt.Sprintf(`
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
		ORDER BY %s
		LIMIT ? OFFSET ?
		`, sortClause)

		rows, err = r.db.QueryContext(
			ctx,
			query,
			userID,
			limit,
			offset,
		)
	}

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

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return expenses, nil
}
func (r *ExpenseRepository) GetByID(ctx context.Context, userID int64, expenseID int64) (*model.Expense, error) {

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
func (r *ExpenseRepository) Update(ctx context.Context, expense *model.Expense) error {

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
func (r *ExpenseRepository) Delete(ctx context.Context, userID int64, expenseID int64) error {

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
func (r *ExpenseRepository) Dashboard(ctx context.Context, userID int64) (*model.Dashboard, error) {

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
func (r *ExpenseRepository) CategorySummary(ctx context.Context, userID int64) ([]model.CategorySummary, error) {

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
func getSortClause(sort string) string {
	switch sort {
	case "date_asc":
		return "expense_date ASC"
	case "date_desc":
		return "expense_date DESC"
	case "amount_asc":
		return "amount ASC"
	case "amount_desc":
		return "amount DESC"
	case "created_asc":
		return "created_at ASC"
	case "created_desc":
		return "created_at DESC"
	case "category_asc":
		return "category ASC"
	case "category_desc":
		return "category DESC"
	default:
		return "expense_date DESC"
	}
}
