package domain

import (
	"time"

	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
	"github.com/kritpi/499-senior-project-trip-service/internal/repository/entity"
)

// Trip represents a trip entity
type Trip struct {
	ID           int       `json:"id"`
	OwnerId      string    `json:"owner_id"`
	TripName     string    `json:"name"`
	Description  string    `json:"description"`
	StartDate    time.Time `json:"start_date"`
	EndDate      time.Time `json:"end_date"`
	MainLocation string    `json:"main_location"`
	ImageUrl     string    `json:"image_url"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (t Trip) FromEntity(e entity.Trip) *Trip {
	return &Trip{
		ID:           e.ID,
		OwnerId:      e.OwnerId,
		TripName:     e.TripName,
		Description:  e.Description,
		StartDate:    e.StartDate,
		EndDate:      e.EndDate,
		MainLocation: e.MainLocation,
		ImageUrl:     e.ImageUrl,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
	}
}

// GetTripsRequest is the request for getting trips
type GetTripsRequest struct {
}

// GetTripsResponse is the response for getting trips
type GetTripsResponse struct {
	Trips []Trip `json:"trips"`
}

// Upsert trip request
type UpsertTripRequest struct {
	ID           *int       `json:"id"`
	OwnerId      string     `json:"owner_id"`
	TripName     string     `json:"trip_name"`
	Desciption   *string    `json:"description"`
	StartDate    *time.Time `json:"start_date"` // 2025-07-11
	EndDate      *time.Time `json:"end_date"`   // 2025-07-11
	MainLocation *string    `json:"main_location"`
	ImageUrl     *string    `json:"image_url"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type UpsertTripResponse struct {
	TripId int `json:"trip_id"`
}

func (t UpsertTripResponse) FromEntity(e entity.UpsertTripResponse) *UpsertTripResponse {
	return &UpsertTripResponse{
		TripId: t.TripId,
	}
}

type TripMember struct {
	MemberId string          `json:"member_id"`
	Role     enum.MemberRole `json:"role"`
}

type CreateTripMemberRequest struct {
	TripId    int             `json:"trip_id"`
	MemberId  string          `json:"member_id"`
	Role      enum.MemberRole `json:"role"`
	CreatedAt time.Time       `json:"created_at"`
}

type GetMemberTripsRequest struct {
	MemberId string `json:"member_id"`
}

type GetMemberTripsResponse struct {
	Trips []MemberTrips `json:"trips"`
}

type MemberTrips struct {
	TripId       int             `json:"trip_id"`
	TripName     string          `json:"trip_name"`
	StartDate    *time.Time      `json:"start_date"`
	EndDate      *time.Time      `json:"end_date"`
	MainLocation *string         `json:"main_location"`
	ImageUrl     *string         `json:"image_url"`
	Role         enum.MemberRole `json:"role"`
}

func (t GetMemberTripsResponse) FromEntity(e entity.GetMemberTripsResponse) *GetMemberTripsResponse {
	trips := make([]MemberTrips, len(e.Trips))
	for i, t := range e.Trips {
		trips[i] = MemberTrips{
			TripId:       t.TripId,
			TripName:     t.TripName,
			StartDate:    t.StartDate,
			EndDate:      t.EndDate,
			MainLocation: t.MainLocation,
			ImageUrl:     t.ImageUrl,
			Role:         t.Role,
		}
	}
	return &GetMemberTripsResponse{
		Trips: trips,
	}
}

// TripMemberDetails represents a member's details in a trip
type TripMemberDetails struct {
	MemberId string          `json:"member_id"`
	Email    string          `json:"email"`
	Name     string          `json:"name"`
	ImageUrl string          `json:"image_url"`
	Role     enum.MemberRole `json:"role"`
}

func (t TripMemberDetails) FromEntity(e entity.TripMemberWithDetails) TripMemberDetails {
	return TripMemberDetails{
		MemberId: e.MemberId,
		Email:    e.Email,
		Name:     e.Name,
		ImageUrl: e.ImageUrl,
		Role:     e.Role,
	}
}

// GetTripByIdRequest is the request for getting a trip by ID
type GetTripByIdRequest struct {
	TripId   int    `json:"trip_id"`
	MemberId string `json:"member_id"`
}

// GetTripByIdResponse is the response for getting a trip by ID
type GetTripByIdResponse struct {
	Trip    *Trip               `json:"trip"`
	Role    enum.MemberRole     `json:"role"`
	Members []TripMemberDetails `json:"members"`
}

type TripInvitationRequest struct {
	MemberId string          `json:"member_id"`
	TripId   int             `json:"trip_id"`
	Email    string          `json:"email"`
	Role     enum.MemberRole `json:"role"`
}

type TripInvitationResponse struct {
	Email string          `json:"email"`
	Role  enum.MemberRole `json:"role"`
}

type DeleteTripMemberRequest struct {
	TripId int    `json:"trip_id"`
	Email  string `json:"email"`
}

type GetTripMemberRoleRequest struct {
	TripId   int    `json:"trip_id"`
	MemberId string `json:"member_id"`
}

type GetTripMemberRoleResponse struct {
	TripId   int             `json:"trip_id"`
	MemberId string          `json:"member_id"`
	Role     enum.MemberRole `json:"role"`
}

type GetTripMembersRequest struct {
	TripId   int    `json:"trip_id"`
	MemberId string `json:"member_id"`
}

type GetTripMembersResponse struct {
	Members []TripMemberDetails `json:"members"`
}
