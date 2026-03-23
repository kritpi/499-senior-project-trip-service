package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
)

func (s *service) UpsertTrip(ctx context.Context, in domain.UpsertTripRequest) (*domain.UpsertTripResponse, error) {
	now := time.Now().Local()

	member, err := s.repo.GetMemberById(ctx, in.OwnerId)
	if err != nil || member == nil {
		log.Errorf("member id not found: %+v", err)
		return nil, err
	}

	if in.StartDate != nil && in.EndDate != nil {
		// validate if end date is not before start date
		if in.EndDate.Before(*in.StartDate) {
			return nil, fmt.Errorf("end date cannot be before start date")
		}
	}

	// Insert CreatedAt for a new trip
	if in.ID == nil {
		in.CreatedAt = now
	}
	in.UpdatedAt = now

	var resp *domain.UpsertTripResponse

	err = s.repo.Transactional(ctx, func(txCtx context.Context) error {
		var txErr error
		resp, txErr = s.repo.UpsertTrip(txCtx, in)
		if txErr != nil {
			log.Errorf("unable to insert/update trip: %+v", txErr)
			return txErr
		}

		// Checking if trip is existed. If not, insert trip member (OWNER)
		if in.ID == nil {
			createReq := domain.CreateTripMemberRequest{
				TripId:    resp.TripId,
				MemberId:  member.ID,
				Role:      enum.MemberRoleOwner,
				CreatedAt: now,
			}
			txErr = s.repo.CreateTripMember(txCtx, createReq)
			if txErr != nil {
				log.Errorf("unable to insert trip member: %+v", txErr)
				return txErr
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
