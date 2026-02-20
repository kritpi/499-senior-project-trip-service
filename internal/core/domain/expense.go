package domain

import (
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
	"github.com/shopspring/decimal"
)

type UpsertExpenseRequest struct {
	TripId      int             `json:"trip_id"`
	ExpenseId   string          `json:"expense_id"`
	Title       string          `json:"title"`
	Amount      decimal.Decimal `json:"amount"`
	CreatedBy   string          `json:"created_by"` //uuid
	SplitType   enum.SplitType  `json:"split_type"`
	ImageUrl    *string         `json:"image_url"`
	Participant []ExpenseMember `json:"participant"`
}

type ExpenseMember struct {
	MemberId string           `json:"member_id"`
	Amount   *decimal.Decimal `json:"amount"`
}

type ExpenseMemberRequest struct {
	ExpenseId     string
	ExpenseMember []ExpenseMember
}

type ExpenseMemberResponse struct {
	MemberId string          `json:"member_id"`
	Name     string          `json:"name"`
	ImageUrl string          `json:"image_url"`
	Amount   decimal.Decimal `json:"amount"`
}

type TripExpenseRequest struct {
	MemberId string `json:"member_id"`
	TripId   int    `json:"trip_id"`
}

type TripExpenseResponse struct {
	TripId        int               `json:"trip_id"`
	TotalAmount   decimal.Decimal   `json:"total_amount"`
	MyTotalAmount decimal.Decimal   `json:"my_total_amount"`
	Expenses      []ExpenseResponse `json:"expenses"`
}

type ExpenseResponse struct {
	ExpenseId   string                  `json:"expense_id"`
	Title       string                  `json:"title"`
	Amount      decimal.Decimal         `json:"amount"`
	MyShared    decimal.Decimal         `json:"my_shared"`
	CreatedBy   string                  `json:"created_by"` //name
	ImageUrl    *string                 `json:"image_url"`
	SplitType   enum.SplitType          `json:"split_type"`
	Participant []ExpenseMemberResponse `json:"participant"`
}

type GetTripMemberExpense struct {
	MemberId string
	TripId   int
}

type DeleteTripExpenseRequest struct {
	MemberId  string `json:"member_id"`
	TripId    int    `json:"trip_id"`
	ExpenseId string `json:"expense_id"`
}
