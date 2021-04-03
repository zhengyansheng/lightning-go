package service

import (
	"encoding/json"
	"fmt"
	"lightning-go/pkg/request"
	"lightning-go/pkg/tools"
	"time"

	"github.com/douyu/jupiter/pkg/conf"
	"github.com/tidwall/gjson"
)

// 验证结构体
type DagRunDataSerializer struct {
	DagName    string                 `json:"dag_name" binding:"required"`
	Data       map[string]interface{} `json:"data" binding:"required"`
	CreateUser string                 `json:"create_user" binding:"required"`
}

type DagRunResponse struct {
	DagRunId        string `json:"dag_run_id"`
	StartDate       string `json:"start_date"`
	EndDate         string `json:"end_date"`
	ExecutionDate   string `json:"execution_date"`
	ExternalTrigger string `json:"external_trigger"`
	State           string `json:"state"`
}

/*

{
  "dags": [
    {
      "dag_id": "cron_demo",
      "description": "A cron demo test",
      "file_token": "Ii9yb290L2FpcmZsb3cvZGFncy9jcm9uX2RlbW8ucHki.ZEWrY1Ti5R8nGaywU-AcFmj6ixE",
      "fileloc": "/root/airflow/dags/cron_demo.py",
      "is_paused": true,
      "is_subdag": false,
      "owners": [
        "zhengshuai"
      ],
      "root_dag_id": null,
      "schedule_interval": {
        "__type": "CronExpression",
        "value": "@daily"
      },
      "tags": []
    },
    {
      "dag_id": "delivery_machine",
      "description": "交付机器",
      "file_token": "Ii9yb290L2FpcmZsb3cvZGFncy9kZWxpdmVyeV9tYWNoaW5lLnB5Ig.BI52o4f3FdqjUptIKk-4U1Lcc_Y",
      "fileloc": "/root/airflow/dags/delivery_machine.py",
      "is_paused": false,
      "is_subdag": false,
      "owners": [
        "zhengshuai"
      ],
      "root_dag_id": null,
      "schedule_interval": null,
      "tags": []
    },
    {
      "dag_id": "example_bash_operator",
      "description": null,
      "file_token": ".eJw9yjEOgCAMAMC_sEsHE79DilYgIm1aovJ7Jx0vOQfK3AGL7pXvQO2CWiLI6Jnb7Bew0mkSXA9MZN8DevCUSmHDZD8iWg4spNhZvQz3AmOfI1c.JFGlbq1mSnq865555vKdAlpQONo",
      "fileloc": "/root/airflow_env/lib/python3.6/site-packages/airflow/example_dags/example_bash_operator.py",
      "is_paused": true,
      "is_subdag": false,
      "owners": [
        "airflow"
      ],
      "root_dag_id": null,
      "schedule_interval": {
        "__type": "CronExpression",
        "value": "0 0 * * *"
      },
      "tags": [
        {
          "name": "example"
        },
        {
          "name": "example2"
        }
      ]
    },
*/

type DagDetailResponse struct {
	DagId            string              `json:"dag_id"`
	Description      string              `json:"description"`
	FileToken        string              `json:"file_token"`
	IsPaused         bool                `json:"is_paused"`
	IsSubdag         bool                `json:"is_subdag"`
	RootDagId        string              `json:"root_dag_id"`
	Fileloc          string              `json:"fileloc"`
	Owners           []string            `json:"owners"`
	Tags             []map[string]string `json:"tags"`
	ScheduleInterval interface{}         `json:"schedule_interval"`
}

type DagListResponse struct {
	Dags []DagDetailResponse `json:"dags"`
}

func GetAirflowHeader() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
		"Authorization": fmt.Sprintf("Basic %s", tools.Base64Encode(fmt.Sprintf("%s:%s",
			conf.GetString("go-ops.scheduler.airflowUserName"),
			conf.GetString("go-ops.scheduler.airflowUserPassword"),
		))),
	}
}

func (s DagRunDataSerializer) Trigger() (interface{}, error) {
	var (
		timeout     = time.Duration(time.Second * 10)
		body        = make(map[string]interface{})
		dagResponse DagRunResponse
	)

	dagUrl := fmt.Sprintf("%s/api/v1/dags/%s/dagRuns",
		conf.GetString("go-ops.scheduler.airflowUrl"),
		s.DagName,
	)
	header := map[string]string{
		"Content-Type": "application/json",
		"Authorization": fmt.Sprintf("Basic %s", tools.Base64Encode(fmt.Sprintf("%s:%s",
			conf.GetString("go-ops.scheduler.airflowUserName"),
			conf.GetString("go-ops.scheduler.airflowUserPassword"),
		))),
	}

	body["conf"] = s.Data
	tools.PrettyPrint(body)
	bodyByte, err := tools.JsonToByte(body)
	if err != nil {
		return "JsonToByte error", err
	}
	responseByte, err := request.Post(bodyByte, header, dagUrl, timeout)
	if err != nil {
		return "request.Post err", err
	}

	// 反序列化 response
	_ = json.Unmarshal(responseByte, &dagResponse)
	tools.PrettyPrint(dagResponse)
	return dagResponse, nil
}

type DagMgt struct {
	DagName  string `json:"dag_name"`
	DagRunId string `json:"dag_run_id"`
	Limit    int    `json:"limit"`
}

func (d DagMgt) ListDag() (DagListResponse, error) {
	var (
		dagUrl          string
		dagListResponse DagListResponse
	)
	d.Limit = 100
	dagUrl = fmt.Sprintf("%s/api/v1/dags?limit=%d",
		conf.GetString("go-ops.scheduler.airflowUrl"),
		d.Limit,
	)

	fmt.Printf("dagUrl %s\n", dagUrl)
	responseByte, err := request.GetWithHeader(dagUrl, GetAirflowHeader(), time.Duration(time.Second*5))
	if err != nil {
		return dagListResponse, err
	}
	// 反序列化 response
	_ = json.Unmarshal(responseByte, &dagListResponse)
	tools.PrettyPrint(dagListResponse)
	return dagListResponse, nil
}

func (d DagMgt) DetailDag() (DagDetailResponse, error) {
	var (
		dagUrl            string
		dagDetailResponse DagDetailResponse
	)
	dagUrl = fmt.Sprintf("%s/api/v1/dags/%s",
		conf.GetString("go-ops.scheduler.airflowUrl"),
		d.DagName,
	)

	fmt.Printf("dagUrl %s\n", dagUrl)
	responseByte, err := request.GetWithHeader(dagUrl, GetAirflowHeader(), time.Duration(time.Second*5))
	if err != nil {
		return dagDetailResponse, err
	}
	// 反序列化 response
	_ = json.Unmarshal(responseByte, &dagDetailResponse)
	tools.PrettyPrint(dagDetailResponse)
	return dagDetailResponse, nil
}

func (d DagMgt) TaskInstance() (map[string]interface{}, error) {
	var (
		dagUrl  string
		jsonMap = make(map[string]interface{})
	)
	d.Limit = 100
	dagUrl = fmt.Sprintf("%s/api/v1/dags/%s/dagRuns/%s/taskInstances?limit=%d",
		conf.GetString("go-ops.scheduler.airflowUrl"),
		d.DagName,
		d.DagRunId,
		d.Limit,
	)

	fmt.Printf("dagUrl %s\n", dagUrl)
	responseByte, err := request.GetWithHeader(dagUrl, GetAirflowHeader(), time.Duration(time.Second*5))
	if err != nil {
		return jsonMap, err
	}
	// 反序列化 response
	fmt.Println(string(responseByte))
	result := gjson.Get(string(responseByte), "task_instances")

	err = json.Unmarshal([]byte(result.Raw), &jsonMap)
	if err != nil {
		return jsonMap, err
	}
	return jsonMap, nil
}
