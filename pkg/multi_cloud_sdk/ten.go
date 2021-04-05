package multi_cloud_sdk

import (
	"errors"
	"fmt"

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
		Password: common.StringPtr("fafa8989afjak@9a"),
	}
	request.VirtualPrivateCloud = &cvm.VirtualPrivateCloud{
		VpcId:    common.StringPtr(vpcId),
		SubnetId: common.StringPtr(subnetId),
	}
	request.HostName = common.StringPtr(hostname)
	request.ImageId = common.StringPtr(imageId)
	request.SecurityGroupIds = common.StringPtrs(securityGroupIds)
	request.DryRun = common.BoolPtr(dryRun)

	response, err := ten.cvmClt.RunInstances(request)
	//if _, ok := err.(*errors.tcErrors); ok {
	//	return fmt.Sprintf("An API error has returned: %s", err), err
	//}
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
	fmt.Printf("%s", response.ToJsonString())
	return "", nil
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
	fmt.Printf("%s", response.ToJsonString())
	return "", nil
}

func (ten *tencentcloudClient) RebootInstance(instanceId string, forceStop bool) (string, error) {
	return "", nil
}

func (ten *tencentcloudClient) DestroyInstance(instanceId string, forceStop bool) (string, error) {
	return "", nil
}

func (ten *tencentcloudClient) ModifyInstanceName(instanceId string, instanceName string) (string, error) {
	return "", nil
}

func (ten *tencentcloudClient) ListInstance(instanceId string) (map[string]interface{}, error) {
	return nil, nil
}

func (ten *tencentcloudClient) ListInstances() ([]map[string]interface{}, error) {
	return nil, nil
}

func (ten *tencentcloudClient) ListRegions() ([]map[string]string, error) {
	return nil, nil
}
