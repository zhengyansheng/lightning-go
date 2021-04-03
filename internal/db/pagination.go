package db

import "github.com/douyu/jupiter/pkg/store/gorm"

// 单表分页
type PaginationQ struct {
	Ok     bool        `json:"ok"`
	Limit  int         `form:"size" json:"size"`
	Offset int         `form:"page" json:"page"`
	Data   interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list
	Total  int         `json:"total"`
}

func CRUD(p *PaginationQ, queryTx *gorm.DB, list interface{}) (int, error) {
	if p.Limit < 1 {
		p.Limit = 10
	}
	if p.Offset < 1 {
		p.Offset = 1
	}

	var total int
	err := queryTx.Count(&total).Error
	if err != nil {
		return 0, err
	}
	offset := p.Limit * (p.Offset - 1)
	err = queryTx.Limit(p.Limit).Offset(offset).Find(list).Error
	if err != nil {
		return 0, err
	}
	return total, err
}
