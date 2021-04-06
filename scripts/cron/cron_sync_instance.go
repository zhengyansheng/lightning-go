package main

import (
	"encoding/json"
	"fmt"
	"lightning-go/pkg/request"
	"lightning-go/pkg/tools"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

var (
	multiCloudHost = "http://127.0.0.1:9900"
	//opsHost        = "http://127.0.0.1:9000"
	opsHost       = "http://ops.aiops724.com"
	regionUri     = "/api/v1/multi-cloud/regions/"
	instanceUri   = "/api/v1/multi-cloud/instance/"
	cmdbUri       = "/api/v1/cmdb/instances/"
	cmdbUpdateUri = "/api/v1/cmdb/instances/multi_update/"
	accounts      = []string{
		"ali.lightning",
	}
)

func getCloudPlatRegions(account string) (m []map[string]string, err error) {
	var regionId string
	accountArray := strings.Split(account, ".")
	switch accountArray[0] {
	case "ali":
		regionId = "cn-beijing"
	case "ten":
		regionId = "ap-beijing"
	case "aws":
		regionId = "cn-north-1"
	}
	url := fmt.Sprintf("%s%s?account=%s&region_id=%s", multiCloudHost, regionUri, account, regionId)
	dataByte, err := request.Get(url, time.Duration(time.Second*10))
	if err != nil {
		return
	}

	if gjson.Get(string(dataByte), "code").Num == 0 {
		dataArray := gjson.Get(string(dataByte), "data").Array()
		for _, region := range dataArray {
			regionMap := make(map[string]string)
			regionMap["local_name"] = region.Map()["local_name"].Str
			regionMap["region_id"] = region.Map()["region_id"].Str
			m = append(m, regionMap)
		}
		return
	}
	return
}

type response struct {
	Code    int                      `json:"code"`
	Message string                   `json:"message"`
	Data    []map[string]interface{} `json:"data"`
}

func getInstances(account, regionId string) (m []map[string]interface{}, err error) {
	url := fmt.Sprintf("%s%s?account=%s&region_id=%s", multiCloudHost, instanceUri, account, regionId)
	dataByte, err := request.Get(url, time.Duration(time.Second*10))
	if err != nil {
		return
	}

	var res response
	err = json.Unmarshal(dataByte, &res)
	m = res.Data
	return
}

func postCmdb(instances map[string]interface{}) (string, error) {
	url := fmt.Sprintf("%s%s", opsHost, cmdbUri)
	fmt.Printf("url:%s\n", url)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	dataByte, err := json.Marshal(instances)
	if err != nil {
		return "json.Marshal err", err
	}
	dataByte, err = request.Post(dataByte, headers, url, time.Duration(time.Second*10))
	if err != nil {
		return string(dataByte), err
	}
	return "", nil
}

func UpdateCmdb(instances []map[string]interface{}) (string, error) {
	url := fmt.Sprintf("%s%s", opsHost, cmdbUpdateUri)
	fmt.Printf("url:%s\n", url)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	dataByte, err := json.Marshal(instances)
	if err != nil {
		return "json.Marshal err", err
	}
	dataByte, err = request.Put(dataByte, headers, url, time.Duration(time.Second*10))
	if err != nil {
		return string(dataByte), err
	}
	return "", nil
}

func main() {
	for _, account := range accounts {
		//1. 获取平台下所有的地域
		regionMap, err := getCloudPlatRegions(account)
		if err != nil {
			fmt.Println("getCloudPlatRegions err", err)
			return
		}

		//2. 获取云主机信息
		for _, info := range regionMap {
			instances, err := getInstances(account, info["region_id"])
			if err != nil {
				fmt.Println("getInstances err", err)
				continue
			}
			if len(instances) == 0 {
				fmt.Printf("region_id: %s instances not found.\n", info["region_id"])
				continue
			}
			for _, instanceInfo := range instances {
				//3. 同步云主机信息
				instanceInfo["account"] = account
				tools.PrettyPrint(instanceInfo)
				res, err := postCmdb(instanceInfo)
				if err != nil {
					fmt.Printf("Post cmdb err: %v, response: %v\n", err, res)
				}
				//4. 变更云主机信息到cmdb
				instanceArr := []map[string]interface{}{}
				instanceArr = append(instanceArr, instanceInfo)
				res, err = UpdateCmdb(instanceArr)
				if err != nil {
					//fmt.Printf("Update cmdb err: %v, response: %v\n", err, res)
				}
			}
		}

	}

}
