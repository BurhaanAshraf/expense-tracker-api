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

func (s *ExpenseService) Create(
	ctx context.Context,
	userID int64,
	title string,
	amount float64,
	category string,
	expenseDate string,
	notes string,
) (*model.Expense, error) {

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
