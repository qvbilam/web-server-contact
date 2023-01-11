package resource

type ObjectResource struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Remark string `json:"remark"`
	IsDND  bool   `json:"is_dnd"`
}
