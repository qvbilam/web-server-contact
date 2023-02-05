package resource

import proto "contact/api/qvbilam/contact/v1"

type GroupMembersResource struct {
	Total int64          `json:"total"`
	List  []*groupMember `json:"list"`
}

type groupMember struct {
	ID          int64         `json:"id"`
	Nickname    string        `json:"nickname"`
	User        *UserResource `json:"user"`
	Role        int64         `json:"role"`
	Level       int64         `json:"level"`
	Remark      string        `json:"remark"`
	IsDnd       bool          `json:"is_dnd"`
	IsBanned    bool          `json:"is_banned"`
	CreatedTime int64         `json:"created_time"`
}

func (r *GroupMembersResource) Resource(p *proto.GroupMembersResponse) *GroupMembersResource {
	res := GroupMembersResource{}
	res.Total = p.Total

	var members []*groupMember
	if res.Total > 0 {
		for _, item := range p.Members {
			u := UserResource{}

			members = append(members, &groupMember{
				ID:          item.Id,
				Nickname:    item.Nickname,
				User:        u.Resource(item.User),
				Role:        item.Role,
				Level:       item.Level,
				Remark:      item.Remark,
				IsDnd:       item.IsDnd,
				IsBanned:    item.IsBanned,
				CreatedTime: item.CreatedTime,
			})
		}

		res.List = members
	}

	return &res
}

func (r *groupMember) Resource(p *proto.GroupMemberResponse) *groupMember {
	if p == nil {
		return nil
	}

	u := UserResource{}
	return &groupMember{
		ID:          p.Id,
		Nickname:    p.Nickname,
		User:        u.Resource(p.User),
		Role:        p.Role,
		Level:       p.Level,
		Remark:      p.Remark,
		IsDnd:       p.IsDnd,
		IsBanned:    p.IsBanned,
		CreatedTime: p.CreatedTime,
	}
}
