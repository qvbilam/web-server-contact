package resource

import proto "contact/api/qvbilam/contact/v1"

type GroupsResource struct {
	Total int64    `json:"total"`
	List  []*group `json:"list"`
}

type group struct {
	ID               int64        `json:"id"`
	Code             int64        `json:"code"`
	Name             string       `json:"name"`
	MemberCount      int64        `json:"member_count"`
	AllowMemberCount int64        `json:"allow_member_count"`
	Member           *groupMember `json:"member"`
}

func (r *GroupsResource) Resource(p *proto.GroupsResponse) *GroupsResource {
	res := GroupsResource{}
	res.Total = p.Total

	var groups []*group
	if res.Total > 0 {
		for _, item := range p.Groups {
			m := groupMember{}
			groups = append(groups, &group{
				ID:               item.Id,
				Code:             item.Code,
				Name:             item.Name,
				MemberCount:      item.MemberCount,
				AllowMemberCount: item.AllowMemberCount,
				Member:           m.Resource(item.Member),
			})
		}

		res.List = groups
	}
	return &res
}
