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
	"data": {
	  "account": "ali.lightning",
	  "pay_type": "PostPaid",
	  "region_id": "cn-beijing",
	  "zone_id": "cn-beijing-c",
	  "instance_type": "ecs.sn1.medium",
	  "image_id": "centos_7_8_x64_20G_alibase_20200914.vhd",
	  "vpc_id": "vpc-2ze80et76jwcsuc2asobq",
	  "subnet_id": "vsw-2zelzctt7aougw3bhz5ev",
	  "disks": [{
		  "disk_type": "system",
		  "disk_storage_type": "cloud",
		  "disk_storage_size": 100
	  }],
	  "security_group_ids": ["sg-2zeb9n0qzygmjwn7hwae"],
	  "hostname": "lightning-fe2.ops.prod.ali"
  }
}
```