package db

func (s *SendDingHistory) TableName() string {
	return "send_ding_history"
}

func (s *SendDingHistory) Create() error {

	err := DB.Debug().Table(s.TableName()).Create(s).Error
	return err

}

type SendDingHistoryQ struct {
	SendDingHistory
	PaginationQ
}

func (h *SendDingHistoryQ) Search() (list *[]SendDingHistory, total int, err error) {
	list = &[]SendDingHistory{}
	//创建 db-query
	tx := DB.Model(h.SendDingHistory).Where(h.SendDingHistory).Order("id desc")
	//CRUD
	total, err = CRUD(&h.PaginationQ, tx, list)
	return
}
