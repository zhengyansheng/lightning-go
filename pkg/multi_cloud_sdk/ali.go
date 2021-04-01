package multi_cloud_sdk

import (
	"errors"
	"fmt"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	ecs "github.com/alibabacloud-go/ecs-20140526/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

type aliyunClient struct {
	regionId string
	ecsClt   *ecs.Client
}

func NewAliEcs(accessKeyId, accessKeySecret, regionId string) (*aliyunClient, error) {
	openCfg := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyId),
		AccessKeySecret: tea.String(accessKeySecret),
		RegionId:        tea.String(regionId),
		//ConnectTimeout:  tea.Int(10),
		//ReadTimeout:     tea.Int(10),
	}
	clt, err := ecs.NewClient(openCfg)
	if err != nil {
		return &aliyunClient{}, err
	}
	return &aliyunClient{
		regionId: regionId,
		ecsClt:   clt,
	}, nil
}

func (ali *aliyunClient) CreateInstance(PayType, hostname, instanceType, zoneId, imageId, vpcId, subnetId string, securityGroupIds []string, dryRun bool) (instanceId string, err error) {
	/*
		实例的付费方式。取值范围：
			PrePaid：包年包月
			PostPaid（默认）：按量付费
			选择包年包月时，您必须确认自己的账号支持余额支付或者信用支付，否则将返回InvalidPayMethod的错误提示。
	*/
	request := &ecs.CreateInstanceRequest{
		InstanceChargeType: tea.String(PayType),
		InstanceName:       tea.String(hostname),
		Password:           tea.String(password),
		InstanceType:       tea.String(instanceType),
		RegionId:           tea.String(ali.regionId),
		ZoneId:             tea.String(zoneId),
		ImageId:            tea.String(imageId),
		DryRun:             tea.Bool(dryRun),
		VSwitchId:          tea.String(subnetId),
		SecurityGroupId:    tea.String(securityGroupIds[0]),
	}

	if request.InstanceChargeType == tea.String("PrePaid") {
		// 包年包月 特殊设置
		request.PeriodUnit = tea.String("Month")
		request.Period = tea.Int32(1)
		request.AutoRenew = tea.Bool(true)
		request.AutoRenewPeriod = tea.Int32(1)
	}
	response, err := ali.ecsClt.CreateInstance(request)
	if err != nil {
		return "Call RunInstances error", err
	}

	return tea.StringValue(response.Body.InstanceId), nil
}

func (ali *aliyunClient) StartInstance(instanceId string) (string, error) {
	startInstancesRequest := &ecs.StartInstancesRequest{
		RegionId:          tea.String(ali.regionId),
		BatchOptimization: tea.String("AllTogether"),
		InstanceId:        []*string{tea.String(instanceId)},
	}
	response, err := ali.ecsClt.StartInstances(startInstancesRequest)
	if err != nil {
		return err.Error(), err
	}
	for _, v := range response.Body.InstanceResponses.InstanceResponse {
		if *v.Message == "success" && *v.Code == "200" {
			return *response.Body.RequestId, nil
		}
	}
	return "", errors.New("Start fail. ")
}

func (ali *aliyunClient) StopInstance(instanceId string) (string, error) {
	stopInstancesRequest := &ecs.StopInstancesRequest{
		RegionId:          tea.String(ali.regionId),
		BatchOptimization: tea.String("AllTogether"),
		InstanceId:        []*string{tea.String(instanceId)},
		ForceStop:         tea.Bool(true),
	}
	response, err := ali.ecsClt.StopInstances(stopInstancesRequest)
	if err != nil {
		return err.Error(), err
	}

	for _, v := range response.Body.InstanceResponses.InstanceResponse {
		if *v.Message == "success" && *v.Code == "200" {
			return *response.Body.RequestId, nil
		}
	}
	return "", errors.New("Stop fail. ")
}

func (ali *aliyunClient) RebootInstance(instanceId string, forceStop bool) (string, error) {
	rebootInstanceRequest := &ecs.RebootInstanceRequest{
		InstanceId: tea.String(instanceId),
		ForceStop:  tea.Bool(forceStop),
	}
	response, err := ali.ecsClt.RebootInstance(rebootInstanceRequest)
	if err != nil {
		return err.Error(), err
	}
	requestId := response.Body.RequestId
	return fmt.Sprintf("requestId: %s", *requestId), nil
}

func (ali *aliyunClient) ListInstance(instanceId string) (map[string]interface{}, error) {
	return nil, nil
}

func (ali *aliyunClient) ListInstances() ([]map[string]interface{}, error) {
	var instanceArray []map[string]interface{}
	for pageNumber := 1; pageNumber <= 3; pageNumber++ {
		describeInstancesRequest := &ecs.DescribeInstancesRequest{
			PageSize:   tea.Int32(100),
			PageNumber: tea.Int32(int32(pageNumber)),
			RegionId:   tea.String(ali.regionId),
			DryRun:     tea.Bool(false),
		}
		// 复制代码运行请自行打印 API 的返回值
		instances, err := ali.ecsClt.DescribeInstances(describeInstancesRequest)
		if err != nil {
			return instanceArray, err
		}
		// 如果返回结果集中为空 则退出循环
		if len(instances.Body.Instances.Instance) == 0 {
			break
		}
		//fmt.Println(instances.Body.Instances.Instance)
		//fmt.Println("->", *instances.Body.TotalCount)
		//fmt.Println("->", len(instances.Body.Instances.Instance))
		for _, instance := range instances.Body.Instances.Instance {
			info, _ := ali.processInstance(instance)
			instanceArray = append(instanceArray, info)

		}
	}
	return instanceArray, nil
}

func (ali *aliyunClient) processInstance(instance *ecs.DescribeInstancesResponseBodyInstancesInstance) (map[string]interface{}, error) {
	info := make(map[string]interface{})
	info = map[string]interface{}{
		"type":                 cloudType,
		"platform":             "ali",
		"instance_charge_type": instance.InstanceChargeType,
		"instance_id":          instance.InstanceId,
		"private_ip":           instance.VpcAttributes.PrivateIpAddress.IpAddress[0],
		"public_ip":            instance.PublicIpAddress.IpAddress[0],
		"eip_ip":               instance.EipAddress.IpAddress,
		"instance_name":        instance.InstanceName,
		"hostname":             instance.HostName,
		"image_id":             instance.ImageId,
		"os_system":            instance.OSType,
		"os_version":           instance.OSName,
		"instance_type":        instance.InstanceType,
		"region_id":            instance.RegionId,
		"zone_id":              instance.ZoneId,
		"vpc_id":               instance.VpcAttributes.VpcId,
		"subnet_id":            instance.VpcAttributes.VSwitchId,
		"security_group_ids":   instance.SecurityGroupIds.SecurityGroupId,
		"state":                strings.ToLower(*instance.Status),
		"mem_total":            *instance.Memory / 1024,
		"cpu_total":            instance.Cpu,
		"start_time":           instance.StartTime,
		"create_time":          instance.CreationTime,
		"expired_time":         instance.ExpiredTime,
	}
	return info, nil
}

func (ali *aliyunClient) ListRegions() ([]map[string]string, error) {
	var regions []map[string]string
	describeRegionsRequest := &ecs.DescribeRegionsRequest{
		AcceptLanguage: tea.String("zh-CN"),
		ResourceType:   tea.String("instance"),
	}
	// 复制代码运行请自行打印 API 的返回值
	regionInfo, err := ali.ecsClt.DescribeRegions(describeRegionsRequest)
	if err != nil {
		return regions, err
	}
	for _, reg := range regionInfo.Body.Regions.Region {
		info := map[string]string{
			"region_id":  *reg.RegionId,
			"local_name": *reg.LocalName,
		}
		regions = append(regions, info)
	}
	return regions, nil
}
