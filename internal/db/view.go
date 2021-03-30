package db

import (
	"github.com/jinzhu/gorm"
)

// 钉钉发送历史数据表
type SendDingHistory struct {
	gorm.Model
	MsgId          int64
	SendDingConfig []SendDingConfig `foreignkey:MsgId;association_foreignkey:MsgId"`
	Data           string           `json:"data" gorm:"column:data;type:text"`
	Status         bool
	ErrMsg         string `gorm:"column:err_msg;type:text"`
}

// 钉钉配置表
type SendDingConfig struct {
	gorm.Model
	MsgId      int64  `json:"msg_id"`      //  渠道ID
	Subject    string `json:"subject"`     //  主题
	DingToken  string `json:"ding_token"`  //  04c381fc31944ad2905f31733e31fa15570ae12efc857062dab16b605a369e4c
	DingSecret string `json:"ding_secret"` //  SECb90923e19e58b466481e9e7b7a54531a3967fb29f0eae5c68

}

// 邮件配置表
type SendMailConfig struct {
	gorm.Model
	MsgId    int64  `json:"msg_id"`    //  渠道ID
	Subject  string `json:"subject"`   //  主题名
	User     string `json:"user"`      //  用户名
	PassWord string `json:"pass_word"` //  密码
	Host     string `json:"host"`      //  主机名
	Port     string `json:"port"`      //  端口

}

// 钉钉分页查询结构体
type SendDingConfigQ struct {
	SendDingConfig
	PaginationQ
}

// 邮件分页查询结构体
type SendMailConfigQ struct {
	SendMailConfig
	PaginationQ
}
