package multi_cloud

import (
	"github.com/gin-gonic/gin"
	"go-ops/internal/models/multi_cloud"
	"go-ops/pkg/tools"
)

func CreateAccountView(c *gin.Context) {
	var s multi_cloud.Account
	if err := c.ShouldBindJSON(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	err := s.Create()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// 返回
	tools.JSONOk(c, "Create ok.")
}

func ListAccountView(c *gin.Context) {
	var s multi_cloud.Account
	accountsInfo, err := s.List()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// 返回
	tools.JSONOk(c, accountsInfo)
}

func DeleteAccountView(c *gin.Context) {
	var s multi_cloud.Account
	pk := c.Param("id")
	pkUint, _ := tools.StringToUint(pk)
	s.ID = pkUint
	err := s.Delete()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// 返回
	tools.JSONOk(c, "Delete ok.")
}
