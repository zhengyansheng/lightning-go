package message

type ResSendDingConfig struct {
	Subject    string `json:"subject"`     //  主题
	DingToken  string `json:"ding_token"`  //  04c381fc31944ad2905f31733e31fa15570ae12efc857062dab16b605a369e4c
	DingSecret string `json:"ding_secret"` //  SECb90923e19e58b466481e9e7b7a54531a3967fb29f0eae5c68
}

type ResSendMailConfig struct {
	Subject  string `json:"subject"`   //  主题名
	User     string `json:"user"`      //  用户名
	PassWord string `json:"pass_word"` //  密码
	Host     string `json:"host"`      //  主机名
	Port     string `json:"port"`      //  端口
}

type ResSendDing struct {
	MsgId int64    `json:"msg_id"` // 消息模版Id
	Data  struct { // 消息数据总字段
		Markdown struct { // 消息种类为Markdown
			Title string `json:"title"` // 消息数据 title
			Text  string `json:"text"`  // 消息数据 text
		} `json:"markdown"`
	} `json:"data"`
}

type DingMSg struct {
	Data struct { // 消息数据总字段
		Markdown struct { // 消息种类为Markdown
			Title string `json:"title"` // 消息数据 title
			Text  string `json:"text"`  // 消息数据 text
		} `json:"markdown"`
	} `json:"data"`
	MsgId      int64  `json:"msg_id"` // 消息模版Id
	DingToken  string `json:"ding_token"`
	DingSecret string `json:"ding_secret"`
}

type DingChan struct {
	Dm chan *DingMSg
}
