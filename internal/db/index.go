package db

import (
	"github.com/douyu/jupiter/pkg/store/gorm"
)

var DB *gorm.DB

func Init() {

	DB = gorm.StdConfig("master").Build()

}

func InitTest() {

	config := gorm.DefaultConfig()
	config.DSN = "root:123456@tcp(127.0.0.1:3306)/idrac?charset=utf8&parseTime=True&loc=Local&readTimeout=1s&timeout=1s&writeTimeout=3s"
	config.Debug = true

	DB = config.Build()
	defer DB.Close()
}

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
