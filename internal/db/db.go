package db

import (
	"github.com/douyu/jupiter/pkg/store/gorm"
)

var (
	DB      *gorm.DB
	Airflow *gorm.DB
)

func Init() {

	DB = gorm.StdConfig("master").Build()
	Airflow = gorm.StdConfig("airflow").Build()

}
