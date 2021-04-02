package multi_cloud

import (
	"fmt"
	"go-ops/pkg/multi_cloud_sdk"
	"go-ops/pkg/tools"

	"github.com/gin-gonic/gin"
)

func CreateInstanceView(c *gin.Context) {
	// Validate field
	s := struct {
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
	}{}
	if err := c.ShouldBindJSON(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Factory CreateInstance
	tools.PrettyPrint(s)
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	response, err := clt.CreateInstance(s.PayType, s.Hostname, s.InstanceType, s.ZoneId, s.ImageId, s.VpcId, s.SubnetId, s.SecurityGroupIds, s.DryRun)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("CreateInstance %v", err.Error()))
		return
	}
	// Response
	tools.JSONOk(c, response)
}

func StartInstanceView(c *gin.Context) {
	// Validate field
	s := struct {
		Account    string `json:"account" binding:"required"`
		RegionId   string `json:"region_id" binding:"required"`
		InstanceId string `json:"instance_id" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Factory StartInstance
	tools.PrettyPrint(s)
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	response, err := clt.StartInstance(s.InstanceId)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("StartInstance %v", err.Error()))
		return
	}
	// Response
	tools.JSONOk(c, response)
}

func StopInstanceView(c *gin.Context) {
	// Validate field
	s := struct {
		Account    string `json:"account" binding:"required"`
		RegionId   string `json:"region_id" binding:"required"`
		InstanceId string `json:"instance_id" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Factory StopInstance
	tools.PrettyPrint(s)
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	response, err := clt.StopInstance(s.InstanceId)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("StopInstance %v", err.Error()))
		return
	}
	// Response
	tools.JSONOk(c, response)
}

func RebootInstanceView(c *gin.Context) {
	// Validate field
	s := struct {
		Account    string `json:"account" binding:"required"`
		RegionId   string `json:"region_id" binding:"required"`
		InstanceId string `json:"instance_id" binding:"required"`
		ForceStop  bool   `json:"force_stop"`
	}{}
	if err := c.ShouldBindJSON(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Factory RebootInstance
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	response, err := clt.RebootInstance(s.InstanceId, s.ForceStop)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("RebootInstance %v", err.Error()))
		return
	}
	// Response
	tools.JSONOk(c, response)
}

func ListInstancesView(c *gin.Context) {
	// Validate field
	s := struct {
		Account  string `form:"account" binding:"required"`
		RegionId string `form:"region_id" binding:"required"`
	}{}
	if err := c.ShouldBind(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Factory ListInstances
	tools.PrettyPrint(s)
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	response, err := clt.ListInstances()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("ListInstances %v", err.Error()))
		return
	}
	// Response
	tools.JSONOk(c, response)
}

func InstanceDetailView(c *gin.Context) {
	// Validate field
	s := struct {
		Account      string `form:"account" binding:"required"`
		RegionId     string `form:"region_id" binding:"required"`
	}{}
	if err := c.ShouldBind(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Factory ListInstance
	tools.PrettyPrint(s)
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	response, err := clt.ListInstance(c.Param("instance_id"))
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("ListInstance %v", err.Error()))
		return
	}
	// Response
	tools.JSONOk(c, response)
}

func ListRegionView(c *gin.Context) {
	// Validate field
	s := struct {
		Account  string `form:"account" binding:"required"`
		RegionId string `form:"region_id" binding:"required"`
	}{}
	if err := c.ShouldBind(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Factory ListRegions
	tools.PrettyPrint(s)
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	response, err := clt.ListRegions()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("ListRegion %v", err.Error()))
		return
	}
	// Response
	tools.JSONOk(c, response)
}

func DestroyInstanceView(c *gin.Context) {
	// Validate field
	s := struct {
		Account    string `json:"account" binding:"required"`
		RegionId   string `json:"region_id" binding:"required"`
		InstanceId string `json:"instance_id" binding:"required"`
		ForceStop  bool   `json:"force_stop"`
	}{}
	if err := c.ShouldBindJSON(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Factory DestroyInstance
	tools.PrettyPrint(s)
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	response, err := clt.DestroyInstance(s.InstanceId, s.ForceStop)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("DestroyInstance %v", err.Error()))
		return
	}
	// Response
	tools.JSONOk(c, response)
}

func ModifyInstanceNameView(c *gin.Context) {
	// Validate field
	s := struct {
		Account      string `json:"account" binding:"required"`
		RegionId     string `json:"region_id" binding:"required"`
		InstanceId   string `json:"instance_id" binding:"required"`
		InstanceName string `json:"instance_name" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Factory ModifyInstanceName
	tools.PrettyPrint(s)
	clt, err := multi_cloud_sdk.NewFactoryByAccount(s.Account, s.RegionId)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	response, err := clt.ModifyInstanceName(s.InstanceId, s.InstanceName)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("ModifyInstanceName %v", err.Error()))
		return
	}
	// Response
	tools.JSONOk(c, response)
}
