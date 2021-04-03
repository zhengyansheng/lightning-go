package scheduler

import "go-ops/internal/db"

// 创建云主机配置参数
type DagRun struct {
	ID                     uint   `gorm:"primary_key" json:"id"`
	DagId                  string `json:"dag_id"`
	ExecutionDate          string `json:"execution_date"`
	State                  string `json:"state"`
	RunId                  string `json:"run_id"`
	ExternalTrigger        int    `json:"external_trigger"`
	Conf                   string `json:"conf"`
	StartDate              string `json:"start_date"`
	EndDate                string `json:"end_date"`
	RunType                string `json:"run_type"`
	LastSchedulingDecision string `json:"last_scheduling_decision"`
	DagHash                string `json:"dag_hash"`
	CreatingJobId          int    `json:"creating_job_id"`
}

func (s *DagRun) TableName() string {
	return "dag_run"
}

func (s *DagRun) List() (dagRuns []DagRun, err error) {
	err = db.Airflow.Debug().Table(s.TableName()).Find(&dagRuns).Error
	return
}
