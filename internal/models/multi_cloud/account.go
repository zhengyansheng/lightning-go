package multi_cloud

import (
	"go-ops/internal/db"
)

// 创建云主机配置参数
type Account struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	EnName      string `json:"en_name" binding:"required"`       // 英文名
	CnName      string `json:"cn_name" binding:"required"`       // 中文名
	Platform    string `json:"platform" binding:"required"`      // 平台; aws/ali/ten/hw
	AccessKeyId string `json:"access_key_id" binding:"required"` // 密钥
	SecretKeyId string `json:"secret_key_id" binding:"required"` // 密钥
	RootId      string `json:"root_id" binding:"omitempty"`      // 主帐号ID
}

func (s *Account) TableName() string {
	return "cloud_account"
}

func (s *Account) Create() (err error) {
	err = db.DB.Debug().Table(s.TableName()).Create(s).Error
	return
}

func (s *Account) Delete() (err error) {
	err = db.DB.Debug().Table(s.TableName()).Where("ID = ?", s.ID).Delete(&s).Error
	return
}

func (s *Account) List() (accounts []Account, err error) {
	err = db.DB.Debug().Table(s.TableName()).Find(&accounts).Error
	return
}

func (s *Account) GetByAccount() (account Account, err error) {
	err = db.DB.Debug().Table(s.TableName()).Where("en_name = ?", s.EnName).Find(&account).Error
	return
}
