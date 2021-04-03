package scheduler

import "lightning-go/internal/db"

// 创建云主机配置参数
type TaskInstance struct {
	TaskId             string  `json:"task_id"`
	DagId              string  `json:"dag_id" binding:"required"`
	ExecutionDate      string  `json:"execution_date" binding:"required"`
	StartDate          string  `json:"start_date"`
	EndDate            string  `json:"end_date"`
	Duration           float64 `json:"duration"`
	State              string  `json:"state"`
	TryNumber          int     `json:"try_number"`
	Hostname           string  `json:"hostname"`
	Unixname           string  `json:"unixname"`
	JobId              int     `json:"job_id"`
	Pool               string  `json:"pool"`
	Queue              string  `json:"queue"`
	PriorityWeight     string  `json:"priority_weight"`
	Operator           string  `json:"operator"`
	QueuedDttm         string  `json:"queued_dttm"`
	Pid                int     `json:"pid"`
	MaxTries           int     `json:"max_tries"`
	ExecutorConfig     string  `json:"executor_config"`
	PoolSlots          int     `json:"pool_slots"`
	QueuedByJobId      int     `json:"queued_by_job_id"`
	ExternalExecutorId string  `json:"external_executor_id"`
}

func (s *TaskInstance) TableName() string {
	return "task_instance"
}

func (s *TaskInstance) ListByDagAndExecDate() (tasInstances []TaskInstance, err error) {
	table := db.Airflow.Debug().Table(s.TableName())
	err = table.Where("dag_id = ? AND execution_date = ?", s.DagId, s.ExecutionDate).Find(&tasInstances).Error
	return
}
