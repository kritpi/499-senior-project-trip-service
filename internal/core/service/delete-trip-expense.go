package service

import (
	"context"
	"errors"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
)

func (s *service) DeleteExpense(ctx context.Context, in domain.DeleteTripExpenseRequest) error {
	role, err := s.repo.GetTripMemberRole(ctx, in.MemberId, in.TripId)
	if err != nil || role == nil || *role == enum.MemberRoleViewer {

		return errors.New("insufficient permissions - only owner or editor can delete expense")
	}

	return s.repo.Transactional(ctx, func(txCtx context.Context) error {
		if err := s.repo.DeleteExpenseMember(txCtx, in.ExpenseId); err != nil {
			return err
		}

		if err := s.repo.DeleteExpense(txCtx, in.ExpenseId); err != nil {
			return err
		}

		return nil
	})
}
