package base_req

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// GetById Find by id structure
type GetById struct {
	ID float64 `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint64 {
	return uint64(r.ID)
}

func (r *GetById) Int() int64 {
	return int64(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId"` // 角色ID
}

type ReqIDList struct {
	Ids string `json:"ids" form:"ids"`
}
type Empty struct{}
