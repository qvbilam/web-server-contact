package resource

import userProto "contact/api/qvbilam/user/v1"

type UserResource struct {
	ID       int64  `json:"id"`
	Code     int64  `json:"code"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Gender   string `json:"gender"`
}

func (r *UserResource) Resource(p *userProto.UserResponse) *UserResource {
	if p == nil {
		return nil
	}
	return &UserResource{
		ID:       p.Id,
		Code:     p.Code,
		Nickname: p.Nickname,
		Avatar:   p.Avatar,
		Gender:   p.Gender,
	}
}
