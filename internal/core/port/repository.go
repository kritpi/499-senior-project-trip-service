package port

import (
	"context"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

type Repository interface {
	GetTripsInfo(ctx context.Context) ([]domain.Trip, error)
	GetTripsInfoTest(ctx context.Context) ([]domain.Trip, error)

	GetMemberByEmail(ctx context.Context, email string) (*domain.Member, error)
	CreateMember(ctx context.Context, member domain.Member) error
}
