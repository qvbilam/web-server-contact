package validate

type FriendApply struct {
	UserID  int64  `form:"user_id" json:"user_id" binding:"required,numeric"`
	Content string `form:"content" json:"content" binding:"omitempty,min=1,max=100"`
}
