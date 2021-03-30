package service

import (
	"errors"
	"go-ops/internal/models/onduty"
	"go-ops/pkg/tools"
)

// 验证结构体
type OnDutySerializer struct {
	ID     uint   `json:"-" binding:"omitempty"`                         // 可传可不传
	Name   string `json:"name" binding:"omitempty,min=1"`                // 节点名称
	Parent int    `json:"parent" form:"parent" binding:"omitempty,gt=0"` // 允许为空，但是如果传递值必须大于等于1
}

func (d OnDutySerializer) Create() (err error) {
	duty := onduty.OnDuty{}
	if len(d.Name) == 0 {
		err = errors.New("name is required. ")
		return
	}

	duty.Name = d.Name
	duty.Parent = d.Parent

	_, err = duty.IsNameExists()
	if err == nil {
		return
	}

	// 格式化打印验证的数据
	tools.PrettyPrint(duty)
	err = duty.Create()
	return
}

func (d OnDutySerializer) Update() (err error) {
	var v = make(map[string]interface{})
	if len(d.Name) != 0 {
		v["name"] = d.Name
	}
	if d.Parent != 0 {
		v["parent"] = d.Parent
	}

	duty := onduty.OnDuty{}
	duty.ID = d.ID

	// 格式化打印验证的数据
	tools.PrettyPrint(v)
	err = duty.Update(v)
	return
}

func (d OnDutySerializer) Delete() (err error) {
	duty := onduty.OnDuty{}
	duty.ID = d.ID
	err = duty.Delete()
	return
}
