package dto

import (
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
)

type SocketContext struct {
	MemberId string          `json:"member_id"`
	TripId   int             `json:"trip_id"`
	TripDate string          `json:"trip_date"`
	Role     enum.MemberRole `json:"role"`
	Joined   bool            `json:"joined"`
}

func (s SocketContext) ToDomain() *domain.SocketContext {
	return &domain.SocketContext{
		MemberId: s.MemberId,
		TripId:   s.TripId,
		TripDate: s.TripDate,
		Role:     s.Role,
		Joined:   s.Joined,
	}
}
