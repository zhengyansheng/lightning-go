package multi_cloud_sdk

import (
	"go-ops/internal/models/multi_cloud"
)

type client interface {
	CreateInstance(PayType, hostname, instanceType, zoneId, imageId, vpcId, subnetId string, securityGroupIds []string, dryRun bool) (string, error)
	StartInstance(string) (string, error)
	StopInstance(string) (string, error)
	RebootInstance(string, bool) (string, error)
	DestroyInstance(string, bool) (string, error)
	ModifyInstanceName(string, string) (string, error)
	ListInstances() ([]map[string]interface{}, error)
	ListInstance(string) (map[string]interface{}, error)
	ListRegions() ([]map[string]string, error)
}

func NewFactory(cloudPlatform, accessKeyId, secretKeyId, regionId string) (clt client, err error) {
	switch cloudPlatform {
	case "ali":
		clt, err = NewAliEcs(accessKeyId, secretKeyId, regionId)
		return
	case "ten":
		clt, err = NewTenCvm(accessKeyId, secretKeyId, regionId)
		return
	}
	return
}

func NewFactoryByAccount(account, regionId string) (clt client, err error) {
	var (
		mc = multi_cloud.Account{EnName: account}
	)
	accessKeyInfo, err := mc.GetByAccount()
	if err != nil {
		return nil, err
	}
	switch accessKeyInfo.Platform {
	case "ali":
		clt, err = NewAliEcs(accessKeyInfo.AccessKeyId, accessKeyInfo.SecretKeyId, regionId)
		return
	case "ten":
		clt, err = NewTenCvm(accessKeyInfo.AccessKeyId, accessKeyInfo.SecretKeyId, regionId)
		return
	}
	return
}
