package enum

// TripRole represents the role of a member in a trip
type MemberRole string

const (
	MemberRoleOwner  MemberRole = "OWNER"
	MemberRoleEditor MemberRole = "EDITOR"
	MemberRoleViewer MemberRole = "VIEWER"
)

type ActivityCategory string

const (
	ActivityCategoryFood          ActivityCategory = "FOOD"
	ActivityCategoryAttraction    ActivityCategory = "ATTRACTION"
	ActivityCategoryAccommodation ActivityCategory = "ACCOMMODATION"
	ActivityCategoryTransport     ActivityCategory = "TRANSPORT"
	ActivityOther                 ActivityCategory = "OTHER"
)

// "FOOD" | "ATTRACTION" | "ACCOMMODATION" | "TRANSPORT" | "OTHER" | ""

type SplitType string

const (
	ExpenseCustom        SplitType = "CUSTOM"
	ExpenseAllEqual      SplitType = "ALL_EQUAL"
	ExpenseSelectedEqual SplitType = "SELECTED_EQUAL"
)
