package multi_cloud

import (
	"github.com/gin-gonic/gin"
	"go-ops/internal/service/multi_cloud"
	"go-ops/pkg/tools"
)

func CreateInstanceView(c *gin.Context) {
	var s service.InstanceSerializer
	if err := c.ShouldBindJSON(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	instanceId, err := s.CreateInstance()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// 返回
	tools.JSONOk(c, instanceId)
}

func StartInstanceView(c *gin.Context) {
	var s service.InstanceStartSerializer
	if err := c.ShouldBindJSON(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	tools.PrettyPrint(s)
	response, err := s.StartInstance()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// 返回
	tools.JSONOk(c, response)
}

func StopInstanceView(c *gin.Context) {
	var s service.InstanceStartSerializer
	if err := c.ShouldBindJSON(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	tools.PrettyPrint(s)
	response, err := s.StopInstance()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// 返回
	tools.JSONOk(c, response)
}

func ListInstancesView(c *gin.Context) {
	var s service.InstanceDetailSerializer
	if err := c.ShouldBind(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	//tools.PrettyPrint(s)
	response, err := s.ListInstances()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// 返回
	tools.JSONOk(c, response)
}

func InstanceDetailView(c *gin.Context) {

}

func ListRegionView(c *gin.Context) {
	var s service.InstanceDetailSerializer
	if err := c.ShouldBind(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	//tools.PrettyPrint(s)
	response, err := s.ListRegions()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// 返回
	tools.JSONOk(c, response)
}