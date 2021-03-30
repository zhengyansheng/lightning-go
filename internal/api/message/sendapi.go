package message

import (
	"fmt"
	dingtalk_robot "github.com/JetBlink/dingtalk-notify-go-sdk"
	"github.com/gin-gonic/gin"
	"go-ops/internal/db"
	"go-ops/pkg/tools"
	"time"
)

var GobalSend = newSend()

// @Summary 发送钉钉消息
// @Description 发送钉钉消息
// @Accept  application/json
// @Produce application/json
// @Param   data     body  ResSendDing   true     "消息主题名"
// @Success 200 {string} {"code": 0, data: "", "message": ""}
// @Failure 500 {string} {"code": -1 data: "", "message": ""}
// @Router /api/v1/message/ding/send [post]
func (d *DingChan) SendDing(c *gin.Context) {
	var (
		resDingMsg     *ResSendDing
		respSendConfig db.SendDingConfig
		dingMSg        DingMSg
		err            error
	)

	err = c.Bind(&resDingMsg)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}

	respSendConfig.MsgId = resDingMsg.MsgId

	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, "参数解析错误")
		return
	}

	respSendConfig, err = respSendConfig.GetSend()

	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, "参数解析错误")
		return
	}

	dingMSg.Data = resDingMsg.Data
	dingMSg.DingToken = respSendConfig.DingToken
	dingMSg.DingSecret = respSendConfig.DingSecret

	go GobalSend.acceptDing(&dingMSg)
	tools.JSONOk(c, "Start Sending")

}

func (d *DingChan) acceptDing(dm *DingMSg) {
	d.Dm <- dm
}

func (d *DingChan) ConsumerMsg() {
	for {
		select {
		case msg := <-d.Dm:
			dbSendDingHistory := db.SendDingHistory{}
			msgStr := fmt.Sprintf("### 主题: %v\n   %v  \n 时间: %v\n",
				msg.Data.Markdown.Title, msg.Data.Markdown.Text,
				time.Now().Format("2006-01-02 15:04:05"))
			dbSendDingHistory.Data = fmt.Sprintf("Title:[%v]  Text:[%v]", msg.Data.Markdown.Title, msg.Data.Markdown.Text)
			dbSendDingHistory.MsgId = msg.MsgId
			robot := dingtalk_robot.NewRobot(msg.DingToken, msg.DingSecret)
			err := robot.SendMarkdownMessage(msg.Data.Markdown.Title, msgStr, []string{}, false)

			if err != nil {
				dbSendDingHistory.Status = false
				dbSendDingHistory.ErrMsg = err.Error()
				return
			}
			dbSendDingHistory.Status = true
			dbSendDingHistory.Create()
			time.Sleep(time.Second * 3)
		default:
			time.Sleep(time.Second * 10)
		}
	}

}

func newSend() *DingChan {
	return &DingChan{
		Dm: make(chan *DingMSg),
	}
}
