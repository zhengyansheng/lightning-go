package service

import (
	"fmt"
	"go-ops/pkg/multi_cloud_sdk"
)

// 验证结构体
type InstanceSerializer struct {
	Account          string   `json:"account" binding:"required"`
	PayType          string   `json:"pay_type" binding:"required"`
	Hostname         string   `json:"hostname" binding:"required"`
	RegionId         string   `json:"region_id" binding:"required"`
	ZoneId           string   `json:"zone_id" binding:"required"`
	InstanceType     string   `json:"instance_type" binding:"required"`
	ImageId          string   `json:"image_id" binding:"required"`
	VpcId            string   `json:"vpc_id" binding:"required"`
	SubnetId         string   `json:"subnet_id" binding:"required"`
	SecurityGroupIds []string `json:"security_group_ids" binding:"required"`
	DryRun           bool     `json:"dry_run" binding:"omitempty"`
}

func (s InstanceSerializer) CreateInstance() (string, error) {
	fmt.Println(s.Account, s.RegionId)
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		return "new factory err.", err
	}
	return clt.CreateInstance(s.PayType, s.Hostname, s.InstanceType, s.ZoneId, s.ImageId, s.VpcId, s.SubnetId, s.SecurityGroupIds, s.DryRun)
}

type InstanceStartSerializer struct {
	Account    string `json:"account" binding:"required"`
	RegionId   string `json:"region_id" binding:"required"`
	InstanceId string `json:"instance_id" binding:"required"`
}

func (s InstanceStartSerializer) StartInstance() (string, error) {
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		return "new factory err.", err
	}
	return clt.StartInstance(s.InstanceId)
}

func (s InstanceStartSerializer) StopInstance() (string, error) {
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		return "new factory err.", err
	}
	return clt.StopInstance(s.InstanceId)
}

type InstanceDetailSerializer struct {
	Account    string `form:"account" binding:"required"`
	RegionId   string `form:"region_id" binding:"required"`
	InstanceId string `form:"instance_id"`
}

func (s InstanceDetailSerializer) ListInstances() ([]map[string]interface{}, error) {
	var instances []map[string]interface{}
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		return instances, err
	}
	return clt.ListInstances()
}