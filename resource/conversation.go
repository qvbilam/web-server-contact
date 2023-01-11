package resource

import (
	proto "contact/api/qvbilam/contact/v1"
)

type ConversationsResource struct {
	Total int64                   `json:"total"`
	List  []*ConversationResource `json:"list"`
}

type ConversationResource struct {
	ID              int64           `json:"id"`
	UserID          int64           `json:"user_id"`
	ObjectType      string          `json:"object_type"`
	ObjectID        int64           `json:"object_id"`
	Object          *ObjectResource `json:"object"`
	NewsCount       int64           `json:"news_count"`
	Tips            string          `json:"tips"`
	LastMessage     string          `json:"last_message"`
	LastMessageTime int64           `json:"last_message_time"`
}

func (r *ConversationsResource) Resource(p *proto.ConversationsResponse) *ConversationsResource {

	c := ConversationsResource{}
	c.Total = p.Total
	var cs []*ConversationResource

	for _, item := range p.Conversations {
		cs = append(cs, &ConversationResource{
			ID:         item.Id,
			UserID:     item.UserId,
			ObjectType: item.ObjectType,
			ObjectID:   item.ObjectId,
			Object: &ObjectResource{
				ID:     item.Object.Id,
				Name:   item.Object.Name,
				Avatar: item.Object.Avatar,
				Remark: item.Object.Remark,
				IsDND:  item.Object.IsDND,
			},
			NewsCount:       item.NewsCount,
			Tips:            item.Tips,
			LastMessage:     item.LastMessage,
			LastMessageTime: item.LastMessageTime,
		})
	}

	c.List = cs
	return &c
}
