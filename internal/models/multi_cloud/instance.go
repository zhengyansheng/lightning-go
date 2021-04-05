package multi_cloud

import (
	"lightning-go/internal/db"
	"lightning-go/internal/models"
)

// 记录生命周期
type InstanceLifeCycle struct {
	InstanceId string          `json:"instance_id"`
	Uri        string          `json:"uri"`
	Method     string          `json:"method"`
	Query      string          `json:"query"`
	Body       models.JSON     `json:"body" gorm:"type:text"`
	RemoteIp   string          `json:"remote_ip"`
	CreateUser string          `json:"create_user"`
	Response   models.JSON     `json:"response" gorm:"type:text"`
	IsSuccess  bool            `json:"is_success"` // 0 fail || 1 success
	CreatedAt  models.JSONTime `json:"created_at"`
}

func (s *InstanceLifeCycle) TableName() string {
	return "instance_life_cycle"
}

func (s *InstanceLifeCycle) Create() (err error) {
	err = db.DB.Debug().Table(s.TableName()).Create(s).Error
	return
}

func (s *InstanceLifeCycle) GetByInstanceId() (instances []InstanceLifeCycle, err error) {
	err = db.DB.Debug().Table(s.TableName()).Where("instance_id = ?", s.InstanceId).Find(&instances).Order("created_at").Error
	return
}
