package port

import (
	"context"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

type Service interface {
	GetTrips(ctx context.Context) ([]domain.Trip, error)
}
