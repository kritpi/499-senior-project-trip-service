package port

import (
	"context"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

type Repository interface {
	// Transactional
	Transactional(ctx context.Context, fn func(txCtx context.Context) error) error

	// Auth
	GetMemberByEmail(ctx context.Context, email string) (*domain.Member, error)
	CreateMember(ctx context.Context, member domain.Member) (id string, err error)
	GetMemberById(ctx context.Context, id string) (*domain.Member, error)

	// Trip
	UpsertTrip(ctx context.Context, in domain.UpsertTripRequest) (*domain.UpsertTripResponse, error)
	GetTripById(ctx context.Context, tripId int) (*domain.Trip, error)
	CheckTripMembership(ctx context.Context, tripId int, memberId string) (bool, error)
	GetTripMembers(ctx context.Context, tripId int) ([]domain.TripMemberDetails, error)
	BatchCreateTripMember(ctx context.Context, in domain.BatchCreateTripMemberRequest) error
	GetMemberTrips(ctx context.Context, in domain.GetMemberTripsRequest) (*domain.GetMemberTripsResponse, error)
	CheckExistingMember(ctx context.Context, checkExistingMember []string) (*[]domain.GetExistingMemberResp, error)
}
