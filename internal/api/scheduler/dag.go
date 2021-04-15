package scheduler

import (
	"fmt"
	"io/ioutil"
	"lightning-go/internal/models/scheduler"
	"lightning-go/internal/service/scheduler"
	"lightning-go/pkg/tools"

	"strings"

	"github.com/gin-gonic/gin"
)

// 触发Dag
func TriggerDagRun(c *gin.Context) {
	// 接收所有参数 不做验证
	var dagRun service.DagRunDataSerializer
	dagRun.DagName = c.Param("dagName") // /:id
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	paramData, err := tools.ByteToJson(bytes)
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	dagRun.Data = paramData
	msg, err := dagRun.Trigger()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	tools.JSONOk(c, msg)
}

// 触发Dag
func TriggerDagRunV2(c *gin.Context) {
	// Validate
	var dagRun service.DagRunDataSerializer
	if err := c.ShouldBindJSON(&dagRun); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	dagRun.DagName = c.Param("dagName") // /:id
	tools.PrettyPrint(dagRun)
	msg, err := dagRun.Trigger()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	tools.JSONOk(c, msg)
}

// 查询Dag
// /api/v1/task-scheduler/dag
// /api/v1/task-scheduler/dag?dag_id=delivery_machine&execution_date=2021-04-02 20:12:09.063224
func ListDagRun(c *gin.Context) {
	var (
		dagRun       scheduler.DagRun
		taskInstance scheduler.TaskInstance
	)
	// replace
	// https://blog.csdn.net/Yvken_Zh/article/details/104861765
	c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, "+", "%2b")

	// Get params
	dagId := c.Query("dag_id") // ?dag_id=xxx
	executionDate := c.Query("execution_date")
	fmt.Printf("dagId: %v, executionDate: %v\n", dagId, executionDate)
	if dagId != "" && executionDate != "" {
		// execute sql
		taskInstance.DagId = dagId
		taskInstance.ExecutionDate = executionDate
		dagRuns, err := taskInstance.ListByDagAndExecDate()
		if err != nil {
			tools.JSONFailed(c, tools.MSG_ERR, err.Error())
			return
		}
		// Response
		tools.JSONOk(c, dagRuns)
	} else {
		// execute sql
		dagRuns, err := dagRun.List()
		if err != nil {
			tools.JSONFailed(c, tools.MSG_ERR, err.Error())
			return
		}
		// Response
		tools.JSONOk(c, dagRuns)
	}
	return
}

func ListTaskLog(c *gin.Context) {

}
