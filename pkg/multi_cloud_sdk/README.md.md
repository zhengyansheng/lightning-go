# create instance

## ali create instance
```json

{
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
	  "hostname": "lightning-fe1.ops.prod.ali"
  }

```


## ten create instance
```json

{
	  "account": "ten.lightning",
	  "pay_type": "PostPaid",
	  "region_id": "ap-shanghai",
	  "zone_id": "ap-shanghai-2",
	  "instance_type": "S4.MEDIUM4",
	  "image_id": "img-oikl1tzv",
	  "vpc_id": "vpc-2qckup44",
	  "subnet_id": "subnet-r1tor2m1",
	  "disks": [{
		  "disk_type": "system",
		  "disk_storage_type": "LOCAL_SSD",
		  "disk_storage_size": 100
	  }],
	  "security_group_ids": ["sg-qodmlw0w"],
	  "hostname": "lightning-fe1.ops.prod.ten"
  }

```