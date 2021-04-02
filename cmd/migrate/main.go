package main

import (
	"fmt"

	// "strings"

	"github.com/douyu/jupiter"
	"github.com/douyu/jupiter/pkg/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"go-ops/internal/db"
	"go-ops/internal/models/multi_cloud"
)

const RECREATE = true

// Admin indicates the user is a system administrator.
type InstallAPP struct {
	jupiter.Application
}

func main() {
	eng := &InstallAPP{}
	_ = eng.Startup(
		eng.initApp,
	)
	initDB()
	db.Init()

}

// func migrateDB(cli *cli.Context) error {
func initDB() error {
	var dbName = "chaos"
	gormdb, err := gorm.Open(
		"mysql",
		conf.GetString("jupiter.mysql.master.dsn"),
	)

	if err != nil {
		return err
	}
	if RECREATE {
		var result []struct {
			Sqlstr string
		}
		gormdb.Debug().Raw("SELECT concat('DROP TABLE IF EXISTS `', table_name, '`;') as sqlstr FROM information_schema.tables WHERE table_schema = '" + dbName + "'").Scan(&result)
		for _, v := range result {
			fmt.Println(`sql drop:`, v.Sqlstr)

			gormdb.Exec(v.Sqlstr)
		}
	}

	defer func() {
		_ = gormdb.Close()
	}()

	models := []interface{}{
		&db.SendDingConfig{},
		&db.SendMailConfig{},
		&db.SendDingHistory{},
		&multi_cloud.CloudTemplate{},
		&multi_cloud.Account{},
	}
	if RECREATE {
		gormdb.DropTableIfExists(models...)
	}
	//

	// 删除原来的表
	gormdb.SingularTable(true)
	gormdb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models...)
	fmt.Println("create table ok")
	return nil

}

func (eng *InstallAPP) initApp() error {
	return nil
}
