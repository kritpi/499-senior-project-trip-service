package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
	"github.com/shopspring/decimal"
)

func (s *service) UpsertExpense(ctx context.Context, in domain.UpsertExpenseRequest) error {
	id := uuid.New()
	// validate trip id
	trip, err := s.repo.GetTripById(ctx, in.TripId)
	if err != nil {
		log.Errorf("unable to get trip: %+v", err)
		return err
	}
	if trip == nil {
		errMsg := fmt.Sprintf("unable to get trip id: %d", in.TripId)
		return errors.New(errMsg)
	}

	// assign id for new expense
	if in.ExpenseId == "" {
		in.ExpenseId = id.String()
	}

	members, err := s.repo.GetTripMembers(ctx, in.TripId)
	if err != nil {
		log.Errorf("unable to get trip's members: %+v", err)
		return nil
	}

	var expenseMemberSplit []domain.ExpenseMember

	switch in.SplitType {
	case enum.ExpenseAllEqual:
		expenseMemberSplit = make([]domain.ExpenseMember, 0, len(members))

		lenMemberDecimal := decimal.NewFromInt(int64(len(members)))
		equalSplit := in.Amount.Div(lenMemberDecimal)

		for _, m := range members {
			expenseMemberSplit = append(expenseMemberSplit, domain.ExpenseMember{
				MemberId: m.MemberId,
				Amount:   &equalSplit,
			})
		}
	case enum.ExpenseSelectedEqual:
		expenseMemberSplit = make([]domain.ExpenseMember, 0, len(in.Participant))

		lenMemberDecimal := decimal.NewFromInt(int64(len(in.Participant)))
		equalSplit := in.Amount.Div(lenMemberDecimal)

		for _, p := range in.Participant {
			expenseMemberSplit = append(expenseMemberSplit, domain.ExpenseMember{
				MemberId: p.MemberId,
				Amount:   &equalSplit,
			})
		}
	case enum.ExpenseCustom:
		expenseMemberSplit = in.Participant
	}

	// upsert expense
	err = s.repo.UpsertExpense(ctx, in)
	if err != nil {
		return err
	}
	
	txErr := s.repo.Transactional(ctx, func(txCtx context.Context) error {

		// delete expense member
		err = s.repo.DeleteExpenseMember(txCtx, in.ExpenseId)
		if err != nil {
			return err
		}
		expenseMemberReq := domain.ExpenseMemberRequest{
			ExpenseId: in.ExpenseId,
			ExpenseMember: expenseMemberSplit,
		}
		
		// batch insert expense member
		err = s.repo.BatchInsertExpenseMember(txCtx, expenseMemberReq)
		if err != nil {
			return err
		}

		return nil
	})
	if txErr != nil {
		errMsg := fmt.Sprintf("error in transactional: %+v", txErr)
		return errors.New(errMsg)
	}

	return nil
}
