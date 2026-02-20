package port

import (
	"context"
	"time"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
	"github.com/shopspring/decimal"
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
	CreateTripMember(ctx context.Context, in domain.CreateTripMemberRequest) error
	GetMemberTrips(ctx context.Context, in domain.GetMemberTripsRequest) (*domain.GetMemberTripsResponse, error)
	CheckExistingMember(ctx context.Context, email string) (*domain.GetExistingMemberResp, error)
	GetTripMemberRole(ctx context.Context, memberId string, tripId int) (*enum.MemberRole, error)
	DeleteInvitedMember(ctx context.Context, memberId string, tripId int) error
	GetActivities(ctx context.Context, in domain.GetActivitiesOfTheDayRequest) (*domain.GetActivitiesOfTheDayResponse, error)
	BatchUpsertActivities(ctx context.Context, in []domain.Activity) error
	BatchDeleteActivities(ctx context.Context, in []domain.Activity) error

	// Expense
	UpsertExpense(ctx context.Context, in domain.UpsertExpenseRequest) error
	DeleteExpenseMember(ctx context.Context, expenseId string) error
	DeleteExpense(ctx context.Context, expenseId string) error
	BatchInsertExpenseMember(ctx context.Context, in domain.ExpenseMemberRequest) error
	GetTripExpenseTotalAmount(ctx context.Context, tripId int) (*decimal.Decimal, error)
	GetTripMemberTotalExpenses(ctx context.Context, in domain.TripExpenseRequest) (*decimal.Decimal, error)
	GetTripExpenses(ctx context.Context, in domain.TripExpenseRequest) (*[]domain.ExpenseResponse, error)
}

type RedisRepository interface {
	EnqueueActivitiesUpsert(ctx context.Context, in domain.ActivityUpsertProcessRequest) error
	EnqueueActivitiesBroadcast(ctx context.Context, tripId int, date time.Time) error
}
