package dto

import (
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
	"github.com/shopspring/decimal"
)

type UpsertExpenseRequest struct {
	TripId    int     `json:"trip_id"`
	ExpenseId string  `json:"expense_id"`
	Title     string  `json:"title"`
	Amount    float64 `json:"amount"`
	// CreatedBy   string          `json:"created_by"` //uuid
	SplitType   enum.SplitType  `json:"split_type"`
	ImageUrl    *string         `json:"image_url"`
	Participant []ExpenseMember `json:"participant"`
}

type ExpenseMember struct {
	MemberId string   `json:"member_id"`
	Amount   *float64 `json:"amount"`
}

func (e UpsertExpenseRequest) ToDomain(memberId string) *domain.UpsertExpenseRequest {
	participant := make([]domain.ExpenseMember, len(e.Participant))
	for i, p := range e.Participant {
		amount := decimal.NewFromFloat(*p.Amount)
		participant[i] = domain.ExpenseMember{
			MemberId: p.MemberId,
			Amount:   &amount,
		}
	}
	return &domain.UpsertExpenseRequest{
		TripId:      e.TripId,
		Title:       e.Title,
		ExpenseId:   e.ExpenseId,
		Amount:      decimal.NewFromFloat(e.Amount),
		CreatedBy:   memberId,
		SplitType:   e.SplitType,
		ImageUrl:    e.ImageUrl,
		Participant: participant,
	}
}

type TripExpenseResponse struct {
	TripId        int               `json:"trip_id"`
	TotalAmount   float64           `json:"total_amount"`
	MyTotalAmount float64           `json:"my_total_amount"`
	Expenses      []ExpenseResponse `json:"expenses"`
}

type ExpenseResponse struct {
	ExpenseId   string                  `json:"expense_id"`
	Title       string                  `json:"title"`
	Amount      float64                 `json:"amount"`
	MyShared    float64                 `json:"my_shared"`
	CreatedBy   string                  `json:"created_by"` //name
	ImageUrl    *string                 `json:"image_url"`
	SplitType   enum.SplitType          `json:"split_type"`
	Participant []ExpenseMemberResponse `json:"participant"`
}

type ExpenseMemberResponse struct {
	MemberId string  `json:"member_id"`
	Name     string  `json:"name"`
	ImageUrl string  `json:"image_url"`
	Amount   float64 `json:"amount"`
}

func (t TripExpenseResponse) FromDomain(dm *domain.TripExpenseResponse) *TripExpenseResponse {
	expenses := make([]ExpenseResponse, len(dm.Expenses))
	for i, e := range dm.Expenses {
		expenseMember := make([]ExpenseMemberResponse, len(e.Participant))
		for j, m := range e.Participant {
			expenseMember[j] = ExpenseMemberResponse{
				MemberId: m.MemberId,
				Name:     m.Name,
				ImageUrl: m.ImageUrl,
				Amount:   m.Amount.InexactFloat64(),
			}
		}

		expenses[i] = ExpenseResponse{
			ExpenseId:   e.ExpenseId,
			Title:       e.Title,
			Amount:      e.Amount.InexactFloat64(),
			MyShared:    e.MyShared.InexactFloat64(),
			CreatedBy:   e.CreatedBy,
			ImageUrl:    e.ImageUrl,
			SplitType:   e.SplitType,
			Participant: expenseMember,
		}
	}
	return &TripExpenseResponse{
		TripId:        dm.TripId,
		TotalAmount:   dm.TotalAmount.InexactFloat64(),
		MyTotalAmount: dm.MyTotalAmount.InexactFloat64(),
		Expenses:      expenses,
	}
}

type TripExpenseRequest struct {
	MemberId string `json:"member_id"`
	TripId   int    `json:"trip_id"`
}

func (t TripExpenseRequest) ToDomain() *domain.TripExpenseRequest {
	return &domain.TripExpenseRequest{
		MemberId: t.MemberId,
		TripId:   t.TripId,
	}
}

type DeleteTripExpenseRequest struct {
	TripId    int    `json:"trip_id"`
	ExpenseId string `json:"expense_id"`
}

func (d DeleteTripExpenseRequest) ToDomain(memberId string) *domain.DeleteTripExpenseRequest {
	return &domain.DeleteTripExpenseRequest{
		MemberId:  memberId,
		TripId:    d.TripId,
		ExpenseId: d.ExpenseId,
	}
}
