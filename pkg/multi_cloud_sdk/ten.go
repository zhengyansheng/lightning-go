package multi_cloud_sdk

import (
	"errors"
	"fmt"
	"lightning-go/pkg/tools"
	"strings"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type tencentcloudClient struct {
	regionId string
	cvmClt   *cvm.Client
}

func NewTenCvm(accessKeyId, accessKeySecret, regionId string) (*tencentcloudClient, error) {
	credential := common.NewCredential(
		accessKeyId,
		accessKeySecret,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	clt, err := cvm.NewClient(credential, regionId, cpf)
	if err != nil {
		return &tencentcloudClient{}, errors.New("")
	}
	return &tencentcloudClient{
		regionId: regionId,
		cvmClt:   clt,
	}, nil
}

func (ten *tencentcloudClient) CreateInstance(instanceChargeType, hostname, instanceType, zoneId, imageId, vpcId, subnetId string, securityGroupIds []string, dryRun bool) (instanceInfo map[string]string, err error) {
	//fmt.Println(hostname, instanceType, zoneId, imageId, dryRun)
	/*
		instanceChargeType 实例付费类型
			PREPAID：预付费，即包年包月
			POSTPAID_BY_HOUR：按小时后付费
			CDHPAID：独享子机（基于专用宿主机创建，宿主机部分的资源不收费）
			SPOTPAID：竞价付费
	*/
	instanceInfo = make(map[string]string)
	request := cvm.NewRunInstancesRequest()

	if instanceChargeType == "PostPaid" {
		request.InstanceChargeType = common.StringPtr("POSTPAID_BY_HOUR")
	} else if instanceChargeType == "PrePaid" {
		request.InstanceChargeType = common.StringPtr("PREPAID")
		request.InstanceChargePrepaid = &cvm.InstanceChargePrepaid{
			Period:    common.Int64Ptr(1),
			RenewFlag: common.StringPtr("NOTIFY_AND_AUTO_RENEW"),
		}
	} else {
		return instanceInfo, errors.New("instanceChargeType only support PostPaid or PrePaid")
	}
	request.InstanceType = common.StringPtr(instanceType)
	request.Placement = &cvm.Placement{
		Zone: common.StringPtr(zoneId),
	}
	request.InstanceCount = common.Int64Ptr(1)
	request.InstanceName = common.StringPtr(hostname)
	request.LoginSettings = &cvm.LoginSettings{
		Password: common.StringPtr(password),
	}
	request.VirtualPrivateCloud = &cvm.VirtualPrivateCloud{
		VpcId:    common.StringPtr(vpcId),
		SubnetId: common.StringPtr(subnetId),
	}
	request.HostName = common.StringPtr(hostname)
	request.ImageId = common.StringPtr(imageId)
	request.SecurityGroupIds = common.StringPtrs(securityGroupIds)
	request.SystemDisk = &cvm.SystemDisk {
		DiskType: common.StringPtr("CLOUD_PREMIUM"),
		DiskSize: common.Int64Ptr(100),
	}
	//request.DataDisks = []*cvm.DataDisk {
	//	&cvm.DataDisk {
	//		DiskType: common.StringPtr("LOCAL_SSD"),
	//		DiskSize: common.Int64Ptr(100),
	//		DeleteWithInstance: common.BoolPtr(true),
	//	},
	//}
	request.DryRun = common.BoolPtr(dryRun)

	response, err := ten.cvmClt.RunInstances(request)
	if err != nil {
		return instanceInfo, err
	}
	fmt.Printf("%s", response.ToJsonString())
	fmt.Println(response.Response.InstanceIdSet)

	instanceIdSet := response.Response.InstanceIdSet
	//requestId := response.Response.RequestId
	if len(instanceIdSet) != 1 {
		return instanceInfo, err
	} else {
		instanceInfo["instance_id"] = common.StringValues(instanceIdSet)[0]
		return instanceInfo, nil
	}
}

func (ten *tencentcloudClient) StartInstance(instanceId string) (string, error) {
	request := cvm.NewStartInstancesRequest()

	request.InstanceIds = common.StringPtrs([]string{instanceId})

	response, err := ten.cvmClt.StartInstances(request)
	if err != nil {
		return "An API error", err
	}
	return *response.Response.RequestId, nil
}

func (ten *tencentcloudClient) StopInstance(instanceId string) (string, error) {
	request := cvm.NewStopInstancesRequest()

	request.InstanceIds = common.StringPtrs([]string{instanceId})
	request.ForceStop = common.BoolPtr(true)
	//request.StopType = common.StringPtr("HARD")
	request.StoppedMode = common.StringPtr("STOP_CHARGING")

	response, err := ten.cvmClt.StopInstances(request)
	if err != nil {
		return "An API error", err
	}
	return *response.Response.RequestId, nil
}

func (ten *tencentcloudClient) RebootInstance(instanceId string, forceStop bool) (string, error) {
	request := cvm.NewRebootInstancesRequest()

	request.InstanceIds = common.StringPtrs([]string{instanceId})
	//request.ForceReboot = common.BoolPtr(true)
	if forceStop {
		request.StopType = common.StringPtr("HARD")
	} else {
		request.StopType = common.StringPtr("SOFT")
	}

	response, err := ten.cvmClt.RebootInstances(request)
	if err != nil {
		return "An API error", err
	}
	return *response.Response.RequestId, nil
}

func (ten *tencentcloudClient) DestroyInstance(instanceId string, forceStop bool) (string, error) {
	request := cvm.NewTerminateInstancesRequest()

	request.InstanceIds = common.StringPtrs([]string{instanceId})

	response, err := ten.cvmClt.TerminateInstances(request)
	if err != nil {
		return "An API error", err
	}
	return *response.Response.RequestId, nil
}

func (ten *tencentcloudClient) ModifyInstanceName(instanceId string, instanceName string) (string, error) {
	request := cvm.NewModifyInstancesAttributeRequest()

	request.InstanceIds = common.StringPtrs([]string{instanceId})
	request.InstanceName = common.StringPtr(instanceName)

	response, err := ten.cvmClt.ModifyInstancesAttribute(request)
	if err != nil {
		return "An API error", err
	}
	return *response.Response.RequestId, nil
}

func (ten *tencentcloudClient) ModifyInstanceSecurityGroups(instanceId string, securityGroupIds []string) (string, error) {
	request := cvm.NewModifyInstancesAttributeRequest()

	request.InstanceIds = common.StringPtrs([]string{instanceId})
	request.SecurityGroups = common.StringPtrs(securityGroupIds)

	response, err := ten.cvmClt.ModifyInstancesAttribute(request)
	if err != nil {
		return "An API error", err
	}
	fmt.Printf("%s", response.ToJsonString())
	return "", nil
}

func (ten *tencentcloudClient) ResetInstance() ([]map[string]string, error) {
	return nil, nil
}

func (ten *tencentcloudClient) ListInstance(instanceId string) (map[string]interface{}, error) {
	var instanceInfo = make(map[string]interface{})
	request := cvm.NewDescribeInstancesRequest()
	request.InstanceIds = common.StringPtrs([]string{instanceId})
	response, err := ten.cvmClt.DescribeInstances(request)
	if err != nil {
		return instanceInfo, err
	}
	// 如果返回结果集中为空 则退出循环
	if len(response.Response.InstanceSet) == 0 {
		return instanceInfo, errors.New("Response instanceSet count 0 ")
	}
	info, err := ten.processInstance(response.Response.InstanceSet[0])
	if err != nil {
		return instanceInfo, err
	}
	instanceInfo = info
	return instanceInfo, nil
}

func (ten *tencentcloudClient) ListInstances() ([]map[string]interface{}, error) {
	var instanceArray []map[string]interface{}
	for pageNumber := 0; pageNumber <= 3; pageNumber++ {
		request := cvm.NewDescribeInstancesRequest()
		request.Offset = common.Int64Ptr(int64(pageNumber)*100)
		request.Limit = common.Int64Ptr(100)
		response, err := ten.cvmClt.DescribeInstances(request)
		if err != nil {
			return instanceArray, err
		}
		// 如果返回结果集中为空 则退出循环
		if len(response.Response.InstanceSet) == 0 {
			break
		}
		for _, instance := range response.Response.InstanceSet {
			info, err := ten.processInstance(instance)
			if err == nil {
				instanceArray = append(instanceArray, info)
			}
		}
	}
	tools.PrettyPrint(instanceArray)
	return instanceArray, nil
}

func (ten *tencentcloudClient) processInstance(instance *cvm.Instance) (map[string]interface{}, error) {
	info := make(map[string]interface{})
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recover %s\n", err)
		}
	}()

	tools.PrettyPrint(instance)
	info = map[string]interface{}{
		"type":                 cloudType,
		"platform":             "ten",
		"instance_charge_type": instance.InstanceChargeType,
		"instance_id":          instance.InstanceId,
		"private_ip":           instance.PrivateIpAddresses[0],
		"public_ip":            instance.PublicIpAddresses,
		"eip_ip":               "",
		"instance_name":        instance.InstanceName,
		"hostname":             instance.InstanceName,
		"image_id":             instance.ImageId,
		"os_system":            strings.Fields(*instance.OsName)[0],
		"os_version":           instance.OsName,
		"instance_type":        instance.InstanceType,
		"region_id":            ten.regionId,
		"zone_id":              instance.Placement.Zone,
		"vpc_id":               instance.VirtualPrivateCloud.VpcId,
		"subnet_id":            instance.VirtualPrivateCloud.SubnetId,
		"security_group_ids":   instance.SecurityGroupIds,
		"state":                strings.ToLower(*instance.InstanceState),
		"mem_total":            *instance.Memory,
		"cpu_total":            instance.CPU,
		"start_time":           "",
		"create_time":          tools.ReplaceDateTime(*instance.CreatedTime),
		"expired_time":         tools.ReplaceDateTime(*instance.ExpiredTime),
	}
	return info, nil
}

func (ten *tencentcloudClient) ListRegions() (regions []map[string]string, err error) {
	request := cvm.NewDescribeRegionsRequest()
	response, err := ten.cvmClt.DescribeRegions(request)
	if err != nil {
		return regions, err
	}
	for _, v := range response.Response.RegionSet {
		m := make(map[string]string)
		m["region_id"] = *v.Region
		m["local_name"] = *v.RegionName
		regions = append(regions, m)
	}
	return regions, nil
}
