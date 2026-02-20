package model

import (
	"time"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

type ActivityUpsertProcessRequest struct {
	TripId     int        `json:"trip_id"`
	Date       time.Time  `json:"trip_date"`
	Activities []domain.Activity `json:"activities"`
}
