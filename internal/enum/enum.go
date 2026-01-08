package enum

// TripRole represents the role of a member in a trip
type MemberRole string

const (
	MemberRoleOwner  MemberRole = "OWNER"
	MemberRoleEditor MemberRole = "EDITOR"
	MemberRoleViewer MemberRole = "VIEWER"
)
