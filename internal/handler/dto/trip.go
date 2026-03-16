package dto

import (
	"time"

	"github.com/gofiber/fiber/v2/log"

	"github.com/kritpi/499-senior-project-trip-service/internal/core/domain"
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
	"github.com/kritpi/499-senior-project-trip-service/shared/utils"
)

// Trip represents a trip entity
type Trip struct {
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
	ID *int `json:"trip_id"`
	// OwnerId      string  `json:"owner_id"`
	TripName     string  `json:"trip_name"`
	Desciption   *string `json:"description"`
	StartDate    *string `json:"start_date"` // 2025-07-11
	EndDate      *string `json:"end_date"`   // 2025-07-11
	MainLocation *string `json:"main_location"`
	ImageUrl     *string `json:"image_url"`
}

func (t UpsertTripRequest) ToDomain(memberId string) *domain.UpsertTripRequest {
	var startDate *time.Time
	if t.StartDate != nil {
		var err error
		startDate, err = utils.DateStringToStartDate(*t.StartDate, utils.DATE_FORMAT)
		if err != nil {
			log.Errorf("error parsing date string: %+v", err)
			startDate = nil
		}
	}

	var endDate *time.Time
	if t.EndDate != nil {
		var err error
		endDate, err = utils.DateStringToStartDate(*t.EndDate, utils.DATE_FORMAT)
		if err != nil {
			log.Errorf("error parsing date string: %+v", err)
			endDate = nil
		}
	}

	return &domain.UpsertTripRequest{
		ID:           t.ID,
		OwnerId:      memberId,
		TripName:     t.TripName,
		Desciption:   t.Desciption,
		StartDate:    startDate,
		EndDate:      endDate,
		MainLocation: t.MainLocation,
		ImageUrl:     t.ImageUrl,
	}
}

type UpsertTripResponse struct {
	TripId int `json:"trip_id"`
}

func (t UpsertTripResponse) FromDomain(dm *domain.UpsertTripResponse) *UpsertTripResponse {
	return &UpsertTripResponse{
		TripId: dm.TripId,
	}
}

type TripMember struct {
	MemberId string          `json:"member_id"`
	Role     enum.MemberRole `json:"role"`
}

type CreateTripMemberRequest struct {
	TripId    int          `json:"trip_id"`
	Members   []TripMember `json:"members"`
	CreatedAt time.Time    `json:"created_at"`
}

type GetMemberTripsRequest struct {
	MemberId string `json:"member_id"`
}

func (t GetMemberTripsRequest) ToDomain(memberId string) *domain.GetMemberTripsRequest {
	return &domain.GetMemberTripsRequest{
		MemberId: memberId,
	}
}

type GetMemberTripsResponse struct {
	Trips []MemberTrips `json:"trips"`
}

type MemberTrips struct {
	TripId       int             `json:"trip_id"`
	TripName     string          `json:"trip_name"`
	StartDate    *string         `json:"start_date"`
	EndDate      *string         `json:"end_date"`
	MainLocation *string         `json:"main_location"`
	ImageUrl     *string         `json:"image_url"`
	Role         enum.MemberRole `json:"role"`
}

func (t GetMemberTripsResponse) FromDomain(dm *domain.GetMemberTripsResponse) *GetMemberTripsResponse {
	trip := make([]MemberTrips, len(dm.Trips))
	for i, t := range dm.Trips {
		var startDate *string
		if t.StartDate != nil {
			sd := utils.DateTimeToDateString(t.StartDate) // Corrected: pass t.StartDate (which is *time.Time)
			startDate = &sd
		}

		var endDate *string
		if t.EndDate != nil {
			ed := utils.DateTimeToDateString(t.EndDate) // Corrected: pass t.EndDate (which is *time.Time)
			endDate = &ed
		}

		trip[i] = MemberTrips{
			TripId:       t.TripId,
			TripName:     t.TripName,
			StartDate:    startDate,
			EndDate:      endDate,
			MainLocation: t.MainLocation,
			ImageUrl:     t.ImageUrl,
			Role:         t.Role,
		}
	}

	return &GetMemberTripsResponse{
		Trips: trip,
	}
}

// GetTripByIdRequest is the DTO for getting a trip by ID
type GetTripByIdRequest struct {
	TripId int `json:"trip_id"`
}

// TripMemberDetails represents a member in a trip with their role
type TripMemberDetails struct {
	MemberId string          `json:"member_id"`
	Email    string          `json:"email"`
	Name     string          `json:"name"`
	ImageUrl string          `json:"image_url"`
	Role     enum.MemberRole `json:"role"`
}

// GetTripByIdResponse is the DTO response for getting a trip by ID
type GetTripByIdResponse struct {
	ID           int                 `json:"trip_id"`
	OwnerId      string              `json:"owner_id"`
	TripName     *string             `json:"trip_name"`
	Description  *string             `json:"description"`
	StartDate    *string             `json:"start_date"`
	EndDate      *string             `json:"end_date"`
	MainLocation *string             `json:"main_location"`
	ImageUrl     *string             `json:"image_url"`
	Role         enum.MemberRole     `json:"role"`
	Members      []TripMemberDetails `json:"members"`
}

func (t GetTripByIdResponse) FromDomain(resp *domain.GetTripByIdResponse) *GetTripByIdResponse {
	trip := resp.Trip

	// Convert members
	members := make([]TripMemberDetails, len(resp.Members))
	for i, m := range resp.Members {
		members[i] = TripMemberDetails{
			MemberId: m.MemberId,
			Email:    m.Email,
			Name:     m.Name,
			ImageUrl: m.ImageUrl,
			Role:     m.Role,
		}
	}

	// Convert TripName to *string
	var tripName *string
	if trip.TripName != "" {
		tripName = &trip.TripName
	}

	// Convert Description to *string
	var description *string
	if trip.Description != "" {
		description = &trip.Description
	}

	// Convert StartDate (time.Time) to *string
	startDateStr := utils.DateTimeToDateString(&trip.StartDate)
	var startDate *string
	if startDateStr != "" {
		startDate = &startDateStr
	}

	// Convert EndDate (time.Time) to *string
	endDateStr := utils.DateTimeToDateString(&trip.EndDate)
	var endDate *string
	if endDateStr != "" {
		endDate = &endDateStr
	}

	// Convert MainLocation to *string
	var mainLocation *string
	if trip.MainLocation != "" {
		mainLocation = &trip.MainLocation
	}

	var imageUrl *string
	if trip.ImageUrl != "" {
		imageUrl = &trip.ImageUrl
	}

	return &GetTripByIdResponse{
		ID:           trip.ID,
		OwnerId:      trip.OwnerId,
		TripName:     tripName,
		Description:  description,
		StartDate:    startDate,
		EndDate:      endDate,
		MainLocation: mainLocation,
		ImageUrl:     imageUrl,
		Role:         resp.Role,
		Members:      members,
	}
}

type TripInvitationRequest struct {
	TripId int             `json:"trip_id"`
	Email  string          `json:"email"`
	Role   enum.MemberRole `json:"role"`
}

func (t TripInvitationRequest) ToDomain(memberId string) *domain.TripInvitationRequest {
	return &domain.TripInvitationRequest{
		MemberId: memberId,
		TripId:   t.TripId,
		Email:    t.Email,
		Role:     t.Role,
	}
}

type TripInvitationResponse struct {
	Email string          `json:"email"`
	Role  enum.MemberRole `json:"role"`
}

func (t TripInvitationResponse) FromDomain(resp *domain.TripInvitationResponse) *TripInvitationResponse {
	return &TripInvitationResponse{
		Email: resp.Email,
		Role:  resp.Role,
	}
}

type DeleteTripMemberRequest struct {
	TripId int    `json:"trip_id"`
	Email  string `json:"email"`
}

func (d DeleteTripMemberRequest) ToDomain() *domain.DeleteTripMemberRequest {
	return &domain.DeleteTripMemberRequest{
		TripId: d.TripId,
		Email:  d.Email,
	}
}

type GetTripMemberRoleRequest struct {
	TripId int `json:"trip_id"`
}

func (t GetTripMemberRoleRequest) ToDomain(memberId string) *domain.GetTripMemberRoleRequest {
	return &domain.GetTripMemberRoleRequest{
		TripId:   t.TripId,
		MemberId: memberId,
	}
}

type GetTripMemberRoleResponse struct {
	TripId   int             `json:"trip_id"`
	MemberId string          `json:"member_id"`
	Role     enum.MemberRole `json:"role"`
}

func (t GetTripMemberRoleResponse) FromDomain(d domain.GetTripMemberRoleResponse) *GetTripMemberRoleResponse {
	return &GetTripMemberRoleResponse{
		TripId:   d.TripId,
		MemberId: d.MemberId,
		Role:     d.Role,
	}
}

type GetTripMembersRequest struct {
	TripId int `json:"trip_id"`
}

func (t GetTripMembersRequest) ToDomain(memberId string) *domain.GetTripMembersRequest {
	return &domain.GetTripMembersRequest{
		TripId:   t.TripId,
		MemberId: memberId,
	}
}

type GetTripMembersResponse struct {
	Members []TripMemberDetails `json:"members"`
}

func (t GetTripMembersResponse) FromDomain(dm []domain.TripMemberDetails) *GetTripMembersResponse {
	members := make([]TripMemberDetails, len(dm))
	for i, m := range dm {
		members[i] = TripMemberDetails{
			MemberId: m.MemberId,
			Email:    m.Email,
			Name:     m.Name,
			ImageUrl: m.ImageUrl,
			Role:     m.Role,
		}
	}
	return &GetTripMembersResponse{
		Members: members,
	}
}
