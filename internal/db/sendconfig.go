package db

import (
	"go-ops/pkg/tools"
)

func (sd *SendDingConfig) TableName() string {
	return "send_ding_config"
}

// 拦截器: 不允许Token、user同时为空
func (sd *SendDingConfig) BeforeCreate() (err error) {

	if len(sd.DingToken) < 1 {
		return tools.New(204, "Token and user cannot be empty at the same time")
	}
	return
}

// 创建SendConfig
func (sd *SendDingConfig) SendCreate() (err error) {
	err = DB.Debug().Table(sd.TableName()).Create(sd).Error
	return
}

// 删除SendConfig
func (sd *SendDingConfig) SendDelete() (err error) {

	if sd.MsgIdExis() == false {
		return tools.New(204, "msg_id cannot be empty")
	}
	err = DB.Debug().Table(sd.TableName()).Where("msg_id = ?", sd.MsgId).Unscoped().Delete(sd).Error
	return
}

// 判断MsgId是否存在
func (sd *SendDingConfig) MsgIdExis() bool {
	var (
		exisSendConfig SendDingConfig
	)
	err := DB.Debug().Table(sd.TableName()).Where("msg_id = ?", sd.MsgId).First(&exisSendConfig).Error
	if err != nil {
		return false
	}
	return true
}

// 更新SendConfig
func (sd *SendDingConfig) SendUpdate() (err error) {

	if sd.MsgIdExis() == false {
		return tools.New(204, "msg_id cannot be empty")
	}
	err = DB.Debug().Table(sd.TableName()).Where("msg_id = ?", sd.MsgId).Update(sd).Error
	return
}

// 查询SendConfig
func (sd *SendDingConfig) GetSend() (SendDingConfig, error) {
	var respSendConfig SendDingConfig
	if sd.MsgIdExis() == false {
		return respSendConfig, tools.New(204, "msg_id cannot be empty")
	}
	DB.Debug().Table(sd.TableName()).Where("msg_id = ?", sd.MsgId).First(&respSendConfig)
	return respSendConfig, nil
}

// 分页查询Ding
func (q *SendDingConfigQ) SearchDing() (list *[]SendDingConfig, total int, err error) {
	list = &[]SendDingConfig{}
	//创建 db-query
	tx := DB.Debug().Table(q.SendDingConfig.TableName()).Where(q.SendDingConfig).Order("id desc")
	//CRUD
	total, err = CRUD(&q.PaginationQ, tx, list)
	return
}

func (sm *SendMailConfig) TableName() string {
	return "send_mail_config"
}

func (sm *SendMailConfig) SendUpdate() (err error) {

	err = DB.Debug().Table(sm.TableName()).Where("msg_id = ?", sm.MsgId).Update(sm).Error

	return nil
}

// 创建SendConfig
func (sm *SendMailConfig) SendCreate() (err error) {
	err = DB.Debug().Table(sm.TableName()).Create(sm).Error
	return
}

// 判断MsgId是否存在
func (sm *SendMailConfig) MsgIdExis() bool {
	var (
		exisSendConfig SendDingConfig
	)
	err := DB.Debug().Table(sm.TableName()).Where("msg_id = ?", sm.MsgId).First(&exisSendConfig).Error
	if err != nil {
		return false
	}
	return true
}

// 删除SendConfig
func (sm *SendMailConfig) SendDelete() (err error) {

	if sm.MsgIdExis() == false {
		return tools.New(204, "msg_id cannot be empty")
	}
	err = DB.Debug().Table(sm.TableName()).Where("msg_id = ?", sm.MsgId).Unscoped().Delete(sm).Error
	return
}

// 分页查询Ding
func (q *SendMailConfigQ) SearchDing() (list *[]SendMailConfig, total int, err error) {
	list = &[]SendMailConfig{}
	//创建 db-query
	tx := DB.Debug().Table(q.SendMailConfig.TableName()).Where(q.SendMailConfig).Order("id desc")
	//CRUD
	total, err = CRUD(&q.PaginationQ, tx, list)
	return
}
