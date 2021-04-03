package multi_cloud

import (
	"lightning-go/internal/models/multi_cloud"
	"lightning-go/pkg/tools"

	"github.com/gin-gonic/gin"
)

func CreateAccountView(c *gin.Context) {
	// Validate field
	var s multi_cloud.Account
	if err := c.ShouldBindJSON(&s); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Create
	err := s.Create()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Response
	tools.JSONOk(c, "Create ok.")
}

func ListAccountView(c *gin.Context) {
	var s multi_cloud.Account
	// List
	accountsInfo, err := s.List()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Response
	tools.JSONOk(c, accountsInfo)
}

func DeleteAccountView(c *gin.Context) {
	var s multi_cloud.Account
	pk := c.Param("id")
	pkUint, _ := tools.StringToUint(pk)
	s.ID = pkUint

	// Delete
	err := s.Delete()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	// Response
	tools.JSONOk(c, "Delete ok.")
}
