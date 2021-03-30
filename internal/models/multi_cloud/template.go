package multi_cloud

import (
	"github.com/jinzhu/gorm"
	"go-ops/internal/db"
	"go-ops/internal/models"
)

// 创建云主机配置参数
type CloudTemplate struct {
	gorm.Model
	AppKey           string      `json:"app_key" binding:"required"`            // 关联到service_tree服务
	Env              string      `json:"env"`                                   // 环境 如果绑在非环境节点时
	Account          string      `json:"account" binding:"required"`            // 帐号
	PayType          string      `json:"pay_type"`                              // 支付方式 包年和按月
	RegionId         string      `json:"region_id" binding:"required"`          // 地域
	ZoneId           string      `json:"zone_id"`                               // 可用区
	InstanceType     string      `json:"instance_type" binding:"required"`      // 机型
	IsPublic         bool        `json:"is_public"`                             // 公网地址
	IsEip            bool        `json:"is_eip"`                                // 弹性IP
	ImageId          string      `json:"image_id" binding:"required"`           // 镜像ID
	VpcId            string      `json:"vpc_id" binding:"required"`             // vpc id
	SubnetId         string      `json:"subnet_id"`                             // 子网
	Disks            models.JSON `json:"disks" binding:"required"`              // 云盘
	SecurityGroupIds models.JSON `json:"security_group_ids" binding:"required"` // 安全组
	Count            int         `json:"-" binding:"omitempty"`                 // 数量
}

type Disk struct {
	DiskStorageSize int    `json:"disk_storage_size"` // 云盘存储大小
	DiskStorageType string `json:"disk_storage_type"` // 云盘存储类型 cloud cloud-ssd
	DiskType        string `json:"disk_type"`         // 云盘类型  system | data
}

func (s *CloudTemplate) TableName() string {
	return "cloud_template"
}

func (s *CloudTemplate) Create() (err error) {
	err = db.DB.Debug().Table(s.TableName()).Create(s).Error
	return
}

func (s *CloudTemplate) Get() (template CloudTemplate, err error) {
	err = db.DB.Debug().Table(s.TableName()).Where("id = ?", s.ID).Find(&template).Error
	return
}

func (s *CloudTemplate) Delete() (err error) {
	err = db.DB.Debug().Table(s.TableName()).Where("id = ?", s.ID).Delete(&s).Error
	return
}

func (s *CloudTemplate) ListByAppKey() (templates []CloudTemplate, err error) {
	table := db.DB.Debug().Table(s.TableName())
	table = table.Where("app_key = ?", s.AppKey)

	if s.Account != "" {
		table = table.Where("account = ?", s.Account)
	}
	if s.Env != "" {
		table = table.Where("env = ?", s.Env)
	}
	if s.RegionId != "" {
		table = table.Where("region_id = ?", s.RegionId)
	}

	err = table.Find(&templates).Error
	return
}
