package message

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "go-ops/docs"
	"go-ops/internal/db"
	"go-ops/pkg/tools"
	"math/rand"
	"strconv"
	"time"
)

// @Summary 增加钉钉消息渠道
// @Description 增加钉钉消息渠道
// @Accept  application/json
// @Produce application/json
// @Param   data     body ResSendDingConfig true     "消息主题名"
// @Success 200 {string} {"code": 0, data: "", "message": ""}
// @Failure 500 {string} {"code": -1 data: "", "message": ""}
// @Router /api/v1/message/ding/add [post]
func AddDing(c *gin.Context) {
	var (
		dbSendDingConfig db.SendDingConfig
		err              error
	)
	err = c.Bind(&dbSendDingConfig)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}

	// 生成消息模版ID
	dbSendDingConfig.MsgId = rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000)

	if err = dbSendDingConfig.SendCreate(); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("Creation failed  %v", err.Error()))
		return
	}
	tools.JSONOk(c, "Created successfully")
}

// @Summary 增加邮件消息渠道
// @Description 增加邮件消息渠道
// @Accept application/json
// @Produce application/json
// @Param   data     body   ResSendMailConfig   true     "增加邮件消息渠道"
// @Success 200 {string} {"code": 0, data: "", "message": ""}
// @Failure 500 {string} {"code": -1 data: "", "message": ""}
// @Router /api/v1/message/mail/add [post]
func AddMail(c *gin.Context) {
	var (
		dbSendMailConfig db.SendMailConfig
		err              error
	)
	err = c.Bind(&dbSendMailConfig)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}

	// 生成消息模版ID
	dbSendMailConfig.MsgId = rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(10000000)
	if err = dbSendMailConfig.SendCreate(); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("Creation failed  %v", err.Error()))
		return
	}
	tools.JSONOk(c, "Created successfully")
}

// @Summary 删除消息渠道
// @Description 删除DingDing消息渠道
// @Accept application/json
// @Produce application/json
// @Param   msg_id     query    int64     true        "消息渠道id"
// @Success 200 {string} {"code": 0, data: "", "message": ""}
// @Failure 500 {string} {"code": -1 data: "", "message": ""}
// @Router /api/v1/message/ding/delete [delete]
func DelDing(c *gin.Context) {
	var (
		dbSendDingConfig db.SendDingConfig
		err              error
	)
	msg := c.Query("msg_id")
	fmt.Println(msg)
	dbSendDingConfig.MsgId, err = strconv.ParseInt(msg, 10, 64)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("Del failed  %v", err.Error()))
		return
	}
	if err = dbSendDingConfig.SendDelete(); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("Del failed  %v", err.Error()))
		return
	}
	tools.JSONOk(c, "Del successfully")

}

// @Summary 更新钉钉消息渠道
// @Description 更新钉钉消息渠道
// @Accept application/json
// @Produce application/json
// @Param   msg_id      query    int64     true        "消息渠道id"
// @Param   subject     query    string     true        "消息主题"
// @Param   ding_token  query    string     true        "钉钉Token"
// @Param   ding_secret query    string     true        "钉钉机器人secret"
// @Success 200 {string} {"code": 0, data: "", "message": ""}
// @Failure 500 {string} {"code": -1 data: "", "message": ""}
// @Router /api/v1/message/ding/update [put]
func UpdateDing(c *gin.Context) {
	var (
		dbSendDingConfig db.SendDingConfig

		err error
	)

	msg := c.Query("msg_id")
	dbSendDingConfig.MsgId, err = strconv.ParseInt(msg, 10, 64)
	dbSendDingConfig.Subject = c.Query("subject")
	dbSendDingConfig.DingToken = c.Query("ding_token")
	dbSendDingConfig.DingSecret = c.Query("ding_secret")

	if err = dbSendDingConfig.SendUpdate(); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("Update failed  %v", err.Error()))
		return
	}
	tools.JSONOk(c, "Update successfully")
}

// @Summary 更新邮件消息渠道
// @Description 更新邮件消息渠道
// @Accept application/json
// @Produce application/json
// @Param   msg_id     query    int64      true        "消息渠道id"
// @Param   subject    query    string     true        "消息主题"
// @Param   user       query    string     true        "邮箱用户"
// @Param   pass_word  query    string     true        "邮箱密码"
// @Param   port       query    string     true        "邮箱端口"
// @Param   host       query    string     true        "邮箱地址"
// @Success 200 {string} {"code": 0, data: "", "message": ""}
// @Failure 500 {string} {"code": -1 data: "", "message": ""}
// @Router /api/v1/message/mail/update [put]
func UpdateMail(c *gin.Context) {
	var (
		dbSendDingConfig db.SendMailConfig

		err error
	)

	msg := c.Query("msg_id")
	dbSendDingConfig.MsgId, err = strconv.ParseInt(msg, 10, 64)
	dbSendDingConfig.Subject = c.Query("subject")
	dbSendDingConfig.User = c.Query("user")
	dbSendDingConfig.PassWord = c.Query("pass_word")
	dbSendDingConfig.Port = c.Query("port")
	dbSendDingConfig.Host = c.Query("host")

	if err = dbSendDingConfig.SendUpdate(); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("Update failed  %v", err.Error()))
		return
	}
	tools.JSONOk(c, "Update successfully")
}

// @Summary 查询消息渠道
// @Description 查询钉钉消息渠道
// @Accept application/json
// @Produce application/json
// @Param   size     query    string     true        "分页大小"
// @Param   page     query    string     true        "第几页"
// @Success 200 {string} {"code": 0, data: "", "message": ""}
// @Failure 500 {string} {"code": -1 data: "", "message": ""}
// @Router  /api/v1/message/ding/query [get]
func QueryDing(c *gin.Context) {
	var (
		dbSendDingConfigQ db.SendDingConfigQ
		dingList          *[]db.SendDingConfig
		total             int
		err               error
	)

	dbSendDingConfigQ.Limit, _ = strconv.Atoi(c.DefaultQuery("size", "10"))
	dbSendDingConfigQ.Offset, _ = strconv.Atoi(c.DefaultQuery("page", "1"))

	dingList, total, err = dbSendDingConfigQ.SearchDing()
	fmt.Println(dingList)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	tools.JSONokQ(c, total, dingList)
	return

}

// @Summary 删除消息渠道
// @Description 删除mail消息渠道
// @Accept application/json
// @Produce application/json
// @Param   msg_id     query    int64     true        "消息渠道id"
// @Success 200 {string} {"code": 0, data: "", "message": ""}
// @Failure 500 {string} {"code": -1 data: "", "message": ""}
// @Router /api/v1/message/mail/delete [delete]
func DelMail(c *gin.Context) {
	var (
		dbSendMailConfig db.SendMailConfig
		err              error
	)
	msg := c.Query("msg_id")
	fmt.Println(msg)
	dbSendMailConfig.MsgId, err = strconv.ParseInt(msg, 10, 64)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("Del failed  %v", err.Error()))
		return
	}
	if err = dbSendMailConfig.SendDelete(); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, fmt.Sprintf("Del failed  %v", err.Error()))
		return
	}
	tools.JSONOk(c, "Del successfully")

}

// @Summary 查询消息渠道
// @Description 查询邮件消息渠道
// @Accept application/json
// @Produce application/json
// @Param   size     query    string     true        "分页大小"
// @Param   page     query    string     true        "第几页"
// @Success 200 {string} {"code": 0, data: "", "message": ""}
// @Failure 500 {string} {"code": -1 data: "", "message": ""}
// @Router  /api/v1/message/mail/query [get]
func QueryMail(c *gin.Context) {
	var (
		dbSendMailConfigQ db.SendMailConfigQ
		dingList          *[]db.SendMailConfig
		total             int
		err               error
	)

	dbSendMailConfigQ.Limit, _ = strconv.Atoi(c.DefaultQuery("size", "10"))
	dbSendMailConfigQ.Offset, _ = strconv.Atoi(c.DefaultQuery("page", "1"))

	dingList, total, err = dbSendMailConfigQ.SearchDing()
	fmt.Println(dingList)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	tools.JSONokQ(c, total, dingList)
	return

}
