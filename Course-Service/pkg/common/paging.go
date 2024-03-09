package common

type Paging struct {
	Limit int `json:"pageSize" form:"pageSize"`
	Page int `json:"page" form:"page"`
}
func (p *Paging) FullFill() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 1
	}
}
