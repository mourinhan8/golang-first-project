package common

type Paging struct {
	Limit int   `json:"limit" form:"limit"`
	Page  int   `json:"page" form:"page"`
	Total int64 `json:"total" form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit > 100 {
		p.Limit = 10
	}
}
