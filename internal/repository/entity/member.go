package entity

type Member struct {
	ID       string `json:"member_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	ImageUrl string `json:"image_url"`
}

type GetExistingMemberResp struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}