package multi_cloud_sdk

import (
	"errors"
	"fmt"
	"lightning-go/pkg/tools"
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

// create and start
func (ali *aliyunClient) CreateInstance(PayType, hostname, instanceType, zoneId, imageId, vpcId, subnetId string, securityGroupIds []string, dryRun bool) (instanceInfo map[string]string, err error) {
	/*
		实例的付费方式。取值范围：
			PrePaid：包年包月
			PostPaid（默认）：按量付费
			选择包年包月时，您必须确认自己的账号支持余额支付或者信用支付，否则将返回InvalidPayMethod的错误提示。
	*/
	instanceInfo = make(map[string]string)
	runInstancesRequest := &ecs.RunInstancesRequest{
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

	if runInstancesRequest.InstanceChargeType == tea.String("PrePaid") {
		// 包年包月 特殊设置
		runInstancesRequest.PeriodUnit = tea.String("Month")
		runInstancesRequest.Period = tea.Int32(1)
		runInstancesRequest.AutoRenew = tea.Bool(true)
		runInstancesRequest.AutoRenewPeriod = tea.Int32(1)
	}
	response, err := ali.ecsClt.RunInstances(runInstancesRequest)
	if err != nil {
		return instanceInfo, err
	}
	instanceIdSet := response.Body.InstanceIdSets.InstanceIdSet
	if len(instanceIdSet) == 0 {
		return instanceInfo, errors.New("Request ok, but instance create not success. ")
	}
	instanceInfo["instance_id"] = tea.StringValue(instanceIdSet[0])
	return instanceInfo, nil
}

// create
func (ali *aliyunClient) CreateInstanceButStop(PayType, hostname, instanceType, zoneId, imageId, vpcId, subnetId string, securityGroupIds []string, dryRun bool) (instanceId string, err error) {
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
	// 如果机器状态是关机，则无法重启；必须先开机 在关机才行
	rebootInstanceRequest := &ecs.RebootInstanceRequest{
		InstanceId: tea.String(instanceId),
		ForceStop:  tea.Bool(forceStop),
	}
	response, err := ali.ecsClt.RebootInstance(rebootInstanceRequest)
	if err != nil {
		return err.Error(), err
	}
	requestId := response.Body.RequestId
	return *requestId, nil
}

func (ali *aliyunClient) DestroyInstance(instanceId string, forceStop bool) (string, error) {
	deleteInstanceRequest := &ecs.DeleteInstanceRequest{
		InstanceId:            tea.String(instanceId),
		Force:                 tea.Bool(forceStop),
		TerminateSubscription: tea.Bool(true), // 是否强制释放运行中（Running的实例
	}
	response, err := ali.ecsClt.DeleteInstance(deleteInstanceRequest)
	if err != nil {
		return err.Error(), err
	}
	return *response.Body.RequestId, nil
}

func (ali *aliyunClient) ModifyInstanceName(instanceId string, instanceName string) (string, error) {
	modifyInstanceAttributeRequest := &ecs.ModifyInstanceAttributeRequest{
		InstanceId:   tea.String(instanceId),
		HostName:     tea.String(instanceName),
		InstanceName: tea.String(instanceName),
	}
	response, err := ali.ecsClt.ModifyInstanceAttribute(modifyInstanceAttributeRequest)
	if err != nil {
		return err.Error(), err
	}
	return *response.Body.RequestId, nil
}

func (ali *aliyunClient) ListInstance(instanceId string) (map[string]interface{}, error) {
	var instanceInfo = make(map[string]interface{})
	describeInstancesRequest := &ecs.DescribeInstancesRequest{
		PageSize:    tea.Int32(100),
		InstanceIds: tea.String(instanceId),
		RegionId:    tea.String(ali.regionId),
		DryRun:      tea.Bool(false),
	}
	instances, err := ali.ecsClt.DescribeInstances(describeInstancesRequest)
	if err != nil {
		return instanceInfo, err
	}
	if len(instances.Body.Instances.Instance) == 0 {
		return instanceInfo, errors.New(fmt.Sprintf("%s not found ", instanceId))
	}
	for _, instance := range instances.Body.Instances.Instance {
		instanceInfo, _ = ali.processInstance(instance)

	}
	return instanceInfo, nil
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
		for _, instance := range instances.Body.Instances.Instance {
			info, _ := ali.processInstance(instance)
			instanceArray = append(instanceArray, info)

		}
	}
	return instanceArray, nil
}

func (ali *aliyunClient) processInstance(instance *ecs.DescribeInstancesResponseBodyInstancesInstance) (map[string]interface{}, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover %s\n", err)
		}
	}()

	tools.PrettyPrint(instance)

	var publicIpAddress *string
	info := make(map[string]interface{})

	if len(instance.PublicIpAddress.IpAddress) != 0 {
		publicIpAddress = instance.PublicIpAddress.IpAddress[0]
	}
	info = map[string]interface{}{
		"type":                 cloudType,
		"platform":             "ali",
		"instance_charge_type": instance.InstanceChargeType,
		"instance_id":          instance.InstanceId,
		"private_ip":           instance.VpcAttributes.PrivateIpAddress.IpAddress[0],
		"public_ip":            publicIpAddress,
		"eip_ip":               instance.EipAddress.IpAddress,
		//"instance_name":        instance.InstanceName,
		"hostname":             instance.InstanceName,
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
		"start_time":           tools.ReplaceDateTime(*instance.StartTime),
		"create_time":          tools.ReplaceDateTime(*instance.CreationTime),
		"expired_time":         tools.ReplaceDateTime(*instance.ExpiredTime),
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
