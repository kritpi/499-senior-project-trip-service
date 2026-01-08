package domain

import "github.com/kritpi/499-senior-project-trip-service/internal/repository/entity"

type Member struct {
	ID       string `json:"member_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageUrl string `json:"image_url"`
}

func (m Member) FromEntity(e entity.Member) *Member {
	return &Member{
		ID:       e.ID,
		Name:     e.Name,
		Email:    e.Email,
		ImageUrl: e.ImageUrl,
	}
}

type GetExistingMemberResp struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (g GetExistingMemberResp) FromEntity(e entity.GetExistingMemberResp) *GetExistingMemberResp {
	return &GetExistingMemberResp{
		ID:    e.ID,
		Email: e.Email,
	}
}
