package service

import (
	"context"

	"github.com/BurhaanAshraf/finance-api/internal/model"
	"github.com/BurhaanAshraf/finance-api/internal/repository"
)

type ExpenseService struct {
	expenseRepository *repository.ExpenseRepository
}

func NewExpenseService(expenseRepository *repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{
		expenseRepository: expenseRepository,
	}
}
func (s *ExpenseService) Create(ctx context.Context, userID int64, title string, amount float64, category string, expenseDate string, notes string) (*model.Expense, error) {

	date, err := ParseDate(expenseDate)
	if err != nil {
		return nil, err
	}

	expense := &model.Expense{
		UserID:      userID,
		Title:       title,
		Amount:      amount,
		Category:    category,
		ExpenseDate: date,
		Notes:       notes,
	}

	err = s.expenseRepository.Create(ctx, expense)
	if err != nil {
		return nil, err
	}

	return expense, nil
}
func (s *ExpenseService) GetAll(ctx context.Context, userID int64, category string, page int, limit int, sort string) ([]model.Expense, error) {

	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	return s.expenseRepository.GetAll(
		ctx,
		userID,
		category,
		page,
		limit,
		sort,
	)
}
func (s *ExpenseService) GetByID(ctx context.Context, userID int64, expenseID int64) (*model.Expense, error) {
	return s.expenseRepository.GetByID(
		ctx,
		userID,
		expenseID,
	)
}
func (s *ExpenseService) Update(ctx context.Context, userID int64, expenseID int64, title string, amount float64, category string, expenseDate string, notes string) error {

	date, err := ParseDate(expenseDate)
	if err != nil {
		return err
	}

	expense := &model.Expense{
		ID:          expenseID,
		UserID:      userID,
		Title:       title,
		Amount:      amount,
		Category:    category,
		ExpenseDate: date,
		Notes:       notes,
	}

	return s.expenseRepository.Update(ctx, expense)
}
func (s *ExpenseService) Delete(ctx context.Context, userID int64, expenseID int64) error {

	return s.expenseRepository.Delete(
		ctx,
		userID,
		expenseID,
	)
}
func (s *ExpenseService) Dashboard(ctx context.Context, userID int64) (*model.Dashboard, error) {

	return s.expenseRepository.Dashboard(
		ctx,
		userID,
	)
}
func (s *ExpenseService) CategorySummary(ctx context.Context, userID int64) ([]model.CategorySummary, error) {

	return s.expenseRepository.CategorySummary(
		ctx,
		userID,
	)
}
