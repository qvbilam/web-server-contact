package resource

import proto "contact/api/qvbilam/contact/v1"

type NewsResource struct {
	Total int64    `json:"total"`
	List  []*apply `json:"list"`
}

type apply struct {
	ID      int64         `json:"id"`
	UserID  int64         `json:"user_id"`
	Content string        `json:"content"`
	Status  int64         `json:"status"`
	User    *UserResource `json:"user"`
}

func (r *NewsResource) Resource(p *proto.FriendAppliesResponse) *NewsResource {
	res := NewsResource{}
	res.Total = p.Total
	var news []*apply

	if res.Total > 0 {
		for _, item := range p.Users {

			u := UserResource{}
			news = append(news, &apply{
				ID:      item.Id,
				UserID:  item.ApplyUserId,
				Content: item.Content,
				Status:  item.Status,
				User:    u.Resource(item.ApplyUser),
			})
		}

		res.List = news
	}

	return &res
}
