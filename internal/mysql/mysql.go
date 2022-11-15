package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing-platform-go/conf"
)

var Db *gorm.DB

func Init() {
	datasource := conf.Datasource
	url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		datasource.Username, datasource.Password, datasource.Ip, datasource.Port, datasource.Database)
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {

	}
	Db = db
}
