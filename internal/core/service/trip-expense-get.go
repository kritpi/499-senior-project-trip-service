package service

import (
	"context"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (s *service) TripExpenseGet(ctx context.Context, in domain.TripExpenseRequest) (*domain.TripExpenseResponse, error) {
	// get trip total amount
	tripExpenseTotalAmount, err := s.repo.GetTripExpenseTotalAmount(ctx, in.TripId)
	if err != nil || tripExpenseTotalAmount == nil{
		log.Errorf("unable to get trip expenses total amount: %+v", err)
		return nil, err
	}

	// get my trip's total split
	myTotalSplit, err := s.repo.GetTripMemberTotalExpenses(ctx, in)
	if err != nil || myTotalSplit == nil {
		log.Errorf("unable to get member total trip's expenses: %+v", err)
		return nil, err
	}

	// get all trip's expense and my spit in each expense
	expenses, err := s.repo.GetTripExpenses(ctx, in)
	if err != nil || expenses == nil {
		log.Errorf("unable to get trip's expenses: %+v", err)
		return nil, err
	}

	tripExpenses := domain.TripExpenseResponse{
		TripId:        in.TripId,
		TotalAmount:   *tripExpenseTotalAmount,
		MyTotalAmount: *myTotalSplit,
		Expenses:      *expenses,
	}
	return &tripExpenses, nil
}
