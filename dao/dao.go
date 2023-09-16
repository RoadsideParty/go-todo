package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectMysql() (err error) {
	dsn := "root:@tcp(127.0.0.1:3306)/my_database?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn))
	return
}

func Close() {
	sqlDb, _ := Db.DB()
	sqlDb.Close()
}
