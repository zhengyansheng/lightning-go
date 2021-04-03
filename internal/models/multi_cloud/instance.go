package multi_cloud

import "lightning-go/internal/db"

// 记录生命周期
type InstanceLifeCycle struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	Action     string `json:"action" binding:"required"` // 动作; 关机，开机，重启，创建, 销毁
	Uri        string `json:"uri"`
	Method     string `json:"method"`
	InstanceId string `json:"instance_id"`
	Body       string `json:"body"`
	Response   string `json:"response"`
	IsSuccess  bool   `json:"is_success"`
}

func (s *InstanceLifeCycle) TableName() string {
	return "instance_life_cycle"
}

func (s *InstanceLifeCycle) Create() (err error) {
	err = db.DB.Debug().Table(s.TableName()).Create(s).Error
	return
}
