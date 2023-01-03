package validate

type GroupCreate struct {
	Name      string `form:"name" json:"name" binding:"required,min=3,max=100"`
	Avatar    string `form:"avatar" json:"avatar" binding:"omitempty,url"`
	Cover     string `form:"cover" json:"cover" binding:"omitempty,url"`
	Introduce string `form:"introduce" json:"introduce" binding:"omitempty,max=500"`
}
