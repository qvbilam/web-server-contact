package validate

type ConversationCreate struct {
	Type     string `form:"type" json:"type" binding:"required"`
	ObjectId int64  `form:"object_id" json:"object_id" binding:"required,numeric"`
}
