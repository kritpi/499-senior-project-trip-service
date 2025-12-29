package enum

// TripRole represents the role of a member in a trip
type TripRole string

const (
	TripRoleOwner  TripRole = "OWNER"
	TripRoleEditor TripRole = "EDITOR"
	TripRoleViewer TripRole = "VIEWER"
)
