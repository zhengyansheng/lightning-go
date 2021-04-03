# 多云管理

## 模版

- 环境节点能够基本操作
    - 创建/删除模版，交付机器。
- 非环境节点只能查看模版
    - 多展示AppKey字段用来区分
    
    
## 交付

- 触发 Airflow Dag

```json
{
	"data":
	[
		{
			"id": 841,
			"dag_id": "delivery_machine",
			"execution_date": "2021-04-02 20:12:09.063224",
			"state": "success",
			"run_id": "manual__2021-04-02T12:12:09.063224+00:00",
			"external_trigger": 1,
			"conf": "BLOB",
			"end_date": "2021-04-02 20:12:19.226899",
			"start_date": "2021-04-02 20:12:09.072088",
			"run_type": "manual",
			"last_scheduling_decision": "2021-04-02 20:12:19.221751",
			"dag_hash": "74a3b0bbca17561f3a6ef1cc88a54039",
			"creating_job_id": null
		}
	]
}
```

```json
{
	"data":
	[
		{
			"task_id": "echo_hello_world",
			"dag_id": "cron_demo1",
			"execution_date": "2021-04-02 20:50:00.000000",
			"start_date": "2021-04-02 21:00:00.811357",
			"end_date": "2021-04-02 21:00:00.968039",
			"duration": 0.156682,
			"state": "success",
			"try_number": 1,
			"hostname": "zhengshuai",
			"unixname": "root",
			"job_id": 879,
			"pool": "default_pool",
			"queue": "default",
			"priority_weight": 1,
			"operator": "BashOperator",
			"queued_dttm": "2021-04-02 21:00:00.650318",
			"pid": 14216,
			"max_tries": 0,
			"executor_config": "BLOB",
			"pool_slots": 1,
			"queued_by_job_id": 489,
			"external_executor_id": "c6fca705-41c9-4349-bd58-f48156ee0f72"
		}
	]
}
```