package port

import (
	"context"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

type Service interface {
	GoogleAuth(ctx context.Context, idToken domain.GoogleIdToken) (*domain.GoogleAuthResponse, error)

	// Trip
	UpsertTrip(ctx context.Context, in domain.UpsertTripRequest) (*domain.UpsertTripResponse, error)
	GetMemberTrips(ctx context.Context, in domain.GetMemberTripsRequest) (*domain.GetMemberTripsResponse, error)
	GetTripById(ctx context.Context, in domain.GetTripByIdRequest) (*domain.GetTripByIdResponse, error)
	InviteMember(ctx context.Context, in domain.TripInvitationRequest) (*domain.InviteMemberResponse, error)
}
