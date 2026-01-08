package entity

import (
	"time"

	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
)

type UpsertTripResponse struct {
	TripId int `db:"trip_id"`
}

type Trip struct {
	ID           int       `db:"id"`
	OwnerId      string    `db:"owner_id"`
	TripName     string    `db:"name"`
	Description  string    `db:"description"`
	StartDate    time.Time `db:"start_date"`
	EndDate      time.Time `db:"end_date"`
	MainLocation string    `db:"main_location"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

type GetMemberTripsResponse struct {
	Trips []MemberTrips `db:"trips"`
}

type MemberTrips struct {
	TripId       int           `db:"id"`
	TripName     string        `db:"trip_name"`
	StartDate    time.Time     `db:"start_date"`
	EndDate      time.Time     `db:"end_date"`
	MainLocation string        `db:"main_location"`
	Role         enum.MemberRole `db:"member_role"`
}

type TripMemberWithDetails struct {
	MemberId string        `db:"member_id"`
	Name     string        `db:"name"`
	ImageUrl string        `db:"image_url"`
	Role     enum.MemberRole `db:"member_role"`
}
