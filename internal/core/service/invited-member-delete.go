package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (s *service) DeleteInvitedMember(ctx context.Context, in domain.DeleteTripMemberRequest) error {
	member, err := s.repo.CheckExistingMember(ctx, in.Email)
	if err != nil {
		log.Errorf("cannot get member: %+v", err)
		return err
	}
	if member == nil {
		errMsg := fmt.Sprintf("member is not exists")
		return errors.New(errMsg)
	}

	err = s.repo.DeleteInvitedMember(ctx, member.ID, in.TripId)
	if err != nil {
		log.Errorf("unable to delete member: %+v", err)
		return err
	}
	return nil
}
