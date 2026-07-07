package dto

type ExpenseQuery struct {
	Category string
	Page     int
	Limit    int
	Sort     string
}
