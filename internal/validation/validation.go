package validation

import (
	"errors"
	"net/mail"
	"strings"
	"time"

	"github.com/BurhaanAshraf/finance-api/internal/dto"
)

func ValidateRegister(req dto.RegisterRequest) error {
	if strings.TrimSpace(req.Name) == "" {
		return errors.New("name is required")
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return errors.New("invalid email")
	}

	if len(req.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

func ValidateLogin(req dto.LoginRequest) error {
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return errors.New("invalid email")
	}

	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	return nil
}

func ValidateExpense(req dto.CreateExpenseRequest) error {
	if strings.TrimSpace(req.Title) == "" {
		return errors.New("title is required")
	}

	if req.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}

	if strings.TrimSpace(req.Category) == "" {
		return errors.New("category is required")
	}

	if _, err := time.Parse("2006-01-02", req.ExpenseDate); err != nil {
		return errors.New("invalid expense date")
	}

	return nil
}
