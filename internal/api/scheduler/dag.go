package scheduler

import (
	"go-ops/internal/service/scheduler"
	"go-ops/pkg/tools"
	"strings"

	"github.com/gin-gonic/gin"
)

func TriggerDagRun(c *gin.Context) {
	var dagRun service.DagRunDataSerializer
	if err := c.ShouldBindJSON(&dagRun); err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	tools.PrettyPrint(dagRun)
	msg, err := dagRun.Trigger()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	tools.JSONOk(c, msg)
}

func ListDagRun(c *gin.Context) {
	var (
		err     error
		dag     service.DagMgt
		dagData interface{}
	)
	dagName, ok := c.GetQuery("dag_name")
	if ok {
		dag.DagName = dagName
		dagData, err = dag.DetailDag()
	} else {
		dagData, err = dag.ListDag()
	}
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	tools.JSONOk(c, dagData)
}

func TaskInstance(c *gin.Context) {
	// ?dag_name=xxx&dag_run_id=xxx
	// https://blog.csdn.net/Yvken_Zh/article/details/104861765
	var (
		ok       bool
		dagName  string
		dagRunId string
		dag      service.DagMgt
	)
	c.Request.URL.RawQuery = strings.ReplaceAll(c.Request.URL.RawQuery, "+", "%2b")
	dagName, ok = c.GetQuery("dag_name")
	if !ok {
		tools.JSONFailed(c, tools.MSG_ERR, "dag_name is required")
		return
	}

	dagRunId, ok = c.GetQuery("dag_run_id")

	if !ok {
		tools.JSONFailed(c, tools.MSG_ERR, "dag_run_id is required")
		return
	}
	dag.DagName = dagName
	dag.DagRunId = dagRunId
	tasks, err := dag.TaskInstance()
	if err != nil {
		tools.JSONFailed(c, tools.MSG_ERR, err.Error())
		return
	}
	tools.JSONOk(c, tasks)

}
