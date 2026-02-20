package port

import (
	"context"
	"time"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

type Service interface {
	GoogleAuth(ctx context.Context, idToken domain.GoogleIdToken) (*domain.GoogleAuthResponse, error)

	// Trip
	UpsertTrip(ctx context.Context, in domain.UpsertTripRequest) (*domain.UpsertTripResponse, error)
	GetMemberTrips(ctx context.Context, in domain.GetMemberTripsRequest) (*domain.GetMemberTripsResponse, error)
	GetTripById(ctx context.Context, in domain.GetTripByIdRequest) (*domain.GetTripByIdResponse, error)
	InviteMember(ctx context.Context, in domain.TripInvitationRequest) (*domain.TripInvitationResponse, error)
	DeleteInvitedMember(ctx context.Context, in domain.DeleteTripMemberRequest) error
	GetTripMemberRole(ctx context.Context, in domain.GetTripMemberRoleRequest) (*domain.GetTripMemberRoleResponse, error)

	// Storage
	UploadImage(ctx context.Context, in domain.UploadImageRequest) (*domain.UploadImageResponse, error)

	// Activity
	SocketActivityJoin(ctx context.Context, in domain.ActivityJoinRequest) (*domain.ActivityJoinResponse, error)
	SocketActivityUpsert(ctx context.Context, in domain.ActivityUpsertRequest) error
	ActivityUpsertProcess(ctx context.Context, in domain.ActivityUpsertProcessRequest) error
	ActivityUpsertBroadcast(ctx context.Context, tripId int, date time.Time) (*domain.ActivityJoinResponse, error)

	// Expense
	UpsertExpense(ctx context.Context, in domain.UpsertExpenseRequest) error
	TripExpenseGet(ctx context.Context, in domain.TripExpenseRequest) (*domain.TripExpenseResponse, error)
	DeleteExpense(ctx context.Context, in domain.DeleteTripExpenseRequest) error
}
