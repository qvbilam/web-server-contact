package resource

import (
	proto "contact/api/qvbilam/contact/v1"
)

type FriendsResource struct {
	Total int64     `json:"total"`
	List  []*friend `json:"list"`
}

type friend struct {
	ID     int64         `json:"id"`
	Remark string        `json:"remark"`
	UserID int64         `json:"user_id"`
	User   *UserResource `json:"user"`
}

func (r *FriendsResource) Resource(p *proto.FriendsResponse) *FriendsResource {
	res := FriendsResource{}
	res.Total = p.Total
	var friends []*friend

	if p.Total > 0 {
		for _, item := range p.Friends {
			u := UserResource{}

			friends = append(friends, &friend{
				ID:     item.Id,
				UserID: item.Friend.Id,
				Remark: item.Remark,
				User:   u.Resource(item.Friend),
			})
		}
		res.List = friends
	}

	return &res
}
