package onduty

import (
	"errors"
	"github.com/jinzhu/gorm"
	"go-ops/internal/db"
)

// 排班系统Model
type OnDuty struct {
	gorm.Model
	Name   string `json:"name"`
	Parent int    `json:"parent"`
	Level  int    `json:"level"`
}

func (s *OnDuty) TableName() string {
	return "on_duty"
}

// hook: 自动生成 level 字段
func (s *OnDuty) BeforeCreate() (err error) {
	//
	var level int
	if s.Parent == 0 {
		level = 0
	} else {
		onDuty := OnDuty{
			Parent: s.Parent,
		}
		level, err = onDuty.GetLevelByParent()
		if err != nil {
			return err
		}
		level += 1
	}
	s.Level = level
	return
}

// 创建
func (s *OnDuty) Create() (err error) {
	err = db.DB.Debug().Table(s.TableName()).Create(s).Error
	return
}

// 修改
func (s *OnDuty) Update(v map[string]interface{}) (err error) {
	err = db.DB.Debug().Table(s.TableName()).Where("id = ?", s.ID).Updates(v).Error
	return
}

// 删除
func (s *OnDuty) Delete() (err error) {
	_, err = s.isIdExists()
	if err != nil {
		return
	}
	err = db.DB.Debug().Table(s.TableName()).Where("id = ?", s.ID).Delete(s).Error
	return
}

// 查询， 通过 parent 查询 level
func (s *OnDuty) GetLevelByParent() (level int, err error) {
	err = db.DB.Debug().Table(s.TableName()).Where("id = ?", s.Parent).First(&s).Error
	if err != nil {
		level = s.Level
	}
	return
}

// 查询, id是否存在
func (s *OnDuty) isIdExists() (id uint, err error) {
	err = db.DB.Debug().Table(s.TableName()).Where("id = ?", s.ID).First(&s).Error
	if err != nil {
		id = s.ID
	}
	return
}

// 查询, 同一级 name 不允许重复
// 存在 nil; 不存在 err
func (s *OnDuty) IsNameExists() (id uint, err error) {
	var count int
	err = db.DB.Debug().Table(s.TableName()).Where("name = ? and parent = ?", s.Name, s.Parent).Count(&count).Error
	if err != nil {
		return
	}
	if count == 0 {
		err = errors.New("not found. ")
	}
	return
}
