package entity

import (
	"github.com/kritpi/499-senior-project-trip-service/internal/enum"
	"github.com/shopspring/decimal"
)

type ExpenseWithMemberRow struct {
	ExpenseID     string          `db:"expense_id"`
	Title         string          `db:"title"`
	ExpenseAmount decimal.Decimal `db:"expense_amount"`
	MyShared      decimal.Decimal `db:"my_shared"`
	CreatedBy     string          `db:"created_by"`
	OwnerImage    *string         `db:"owner_image"`
	ImageURL      string          `db:"image_url"`
	SplitType     enum.SplitType  `db:"split_type"`

	MemberID     string          `db:"member_id"`
	MemberName   string          `db:"member_name"`
	MemberImage  *string         `db:"member_image"`
	MemberAmount decimal.Decimal `db:"member_amount"`
}
