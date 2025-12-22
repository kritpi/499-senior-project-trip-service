package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
)

func (r *Repository) GetTripsInfoTest(ctx context.Context) ([]domain.Trip, error) {
	// Mock data for trips
	trips := []domain.Trip{
		{
			ID:          uuid.New().String(),
			Name:        "Paris Adventure",
			Destination: "Paris, France",
			StartDate:   time.Date(2025, 3, 15, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2025, 3, 22, 0, 0, 0, 0, time.UTC),
			Status:      "planning",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			Name:        "Tokyo Experience",
			Destination: "Tokyo, Japan",
			StartDate:   time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2025, 4, 15, 0, 0, 0, 0, time.UTC),
			Status:      "planning",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			Name:        "Bali Retreat",
			Destination: "Bali, Indonesia",
			StartDate:   time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC),
			EndDate:     time.Date(2025, 2, 10, 0, 0, 0, 0, time.UTC),
			Status:      "ongoing",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}
	return trips, nil

}
