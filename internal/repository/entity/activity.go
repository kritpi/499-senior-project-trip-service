package entity

import (
	"time"

	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
)

type Activity struct {
	ID               string                 `db:"id"`
	TripId           int                    `db:"trip_id"`
	ActivityDate     time.Time              `db:"activity_date"`
	StartTime        *time.Time             `db:"start_time"`
	EndTime          *time.Time             `db:"end_time"`
	Note             *string                `db:"note"`
	Description      *string                `db:"description"`
	ActivityLocation *Location              `db:"activity_location"`
	Category         *enum.ActivityCategory `db:"category"`
	Rank             int                    `db:"rank"`
	CreatedAt        time.Time              `db:"created_at"`
	UpdatedAt        time.Time              `db:"updated_at"`
}

type Location struct {
	Name    string  `db:"name"`
	Address string  `db:"address"`
	Lat     float64 `db:"lat"`
	Lng     float64 `db:"lng"`
}
